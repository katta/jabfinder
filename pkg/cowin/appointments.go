package cowin

import (
	"encoding/json"
	"fmt"
	"github.com/katta/jabfinder/pkg/table"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const dateFormat = "02-01-2006"

func CheckAvailability(district string, filters *Filters) {
	log.Printf("Checking availability for %v, %+v", district, filters)

	client := &http.Client{Timeout: 60 & time.Second}
	request, err := http.NewRequest("GET", buildAppointmentQuery(district), nil)
	exitOnError(err)

	request.Header.Add("user-agent", "Mozilla/5.0")

	response, err := client.Do(request)
	exitOnError(err)

	if response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		exitOnError(err)
		//log.Printf("Response: %v", string(body))

		var cowinResponse CowinResponse
		if err := json.Unmarshal(body, &cowinResponse); err == nil {
			//log.Printf("Centers: %+v", cowinResponse.Centers)
			printAvailability(cowinResponse, filters)
		} else {
			exitOnError(err)
		}

	} else {
		log.Printf("Cowin responded with status code %v", response.StatusCode)
	}
}

func printAvailability(response CowinResponse, filters *Filters) {
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
