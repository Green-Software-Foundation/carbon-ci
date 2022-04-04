// Package watttime technology—based on real-time grid data, cutting-edge algorithms, and machine learning—provides first-of-its-kind
// insight into your local electricity grid’s marginal emissions rate.
package watttime

import (
	"encoding/base64"
	"strconv"
)

const url string = "https://api2.watttime.org/v2/"

var token string

// Login obtain an access token, it returns an error if failed.
func Login(username string, password string) error {
	header := make(map[string]string)

	header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

	var response loginResp

	request := httpRequestType{
		Url:      url + "login",
		Method:   "GET",
		Header:   header,
		Response: &response,
	}

	err := httpRequest(request)
	if err != nil {
		return err
	}

	token = response.Token
	return nil
}

// DetermineGridRegion returns the details of the balancing authority (BA) serving that location, if known, or a Coordinates not found error if the point lies outside of known/covered BAs.
func DetermineGridRegion(latitude float32, longitude float32) (*determineGridRegionResp, error) {
	header := make(map[string]string)

	header["Authorization"] = "Bearer " + token

	query := make(map[string]string)
	query["latitude"] = strconv.FormatFloat(float64(latitude), 'E', -1, 32)
	query["longitude"] = strconv.FormatFloat(float64(longitude), 'E', -1, 32)

	var response determineGridRegionResp

	err := httpRequest(httpRequestType{
		Url:      url + "ba-from-loc",
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ListOfGridRegions by default this function delivers a list of regions to which you have access. Optionally, it can return a list of all grid regions where WattTime has data coverage.
func ListOfGridRegions(all bool) (*[]listOfGridRegionsResp, error) {
	header := make(map[string]string)

	header["Authorization"] = "Bearer " + token

	query := make(map[string]string)
	query["all"] = strconv.FormatBool(all)

	var response []listOfGridRegionsResp

	err := httpRequest(httpRequestType{
		Url:      url + "ba-access",
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RealTimeEmissionsIndex provides a real-time signal indicating the marginal carbon intensity for the local grid for the current time (updated every 5 minutes).
func RealTimeEmissionsIndex(ba string, latitude float32, longitude float32, style string) (*realTimeEmissionsIndexResp, error) {
	header := make(map[string]string)

	header["Authorization"] = "Bearer " + token

	query := make(map[string]string)
	if ba != "" {
		query["ba"] = ba
	} else {
		query["latitude"] = strconv.FormatFloat(float64(latitude), 'E', -1, 32)
		query["longitude"] = strconv.FormatFloat(float64(longitude), 'E', -1, 32)
	}

	if style != "" {
		query["style"] = style
	}

	var response realTimeEmissionsIndexResp

	err := httpRequest(httpRequestType{
		Url:      url + "index",
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GridEmissionsData obtain historical marginal emissions (CO2 MOER in lbs of CO2 per MWh) for a given grid region (balancing authority abbreviated code, ba) or location (latitude & longitude pair).
func GridEmissionsData(ba string, latitude float32, longitude float32, starttime string, endtime string, style string, moerversion string) (*[]gridEmissionsDataResp, error) {
	header := make(map[string]string)

	header["Authorization"] = "Bearer " + token

	query := make(map[string]string)
	if ba != "" {
		query["ba"] = ba
	} else {
		query["latitude"] = strconv.FormatFloat(float64(latitude), 'E', -1, 32)
		query["longitude"] = strconv.FormatFloat(float64(longitude), 'E', -1, 32)
	}

	if starttime != "" {
		query["starttime"] = starttime
	}

	if endtime != "" {
		query["endtime"] = endtime
	}

	if style != "" {
		query["style"] = style
	}

	if style != "" {
		query["moerversion"] = moerversion
	}

	var response []gridEmissionsDataResp

	err := httpRequest(httpRequestType{
		Url:      url + "data",
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &response,
	})

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// HistoricalEmissions obtain a zip file containing monthly .csv files with the MOER values and timestamps for a given region for (up to) the past two years.
func HistoricalEmissions(ba string, version string) {
	// RETURN CSV
}

// EmissionForecast obtain MOER forecast data for a given region.
func EmissionsForecast(ba string, starttime string, endtime string, extendedForecast bool) (*[]emissionForecastResp, error) {
	header := make(map[string]string)

	header["Authorization"] = "Bearer " + token

	query := make(map[string]string)
	query["ba"] = ba

	if starttime != "" {
		query["starttime"] = starttime
	}

	if endtime != "" {
		query["endtime"] = endtime
	}

	if extendedForecast {
		query["extended_forecast"] = strconv.FormatBool(extendedForecast)
	}

	var response []emissionForecastResp

	err := httpRequest(httpRequestType{
		Url:      url + "forecast",
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &response,
	})

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GridRegionMapGeometry provides a geojson of the grid region boundary for all regions that WattTime covers globally.
func GridRegionMapGeometry() {
	// RETURN CSV
}
