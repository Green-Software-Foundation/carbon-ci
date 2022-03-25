package watttime

import (
	"encoding/base64"
	"strconv"
)

const url string = "https://api2.watttime.org/v2/"

var token string

// LOGIN - RESPONSE object { token }
func Login(username string, password string) error {
	header := make(map[string]string)

	header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

	var response loginResponse

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

// DETERMINE GRID REGION - RESPONSE object { id, abbrev, name }
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

// LIST OF GRID REGIONS - RESPONSE object { ba, name, access, datatype }
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

// REAL-TIME EMISSIONS INDEX - RESPONSE object { freq, ba, percent, moer, point_time }
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

// GRID EMISSIONS DATA - RESPONSE array [object { ba,  datatype, market, point_time, value, version }]
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

// HISTORICAL EMISSIONS - RESPONSE A binary zip file payload containing monthly .csv files with MOERs for the specified balancing authority for the past two years. Save the body to disk and unzip it. : Binary Data Zip File
func HistoricalEmissions(ba string, version string) {
	// RETURN CSV
}

//	EMISSIONS FORECAST - RESPONSE object { generated_at, forecast [{ba, point_time, value, version}]}
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

//	GRID REGION MAP GEOMETRY - RESPONSE A geojson response, that is a Feature Collection with properties that describe each BA, and multipolygon geometry made up of coordinates which define the boundary for each BA. The “meta” object contains the date-time that the geojson was last updated.
func GridRegionMapGeometry() {
	// RETURN CSV
}
