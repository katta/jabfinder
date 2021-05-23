package cowin

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/katta/jabfinder/pkg/notifiers"
	"github.com/katta/jabfinder/pkg/table"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const dateFormat = "02-01-2006"

var exit = make(chan bool)

func CheckAvailability(filters *Filters, notify bool) {
	log.Printf("Checking availability for: %+v", filters)

	if notify {
		go func() {
			for {
				availability := checkVaccineAvailability(filters)
				if availability != nil {
					notifiers.SendMail(createEmail(), smtpConfig())
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
		checkVaccineAvailability(filters)
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

func createEmail() notifiers.EMail {
	return notifiers.EMail{
		From:    "JabFinder <jabfinderindia@gmail.com>",
		To:      viper.GetString("notify.toEmail"),
		Subject: "Vaccination Slot Availability",
		Body:    "Here you go again !!",
	}
}

func checkVaccineAvailability(filters *Filters) [][]string {
	client := &http.Client{Timeout: 60 & time.Second}
	request, err := http.NewRequest("GET", buildAppointmentQuery(filters.DistrictCode), nil)
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

		var cowinResponse CowinResponse
		err = json.Unmarshal(body, &cowinResponse)
		exitOnError(err)

		return printAvailability(cowinResponse, filters)
	} else {
		log.Printf("Cowin responded with status code %v", response.StatusCode)
	}
	return nil
}

func onError(err error, notify bool) bool {
	if err != nil {
		log.Printf("Error while checking availability on cowin: %+v", err)
		if !notify {
			exit <- true
		} else {
			return true
		}
	}
	return false
}

func printAvailability(response CowinResponse, filters *Filters) [][]string {
	headers := []string{"Date", "Vaccine", "Dose 1", "Dose 2", "Center", "Address"}
	rows := [][]string{}

	log.Printf("Filters: %+v", filters)

	for _, center := range response.Centers {
		for _, session := range center.Sessions {
			if session.AvailableCapacity > 0 && session.MinAge == filters.Age {
				if filters.Dose == 1 && session.AvailableCapacityDose1 > 0 {
					rows = appendCenter(center, session, rows)
				} else if filters.Dose == 2 && session.AvailableCapacityDose2 > 0 {
					rows = appendCenter(center, session, rows)
				} else if filters.Dose == 0 {
					rows = appendCenter(center, session, rows)
				}
			}
		}
	}

	table.Render(headers, rows, []string{}, true)
	return rows
}

func appendCenter(center Center, session Session, rows [][]string) [][]string {
	address := fmt.Sprintf("%s, %d", center.Address, center.Pincode)
	row := []string{session.Date,
		session.Vaccine,
		strconv.Itoa(session.AvailableCapacityDose1),
		strconv.Itoa(session.AvailableCapacityDose2),
		//strings.Join(session.Slots[:], ","),
		center.Name,
		address,
	}
	rows = append(rows, row)
	return rows
}
