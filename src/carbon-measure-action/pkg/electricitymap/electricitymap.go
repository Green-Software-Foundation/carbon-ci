package electricitymap

import (
	"fmt"

	"main/pkg/http"

	"strconv"
)

type electricityMap struct {
	zoneKey string
	url     string
}

func New(zoneKey string) electricityMap {
	em := electricityMap{
		zoneKey: zoneKey,
		url:     "https://api.electricitymap.org/v3",
	}
	return em
}

func httpQueryBuilder(zoneKey string, params TypAPIParams) (header map[string]string, query map[string]string) {
	header = make(map[string]string)
	query = make(map[string]string)

	header["auth-token"] = zoneKey

	if params.Zone != "" {
		query["zone"] = params.Zone
	}
	if params.Lon != "" && params.Lat != "" {
		query["lon"] = params.Lon
		query["lat"] = params.Lat
	}
	if params.Datetime != "" {
		query["datetime"] = params.Datetime
	}
	if params.Start != "" {
		query["start"] = params.Start
	}
	if params.End != "" {
		query["end"] = params.End
	}
	if params.EstimationFallback == true {
		query["estimationFallback"] = strconv.FormatBool(params.EstimationFallback)
	}

	return
}

/*
This endpoint returns all zones available if no auth-token is provided.

If an auth-token is provided, it returns a list of zones and routes available with this token
*/
func (e electricityMap) GetZones() (map[string]TypZone, error) {
	url := fmt.Sprintf("%v/zones", e.url)
	data := make(map[string]TypZone)
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = e.zoneKey

	fmt.Println("Getting Electricity Map Zones")

	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

/*
This endpoint retrieves the last known carbon intensity (in gCO2eq/kWh) of electricity consumed in an area. It can either be queried by zone identifier or by geolocation.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)
*/
func (e electricityMap) LiveCarbonIntensity(params TypAPIParams) (TypCI, error) {
	url := fmt.Sprintf("%v/carbon-intensity/latest", e.url)
	var data TypCI

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Live Carbon Intensity")

	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}


	err := request.Send()
	return data, err
}

