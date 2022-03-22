package electricitymap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetZones(zoneKey string) (map[string]typZone, error) {
	url := "https://api.electricitymap.org/v3/zones"
	data := make(map[string]typZone)
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	fmt.Println("Getting Electricity Map Zones")
	err := HttpGet(url, &data, header, query)
	return data, err
}

func LiveCarbonIntensity(zoneKey string, zone string, lon string, lat string) (typCI, error) {
	url := "https://api.electricitymap.org/v3/carbon-intensity/latest"
	var data typCI
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}

	fmt.Println("Getting Electricity Map Live Carbon Intensity")
	err := HttpGet(url, &data, header, query)
	return data, err

}

func LivePowerBreakdown(zoneKey string, zone string, lon string, lat string) (typPB, error) {
	url := "https://api.electricitymap.org/v3/power-breakdown/latest"
	var data typPB
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}

	fmt.Println("Getting Electricity Map Live Power Breakdown")
	err := HttpGet(url, &data, header, query)
	return data, err

}

func RecentCarbonIntensity(zoneKey string, zone string, lon string, lat string) (typRecentCI, error) {
	url := "https://api.electricitymap.org/v3/carbon-intensity/history"
	var data typRecentCI
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}

	fmt.Println("Getting Electricity Map Recent Carbon Intensity")
	err := HttpGet(url, &data, header, query)
	return data, err

}

func RecentPowerBreakdown(zoneKey string, zone string, lon string, lat string) (typRecentPB, error) {
	url := "https://api.electricitymap.org/v3/power-consumption-breakdown/history"
	var data typRecentPB
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}

	fmt.Println("Getting Electricity Map Recent Power Breakdown")
	err := HttpGet(url, &data, header, query)
	return data, err

}

func PastCarbonIntensity(zoneKey string, zone string, lon string, lat string, datetime string, estimationFallback bool) (typCI, error) {
	url := "https://api.electricitymap.org/v3/carbon-intensity/past"
	var data typCI
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}
	query["datetime"] = datetime
	query["estimationFallback"] = strconv.FormatBool(estimationFallback)

	fmt.Println("Getting Electricity Map Past Carbon Intensity")
	err := HttpGet(url, &data, header, query)
	return data, err

}
func PastCarbonIntensityRange(zoneKey string, zone string, lon string, lat string, start string, end string, estimationFallback bool) (map[string][]typCI, error) {
	url := "https://api.electricitymap.org/v3/carbon-intensity/past-range"
	var data = make(map[string][]typCI)
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}
	query["start"] = start
	query["end"] = end
	query["estimationFallback"] = strconv.FormatBool(estimationFallback)

	fmt.Println("Getting Electricity Map Past Carbon Intensity Range")
	err := HttpGet(url, &data, header, query)
	return data, err

}

func PastPowerBreakdown(zoneKey string, zone string, lon string, lat string, datetime string, estimationFallback bool) (typPB, error) {
	url := "https://api.electricitymap.org/v3/power-breakdown/past"
	var data typPB
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}
	query["datetime"] = datetime
	query["estimationFallback"] = strconv.FormatBool(estimationFallback)

	fmt.Println("Getting Electricity Map Past Power Breakdown")
	err := HttpGet(url, &data, header, query)
	return data, err

}

func PastPowerBreakdownRange(zoneKey string, zone string, lon string, lat string, start string, end string, estimationFallback bool) (map[string][]typPB, error) {
	url := "https://api.electricitymap.org/v3/power-breakdown/past-range"
	var data = make(map[string][]typPB)
	header := make(map[string]string)
	query := make(map[string]string)

	header["auth-token"] = zoneKey

	query["zone"] = zone
	if lon != "" {
		query["lon"] = lon
	}
	if lat != "" {
		query["lat"] = lat
	}
	query["start"] = start
	query["end"] = end
	query["estimationFallback"] = strconv.FormatBool(estimationFallback)

	fmt.Println("Getting Electricity Map Past Power Breakdown Range")
	err := HttpGet(url, &data, header, query)
	return data, err

}

