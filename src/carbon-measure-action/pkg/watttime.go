package watttime

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const URL string = "https://api2.watttime.org/v2/"

var token string

// <GRID EMISSIONS INFORMATION>
// LOGIN - RESPONSE object { token }
type LoginResponse struct {
	Token string `json:"token"`
}

func Login(username string, password string) error {
	header := make(map[string]string)

	header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"

	var response LoginResponse

	request := Headers{
		Url:      URL + "login",
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
type DetermineGridRegionResp struct {
	Abbrev string `json:"abbrev"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
}

func DetermineGridRegion(latitude float32, longitude float32) (*DetermineGridRegionResp, error) {
	header := make(map[string]string)

	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Bearer " + token

	query := make(map[string]string)
	query["latitude"] = strconv.FormatFloat(float64(latitude), 'E', -1, 32)
	query["longitude"] = strconv.FormatFloat(float64(longitude), 'E', -1, 32)

	var response DetermineGridRegionResp

	err := httpRequest(Headers{
		Url:      URL + "ba-from-loc",
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
type ListOfGridRegionsResp struct {
	BA       string `json:"ba"`
	Name     string `json:"name"`
	Accept   bool   `json:"accept"`
	DataType string `json:"datatype"`
}

func ListOfGridRegions(all bool) (*[]ListOfGridRegionsResp, error) {
	header := make(map[string]string)

	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Bearer " + token

	query := make(map[string]string)
	query["latitude"] = strconv.FormatBool(all)

	var response []ListOfGridRegionsResp

	err := httpRequest(Headers{
		Url:      URL + "ba-access",
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
type RealTimeEmissionsIndexResp struct {
	Freq      string `json:"freq"`
	BA        string `json:"ba"`
	Percent   string `json:"percent"`
	Moer      string `json:"moer"`
	PointTime string `json:"point_time"`
}

func RealTimeEmissionsIndex(ba string, latitude float32, longitude float32, style string) (*RealTimeEmissionsIndexResp, error) {
	header := make(map[string]string)

	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
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

	var response RealTimeEmissionsIndexResp

	err := httpRequest(Headers{
		Url:      URL + "index",
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
type GridEmissionsDataResp struct {
	BA        string  `json:"ba"`
	DataType  string  `json:"datatype"`
	Frequency int     `json:"frequency"`
	Market    string  `json:"market"`
	PointTime string  `json:"point_time"`
	Value     float32 `json:"value"`
	Version   string  `json:"version"`
}

func GridEmissionsData(ba string, latitude float32, longitude float32, starttime string, endtime string, style string, moerversion string) (*[]GridEmissionsDataResp, error) {
	header := make(map[string]string)

	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
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

	var response []GridEmissionsDataResp

	err := httpRequest(Headers{
		Url:      URL + "data",
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

}

//	EMISSIONS FORECAST - RESPONSE object { generated_at, forecast [{ba, point_time, value, version}]}
type EmissionForecastResp struct {
	GeneratedAt string     `json:"generated_at"`
	Forecast    []Forecast `json:"forecast"`
}
type Forecast struct {
	BA        string  `json:"ba"`
	PointTime string  `json:"point_time"`
	Value     float32 `json:"value"`
	Version   string  `json:"version"`
}

func EmissionsForecast(ba string, starttime string, endtime string, extendedForecast bool) (*[]EmissionForecastResp, error) {
	header := make(map[string]string)

	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
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

	var response []EmissionForecastResp

	err := httpRequest(Headers{
		Url:      URL + "forecast",
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

}

// </GRID EMISSIONS INFORMATION>

// <HTTP REQUEST>

// REQUEST FUNCTIONS
type Headers struct {
	Url      string
	Method   string
	Data     map[string]string
	Header   map[string]string
	Query    map[string]string
	Response interface{}
}

func initRequest(method string, url string, data map[string]string) (*http.Request, error) {
	if len(method) == 0 {
		return http.NewRequest(method, url, nil)
	} else {
		json, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return http.NewRequest(method, url, bytes.NewBuffer(json))
	}
}

func httpRequest(headers Headers) error {
	client := &http.Client{}

	// INITIALIZE NEW REQUEST
	req, err := initRequest(headers.Method, headers.Url, headers.Data)

	if err != nil {
		return err
	}

	// SET HEADERS/S
	for h := range headers.Header {
		req.Header.Add(h, headers.Header[h])
	}

	// SET QUERY STRING
	query := req.URL.Query()
	for q := range headers.Query {
		query.Add(q, headers.Query[q])
	}
	req.URL.RawQuery = query.Encode()

	fmt.Println(req.URL.RawQuery)

	// DO REQUEST
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// EXECUTE AFTER FUNCTION RETURNS
	defer resp.Body.Close()

	// GET RESPONSE DATA
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// CONVERT RESPONSE DATA FROM BYTE ARRAY TO OBJECT
	err = json.Unmarshal(bodyBytes, &headers.Response)
	if err != nil {
		return err
	}

	return nil
}

// </HTTP REQUESST

// func main() {
// 		REMARKS : PLEASE LOGIN FIRST BEFORE CALLING THE GRID EMISSIONS INFORMATION
// 		Login("Jerrico", "@WattTime123")
// 		fmt.Print(DetermineGridRegion(42.372, -72.519))
// 		fmt.Println(ListOfGridRegions(false))
// 		fmt.Println(RealTimeEmissionsIndex("CAISO_NORTH", 0, 0, ""))
// 		fmt.Println(GridEmissionsData("CAISO_NORTH", 0, 0, "2019-02-20T16:00:00-0800", "2019-02-20T16:15:00-0800", "", ""))
// 		fmt.Println(EmissionsForecast("CAISO_NORTH", "2022-03-17T00:00:00-0400", "2022-03-17T17:00:00-0400", false))
// }