/*
This endpoint retrieves the last known data about the origin of electricity in an area.

 - "powerProduction" (in MW) represents the electricity produced in the zone, broken down by production type

 - "powerConsumption" (in MW) represents the electricity consumed in the zone, after taking into account imports and exports, and broken down by production type.

 - "powerExport" and "Power import" (in MW) represent the physical electricity flows at the zone border

 - "renewablePercentage" and "fossilFreePercentage" refers to the % of the power consumption breakdown coming from renewables or fossil-free power plants (renewables and nuclear) It can either be queried by zone identifier or by geolocation.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)
*/
func (e electricityMap) LivePowerBreakdown(params TypAPIParams) (TypPB, error) {
	url := fmt.Sprintf("%v/power-breakdown/latest", e.url)
	var data TypPB

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Live Power Breakdown")
	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

/*
This endpoint retrieves the last 24h of carbon intensity (in gCO2eq/kWh) of an area. It can either be queried by zone identifier or by geolocation. The resolution is 60 minutes.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)
*/
func (e electricityMap) RecentCarbonIntensity(params TypAPIParams) (TypRecentCI, error) {
	url := fmt.Sprintf("%v/carbon-intensity/history", e.url)
	var data TypRecentCI

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Recent Carbon Intensity")
	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

/*
This endpoint retrieves the last 24h of power consumption and production breakdown of an area, which represents the physical origin of electricity broken down by production type. It can either be queried by zone identifier or by geolocation. The resolution is 60 minutes.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)
*/
func (e electricityMap) RecentPowerBreakdown(params TypAPIParams) (TypRecentPB, error) {
	url := fmt.Sprintf("%v/power-consumption-breakdown/history", e.url)
	var data TypRecentPB

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Recent Power Breakdown")
	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

/*
This endpoint retrieves a past carbon intensity (in gCO2eq/kWh) of an area. It can either be queried by zone identifier or by geolocation. The resolution is 60 minutes.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)

datetime | datetime in ISO format

estimationFallback | (optional) boolean (if estimated data should be included)
*/
func (e electricityMap) PastCarbonIntensity(params TypAPIParams) (TypCI, error) {
	url := fmt.Sprintf("%v/carbon-intensity/past", e.url)
	var data TypCI

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Past Carbon Intensity")
	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

/*
This endpoint retrieves a past carbon intensity (in gCO2eq/kWh) of an area within a given date range. It can either be queried by zone identifier or by geolocation. The resolution is 60 minutes. The time range is limited to 10 days.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)

start | datetime in ISO format

end | datetime in ISO format (excluded)

estimationFallback | (optional) boolean (if estimated data should be included)
*/
func (e electricityMap) PastCarbonIntensityRange(params TypAPIParams) (map[string][]TypCI, error) {
	url := fmt.Sprintf("%v/carbon-intensity/past-range", e.url)
	var data = make(map[string][]TypCI)

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Past Carbon Intensity Range")
	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

/*
This endpoint retrieves a past power breakdown of an area. It can either be queried by zone identifier or by geolocation. The resolution is 60 minutes.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)

datetime | datetime in ISO format

estimationFallback | (optional) boolean (if estimated data should be included)
*/
func (e electricityMap) PastPowerBreakdown(params TypAPIParams) (TypPB, error) {
	url := fmt.Sprintf("%v/power-breakdown/past", e.url)
	var data TypPB

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Past Power Breakdown")
	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

/*
This endpoint retrieves a past power breakdown of an area within a given date range. It can either be queried by zone identifier or by geolocation. The resolution is 60 minutes. The time range is limited to 10 days.

QUERY PARAMETERS

Parameter | Description

zone | A string representing the zone identifier

lon | Longitude (if querying with a geolocation)

lat | Latitude (if querying with a geolocation)

start | datetime in ISO format

end | datetime in ISO format (excluded)

estimationFallback | (optional) boolean (if estimated data should be included)
*/
func (e electricityMap) PastPowerBreakdownRange(params TypAPIParams) (map[string][]TypPB, error) {
	url := fmt.Sprintf("%v/power-breakdown/past-range", e.url)
	var data = make(map[string][]TypPB)

	header, query := httpQueryBuilder(e.zoneKey, params)

	fmt.Println("Getting Electricity Map Past Power Breakdown Range")
	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

type TypAPIParams struct {
	Zone               string
	Lon                string
	Lat                string
	Datetime           string
	Start              string
	End                string
	EstimationFallback bool
}

type TypCI struct {
	Zone            string `json:"zone"`
	CarbonIntensity int    `json:"carbonIntensity"`
	Datetime        string `json:"datetime"`
	UpdatedAt       string `json:"updatedAt"`
	CreatedAt       string `json:"createdAt"`
}

type TypPB struct {
	Zone                      string                       `json:"zone"`
	Datetime                  string                       `json:"datetime"`
	PowerProductionBreakdown  TypPowerProductionBreakdown  `json:"powerProductionBreakdown"`
	PowerProductionTotal      int                          `json:"powerProductionTotal"`
	PowerConsumptionBreakdown TypPowerConsumptionBreakdown `json:"powerConsumptionBreakdown"`
	PowerConsumptionTotal     int                          `json:"powerConsumptionTotal"`
	PowerImportBreakdown      TypPowerImpExpBreakdown      `json:"powerImportBreakdown"`
	PowerImportTotal          int                          `json:"powerImportTotal"`
	PowerExportBreakdown      TypPowerImpExpBreakdown      `json:"powerExportBreakdown"`
	PowerExportTotal          int                          `json:"powerExportTotal"`
	FossilFreePercentage      int                          `json:"fossilFreePercentage"`
	RenewablePercentage       int                          `json:"renewablePercentage"`
	UpdatedAt                 string                       `json:"updatedAt"`
	CreatedAt                 string                       `json:"createdAt"`
}

type TypPowerConsumptionBreakdown struct {
	BatteryDischarge string // battery discharge `json:"batteryDischarge"`
	Biomass          int    `json:"biomass"`
	Coal             int    `json:"coal"`
	Gas              int    `json:"gas"`
	Geothermal       int    `json:"geothermal"`
	Hydro            int    `json:"hydro"`
	HydroDischarge   int    //hydro discharge `json:"hydroDischarge"`
	Nuclear          int    `json:"nuclear"`
	Oil              int    `json:"oil"`
	Solar            int    `json:"solar"`
	Unknown          int    `json:"unknown"`
	Wind             int    `json:"wind"`
}

type TypPowerImpExpBreakdown struct {
	DE     int `json:"DE"`
	DK_DK1 int //DK-DK1 `json:"DK_DK1"`
	SE     int `json:"SE"`
}

type TypPowerProductionBreakdown struct {
	Biomass    int `json:"biomass"`
	Coal       int `json:"coal"`
	Gas        int `json:"gas"`
	Geothermal int `json:"geothermal"`
	Hydro      int `json:"hydro"`
	Nuclear    int `json:"nuclear"`
	Oil        int `json:"oil"`
	Solar      int `json:"solar"`
	Unknown    int `json:"unknown"`
	Wind       int `json:"wind"`
}

type TypZone struct {
	CountryName string   `json:"countryName"`
	ZoneName    string   `json:"zoneName"`
	Access      []string `json:"access"`
}

type TypRecentCI struct {
	Zone    string `json:"zone"`
	History []struct {
		CarbonIntensity int    `json:"carbonIntensity"`
		Datetime        string `json:"datetime"`
		UpdatedAt       string `json:"updatedAt"`
		CreatedAt       string `json:"createdAt"`
	} `json:"history"`
}

type TypRecentPB struct {
	Zone    string `json:"zone"`
	History []struct {
		Datetime                  string                       `json:"datetime"`
		FossilFreePercentage      string                       `json:"fossilFreePercentage"`
		PowerConsumptionBreakdown TypPowerConsumptionBreakdown `json:"powerConsumptionBreakdown"`
		PowerConsumptionTotal     int                          `json:"powerConsumptionTotal"`
		PowerImportBreakdown      TypPowerImpExpBreakdown      `json:"powerImportBreakdown"`
		PowerImportTotal          int                          `json:"powerImportTotal"`
		PowerExportBreakdown      TypPowerImpExpBreakdown      `json:"powerExportBreakdown"`
		PowerExportTotal          int                          `json:"powerExportTotal"`
		PowerProductionBreakdown  TypPowerProductionBreakdown  `json:"powerProductionBreakdown"`
		PowerProductionTotal      int                          `json:"powerProductionTotal"`
		RenewablePercentage       int                          `json:"renewablePercentage"`
	} `json:"history"`
}
