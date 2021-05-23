package cowin

import (
	"encoding/json"
	"github.com/katta/jabfinder/pkg/models"
	"github.com/katta/jabfinder/pkg/table"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ListDistricts(stateCode int) {
	log.Printf("Checking supported districts for state code %v", stateCode)

	client := &http.Client{Timeout: cowinTimeout}
	request, err := http.NewRequest("GET", buildDistrictsQuery(stateCode), nil)
	exitOnError(err)

	request.Header.Add("user-agent", "Mozilla/5.0")

	response, err := client.Do(request)
	exitOnError(err)

	if response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		exitOnError(err)
		//log.Printf("Response: %v", string(body))

		var districtsResponse models.DistrictsResponse
		if err := json.Unmarshal(body, &districtsResponse); err == nil {
			//log.Printf("Centers: %+v", cowinResponse.Centers)
			printDistricts(districtsResponse)
		} else {
			exitOnError(err)
		}

	} else {
		log.Printf("Cowin responded with status code %v", response.StatusCode)
	}

}

func ListStates() {
	log.Printf("Checking supported states")

	client := &http.Client{Timeout: cowinTimeout}
	request, err := http.NewRequest("GET", buildStatesQuery(), nil)
	exitOnError(err)

	request.Header.Add("user-agent", "Mozilla/5.0")

	response, err := client.Do(request)
	exitOnError(err)

	if response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		exitOnError(err)
		//log.Printf("Response: %v", string(body))

		var statesResponse models.StatesResponse
		if err := json.Unmarshal(body, &statesResponse); err == nil {
			//log.Printf("Centers: %+v", cowinResponse.Centers)
			printStates(statesResponse)
		} else {
			exitOnError(err)
		}

	} else {
		log.Printf("Cowin responded with status code %v", response.StatusCode)
	}
}

func printStates(statesResponse models.StatesResponse) {
	headers := []string{"Name", "Code"}
	rows := [][]string{}

	for _, state := range statesResponse.States {
		row := []string{state.Name, strconv.Itoa(state.Code)}
		rows = append(rows, row)
	}

	table.Render(headers, rows, nil, false)
}

func printDistricts(response models.DistrictsResponse) {
	headers := []string{"Name", "Code"}
	rows := [][]string{}

	for _, district := range response.Districts {
		row := []string{district.Name, strconv.Itoa(district.Code)}
		rows = append(rows, row)
	}

	table.Render(headers, rows, nil, false)

}