func HttpGet(url string, data interface{}, header map[string]string, query map[string]string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if hasError(err) {
		fmt.Println("http.NewRequest error")
		fmt.Println(err.Error())
		return err
	}

	// Add Headers
	for k := range header {
		fmt.Printf("Adding header %v:%v\n", k, header[k])
		req.Header.Add(k, header[k])
	}

	// Get URL Query String
	q := req.URL.Query()

	for k := range query {
		q.Add(k, query[k])
	}

	// Add query string to URL
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL)
	response, err := client.Do(req)
	if hasError(err) {
		fmt.Println("client.Do error")
		fmt.Println(err.Error())
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if hasError(err) {
		fmt.Println("ioutil.ReadAll error")
		fmt.Println(err.Error())
		return err
	}

	json.Unmarshal(responseData, &data)
	return nil //no error
}

func hasError(err error) bool {
	if err != nil {
		log.Fatal(err)
		return true
	}
	return false
}

type typCI struct {
	Zone            string `json:"zone"`
	CarbonIntensity int    `json:"carbonIntensity"`
	Datetime        string `json:"datetime"`
	UpdatedAt       string `json:"updatedAt"`
	CreatedAt       string `json:"createdAt"`
}

type typPB struct {
	Zone                      string                       `json:"zone"`
	Datetime                  string                       `json:"datetime"`
	PowerProductionBreakdown  typPowerProductionBreakdown  `json:"powerProductionBreakdown"`
	PowerProductionTotal      int                          `json:"powerProductionTotal"`
	PowerConsumptionBreakdown typPowerConsumptionBreakdown `json:"powerConsumptionBreakdown"`
	PowerConsumptionTotal     int                          `json:"powerConsumptionTotal"`
	PowerImportBreakdown      typPowerImpExpBreakdown      `json:"powerImportBreakdown"`
	PowerImportTotal          int                          `json:"powerImportTotal"`
	PowerExportBreakdown      typPowerImpExpBreakdown      `json:"powerExportBreakdown"`
	PowerExportTotal          int                          `json:"powerExportTotal"`
	FossilFreePercentage      int                          `json:"fossilFreePercentage"`
	RenewablePercentage       int                          `json:"renewablePercentage"`
	UpdatedAt                 string                       `json:"updatedAt"`
	CreatedAt                 string                       `json:"createdAt"`
}

type typPowerConsumptionBreakdown struct {
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

type typPowerImpExpBreakdown struct {
	DE     int `json:"DE"`
	DK_DK1 int //DK-DK1 `json:"DK_DK1"`
	SE     int `json:"SE"`
}

type typPowerProductionBreakdown struct {
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

type typZone struct {
	CountryName string   `json:"countryName"`
	ZoneName    string   `json:"zoneName"`
	Access      []string `json:"access"`
}

type typRecentCI struct {
	Zone    string `json:"zone"`
	History []struct {
		CarbonIntensity int    `json:"carbonIntensity"`
		Datetime        string `json:"datetime"`
		UpdatedAt       string `json:"updatedAt"`
		CreatedAt       string `json:"createdAt"`
	} `json:"history"`
}

type typRecentPB struct {
	Zone    string `json:"zone"`
	History []struct {
		Datetime                  string                       `json:"datetime"`
		FossilFreePercentage      string                       `json:"fossilFreePercentage"`
		PowerConsumptionBreakdown typPowerConsumptionBreakdown `json:"powerConsumptionBreakdown"`
		PowerConsumptionTotal     int                          `json:"powerConsumptionTotal"`
		PowerImportBreakdown      typPowerImpExpBreakdown      `json:"powerImportBreakdown"`
		PowerImportTotal          int                          `json:"powerImportTotal"`
		PowerExportBreakdown      typPowerImpExpBreakdown      `json:"powerExportBreakdown"`
		PowerExportTotal          int                          `json:"powerExportTotal"`
		PowerProductionBreakdown  typPowerProductionBreakdown  `json:"powerProductionBreakdown"`
		PowerProductionTotal      int                          `json:"powerProductionTotal"`
		RenewablePercentage       int                          `json:"renewablePercentage"`
	} `json:"history"`
}
