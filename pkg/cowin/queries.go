package cowin

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"time"
)

const cowinTimeout = 60 & time.Second

func buildAppointmentQuery(district string) string {
	cowinUrl := viper.GetString("cowin.baseurl")
	appointmentsPath := viper.GetString("cowin.appointmentsPath")

	query, err := url.Parse(fmt.Sprintf("%s%s", cowinUrl, appointmentsPath))
	exitOnError(err)

	values := url.Values{}
	values.Set("district_id", district)
	values.Set("date", time.Now().Format(dateFormat))
	query.RawQuery = values.Encode()

	log.Printf("Using Query: %v", query.String())
	return query.String()
}

func buildStatesQuery() string {
	cowinUrl := viper.GetString("cowin.baseurl")
	statesPath := viper.GetString("cowin.statesPath")

	query, err := url.Parse(fmt.Sprintf("%s%s", cowinUrl, statesPath))
	exitOnError(err)

	log.Printf("Using Query: %v", query.String())
	return query.String()
}

func buildDistrictsQuery(stateCode int) string {
	cowinUrl := viper.GetString("cowin.baseurl")
	districtsPath := viper.GetString("cowin.districtsPath")

	query, err := url.Parse(fmt.Sprintf("%s%s/%d", cowinUrl, districtsPath, stateCode))
	exitOnError(err)

	log.Printf("Using Query: %v", query.String())
	return query.String()
}

func exitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func warnOnError(err error) bool {
	if err != nil {
		log.Printf("Error: %+v", err)
		return true
	}
	return false
}
