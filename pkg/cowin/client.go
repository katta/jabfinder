package cowin

import (
	"encoding/json"
	"fmt"
	"github.com/katta/jabfinder/pkg/table"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const dateFormat = "02-01-2006"

func CheckAvailability(district string, age int) {
	log.Printf("Checking availability for %v, %v", district, age)

	client := &http.Client{Timeout: 60 & time.Second}
	request, err := http.NewRequest("GET", buildQuery(district), nil)
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
			log.Printf("Centers: %+v", cowinResponse.Centers)
			printAvailability(cowinResponse, age)
		} else {
			exitOnError(err)
		}

	} else {
		log.Printf("Cowin responded with status code %v", response.StatusCode)
	}
}

func printAvailability(response CowinResponse, age int) {
	headers := []string{"Date", "Vaccine", "Dose 1", "Dose 2", "Center", "Address"}
	rows := [][]string{}

	for _, center := range response.Centers {
		for _, session := range center.Sessions {
			if session.AvailableCapacity > 0 && session.MinAge == age {
				address := fmt.Sprintf("%s, %d", center.Address, center.Pincode)
				row := []string{session.Date, session.Vaccine, strconv.Itoa(session.AvailableCapacityDose1), strconv.Itoa(session.AvailableCapacityDose2), center.Name, address}
				rows = append(rows, row)
			}
		}
	}

	table.Render(headers, rows, []string{})
}

func buildQuery(district string) string {
	cowinUrl := viper.GetString("cowin.baseurl")
	log.Printf("Using %v to check for availability", cowinUrl)
	query, err := url.Parse(cowinUrl)
	exitOnError(err)

	values := url.Values{}
	values.Set("district_id", district)
	values.Set("date", time.Now().Format(dateFormat))
	query.RawQuery = values.Encode()

	log.Printf("Query with parameters: %v", query.String())
	return query.String()
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
