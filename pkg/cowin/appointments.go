package cowin

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/katta/jabfinder/pkg/db"
	"github.com/katta/jabfinder/pkg/models"
	"github.com/katta/jabfinder/pkg/notifiers"
	"github.com/katta/jabfinder/pkg/table"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "02-01-2006"

var exit = make(chan bool)

func CheckAvailability(filters *models.Filters, notify bool) {
	log.Printf("Checking availability for: %+v", filters)

	if notify {
		go func() {
			for {
				availableSessions := retrieveAvailableSessions(filters)

				printToConsole(availableSessions)

				newSessions := db.Register(availableSessions)
				if len(newSessions) > 0 {
					notifyByEmail(newSessions)
				}

				interval := viper.GetInt("notify.intervalInSeconds")
				color.Set(color.FgHiGreen)
				log.Printf("Will check again in %v seconds.. \n", interval)
				color.Unset()

				time.Sleep(time.Duration(interval) * time.Second)
			}
		}()
		<-exit
	} else {
		availableSessions := retrieveAvailableSessions(filters)
		printToConsole(availableSessions)
	}
}

func notifyByEmail(sessions []models.FlatSession) {
	if sessions != nil {
		mailer := &notifiers.Mailer{
			EMail: emailConfig(),
			SMTP:  smtpConfig(),
		}
		mailer.Notify(sessions)
	}
}

func smtpConfig() notifiers.SMTP {
	return notifiers.SMTP{
		Host:     viper.GetString("smtp.host"),
		Port:     viper.GetInt("smtp.port"),
		Email:    viper.GetString("smtp.email"),
		Password: viper.GetString("smtp.password"),
	}
}

func emailConfig() notifiers.EMail {
	return notifiers.EMail{
		From:    "JabFinder <jabfinderindia@gmail.com>",
		To:      viper.GetString("notify.toEmail"),
		Subject: "Vaccination Slot Availability",
		Body:    "Here you go again !!",
	}
}

func retrieveAvailableSessions(filters *models.Filters) []models.FlatSession {
	client := &http.Client{Timeout: 60 & time.Second}

	date := filters.Date
	if date == "" {
		date = time.Now().Format(dateFormat)
	}

	request, err := http.NewRequest("GET", buildAppointmentQuery(filters.DistrictCode, date), nil)
	exitOnError(err)

	request.Header.Add("user-agent", "Mozilla/5.0")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error while checking availability on cowin: %+v", err)
		return nil
	}

	if response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Error while reading the response from cowin: %+v", err)
			return nil
		}
		//log.Printf("Response: %v", string(body))

		var cowinResponse models.CowinResponse
		err = json.Unmarshal(body, &cowinResponse)
		exitOnError(err)

		return filterAvailableSessions(cowinResponse, filters)
	} else {
		log.Printf("Cowin responded with status code %v", response.StatusCode)
	}
	return nil
}

func filterAvailableSessions(response models.CowinResponse, filters *models.Filters) []models.FlatSession {
	flatSessions := []models.FlatSession{}

	for _, center := range response.Centers {
		for _, session := range center.Sessions {
			if session.AvailableCapacity > 0 && session.MinAge == filters.Age && strings.EqualFold(session.Vaccine, filters.Vaccine) {
				if filters.Dose == 1 && session.AvailableCapacityDose1 > 0 {
					flatSessions = append(flatSessions, models.FlatSessionsFrom(center, session))
				} else if filters.Dose == 2 && session.AvailableCapacityDose2 > 0 {
					flatSessions = append(flatSessions, models.FlatSessionsFrom(center, session))
				} else if filters.Dose == 0 {
					flatSessions = append(flatSessions, models.FlatSessionsFrom(center, session))
				}
			}
		}
	}

	return flatSessions
}

func printToConsole(flatSessions []models.FlatSession) {
	headers := []string{"Date", "Vaccine", "Dose 1", "Dose 2", "Center", "Address"}
	rows := [][]string{}

	for _, fSession := range flatSessions {
		rows = toTableRow(fSession, rows)
	}

	table.Render(headers, rows, []string{}, true)
}

func toTableRow(flatSession models.FlatSession, rows [][]string) [][]string {
	address := fmt.Sprintf("%s, %d", flatSession.CenterAddress, flatSession.CenterPincode)
	row := []string{flatSession.SessionDate,
		flatSession.Vaccine,
		strconv.Itoa(flatSession.AvailableCapacityDose1),
		strconv.Itoa(flatSession.AvailableCapacityDose2),
		//strings.Join(session.Slots[:], ","),
		flatSession.CenterName,
		address,
	}
	rows = append(rows, row)
	return rows
}
