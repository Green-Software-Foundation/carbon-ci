package watttime

import (
	"errors"
	"os"
	"testing"
)

var (
	forbidden = "403 Forbidden"
)

/*
PLEASE SET ENVIRONMENT VARIABLES FOR WATTTIME CREDENTIALS

BASH

	export WATTTIME_USERNAME = "<USERNAME>"
	export WATTTIME_PASSWORD = "<PASSWORD>"

POWERSHELL OR CMD

	$Env:WATTTIME_USERNAME = "<USERNAME>"
	$Env:WATTTIME_PASSWORD = "<PASSWORD>"
*/
func getCredentials() (string, string) {
	username := os.Getenv("WATTTIME_USERNAME")
	password := os.Getenv("WATTTIME_PASSWORD")
	return username, password
}

func TestLogin(t *testing.T) {
	username, password := getCredentials()

	var credentialTests = []struct {
		username string
		password string
		out      error
	}{
		{"", "", errors.New(forbidden)},
		{"FakeUsername", "FakePassword", errors.New(forbidden)},
		{username, password, nil},
	}

	for _, test := range credentialTests {
		err := Login(test.username, test.password)
		if test.out == err {
			t.Log("Logged in successfully!")
		} else if err.Error() != test.out.Error() {
			t.Errorf("\nUsername: %s\nPassword: %s\nRESULT\n%s\nEXPECTED\n%s", test.username, test.password, err, test.out)
		}
	}
}

func TestListOfGridRegions(t *testing.T) {
	testCases := []struct {
		all bool
	}{
		{true},
		{false},
	}

	username, password := getCredentials()
	Login(username, password)
	for _, testCase := range testCases {
		data, err := ListOfGridRegions(testCase.all)
		if err != nil {
			t.Errorf("%s", err)
		} else {
			t.Log(data)
		}
	}
}

func TestDetermineGridRegion(t *testing.T) {
	testCases := []struct {
		latitude  float32
		longitude float32
	}{
		{42.372, -72.519},
	}

	username, password := getCredentials()
	Login(username, password)
	for _, testCase := range testCases {
		data, err := DetermineGridRegion(testCase.latitude, testCase.longitude)

		if err != nil {
			t.Errorf("%s", err)
		} else {
			t.Log(data)
		}
	}
}

func TestRealTimeEmissionsIndex(t *testing.T) {
	testCases := []struct {
		ba        string
		latitude  float32
		longitude float32
		style     string
	}{
		{"CAISO_NORTH", 0, 0, ""},
	}

	username, password := getCredentials()
	Login(username, password)
	for _, testCase := range testCases {
		data, err := RealTimeEmissionsIndex(testCase.ba, testCase.latitude, testCase.longitude, testCase.style)

		if err != nil {
			t.Errorf("%s", err)
		} else {
			t.Logf("FOUND\nFREQ:%s\nBA:%s\nPERCENT:%s\nMOER:%s\nPOINT_TIME:%s", data.Freq, data.BA, data.Percent, data.Moer, data.PointTime)
		}
	}
}

func TestGridEmissionsData(t *testing.T) {
	testCases := []struct {
		ba          string
		latitude    float32
		longitude   float32
		starttime   string
		endtime     string
		style       string
		moerversion string
	}{
		{"CAISO_NORTH", 0, 0, "2019-02-20T16:00:00-0800", "2019-02-20T16:15:00-0800", "", ""},
	}

	username, password := getCredentials()
	Login(username, password)
	for _, testCase := range testCases {
		data, err := GridEmissionsData(testCase.ba, testCase.latitude, testCase.longitude, testCase.starttime, testCase.endtime, testCase.style, testCase.moerversion)

		if err != nil {
			t.Errorf("%s", err)
		} else {
			t.Log(data)
		}
	}
}

func TestEmissionsForecast(t *testing.T) {
	testCases := []struct {
		ba                string
		starttime         string
		endtime           string
		extended_forecast bool
	}{
		{"CAISO_NORTH", "2021-08-05T09:00:00-0400", "2021-08-05T09:05:00-0400", false},
		{"CAISO_NORTH", "2021-08-05T09:00:00-0400", "2021-10-05T09:05:00-0400", true},
	}

	username, password := getCredentials()
	Login(username, password)
	for _, testCase := range testCases {
		data, err := EmissionsForecast(testCase.ba, testCase.starttime, testCase.endtime, testCase.extended_forecast)

		if err != nil {
			t.Errorf("%s", err)
		} else {
			t.Log(data)
		}
	}
}
