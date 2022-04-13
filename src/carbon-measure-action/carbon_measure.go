package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	iac "main/pkg/infraascode"
	pa "main/pkg/poweradapter"
	"os"
	"strconv"
	"strings"
)

func main() {

	infraFileType := os.Getenv("IACType")
	infraFileName := os.Getenv("IACTemplateFile")
	electricityMapZoneKey := os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")
	var averageKwh int
	var count int
	// TODO: For terraform, we might need to accept a list of multiple files
	var param pa.TypCarbonQueryParams
	param.IacLocation = "eastus" // "US-CAL-CISCO"
	param.IacProvider = "azure"
	param.CarbonRateProvider = "electricitymap"
	param.ElectricityMapZoneKey = electricityMapZoneKey
	param.WattTimeUser = "dtejares"
	param.WattTimePass = "u8Go!T1BOQU+"

	sumary := iac.GetIACSummary(iac.TypIACQuery{Filetype: infraFileType, Filename: infraFileName})

	for _, ts := range sumary {
		count = ts.Count
		Sizes := ts.Sizes
		if ts.Resource == "Microsoft.Compute/virtualMachines" {
			for _, S := range Sizes {
				for _, D := range S.Details {
					fmt.Println(D.Location)
					param.IacLocation = D.Location
					averageKwh = averageKwh + getCarbonIntensity(param)
				}
			}
		}
	}
	averageKwh = averageKwh / count
	totalKwh := iterateOverFile(infraFileName, infraFileType)

	fmt.Println("averageKwh")
	fmt.Println(averageKwh)

	carbonIntensity := float64(averageKwh) * totalKwh
	fmt.Println("grams_carbon_equivalent_per_kwh", averageKwh)
	fmt.Println("grams_emitted_over_24h", carbonIntensity)
	fmt.Println("Successfully ran carbon measure action.")

	//gitHubOutputVariable("grams_carbon_equivalent_per_kwh", fmt.Sprint(averageKwh))
	//gitHubOutputVariable("grams_emitted_over_24h", fmt.Sprint(carbonIntensity))
	//githubNoticeMessage("Successfully ran carbon measure action.")
}

func getCarbonIntensity(param pa.TypCarbonQueryParams) int {

	x := pa.LiveCarbonIntensity(param)
	if x.LiveCarbonIntensity < 1 {
		x.LiveCarbonIntensity = 200 //  dummy
	}

	return x.LiveCarbonIntensity //200
}

func getKwhForComponent(qry TypCloudResourceQuery) float64 {

	qry.SizeName = "Standard_A2_V2"
	qry.Type = "Microsoft.Compute/virtualMachines"
	qry.Provider = "azure"

	return GetWattage(qry) //2.6
}

func iterateOverFile(fileName string, infraFileType string) float64 {
	// TODO: Get kwh for each component and return a summed float
	// TODO: Call a different iterator depending on if it is arm, bicep, terraform, pulumi, etc
	var qry TypCloudResourceQuery
	var TotalKwh float64
	summary := iac.GetIACSummary(iac.TypIACQuery{Filetype: infraFileType, Filename: fileName})

	for _, ts := range summary {
		//dummy
		qry.SizeName = "Standard_A2_V2"
		qry.Type = "Microsoft.Compute/virtualMachines"
		qry.Provider = "azure"
		Sizes := ts.Sizes
		if ts.Resource == "Microsoft.Compute/virtualMachines" {
			for _, S := range Sizes {
				qry.SizeName = S.Size
				for _, D := range S.Details {

					qry.SizeName = D.Location
					TotalKwh = TotalKwh + getKwhForComponent(qry)
				}

			}
		}
	}

	return TotalKwh // getKwhForComponent(qry) * float64(c)
}

func readJSON(jsonPath string) []TypCloudResources {
	file, _ := ioutil.ReadFile(jsonPath)
	var cloudLoc []TypCloudResources
	err := json.Unmarshal([]byte(file), &cloudLoc)
	if err != nil {
		fmt.Println(err.Error())
	}
	return cloudLoc
}

func GetWattage(qry TypCloudResourceQuery) (watt float64) {
	cloudLoc := readJSON("references/resources.json")
	for _, c := range cloudLoc {
		if strings.ToLower(c.Cloud) == strings.ToLower(qry.Provider) {

			for _, l := range c.Resouce {
				if strings.ToLower(l.Type) == strings.ToLower(qry.Type) {

					for _, s := range l.Sizes {
						if strings.ToLower(s.Name) == strings.ToLower(qry.SizeName) {
							watt, _ = strconv.ParseFloat(s.Wattage, 64) //watt, _ = strconv.Atoi(s.wattage) //strconv.ParseFloat( s.wattage)

						}
					}
				}
			}
		}
	}
	return
}

type TypARM struct {
	Resources  []TypResource           `json:"resources"`
	Parameters map[string]TypParameter `json:"parameters"`
	Variables  map[string]string       `json:"variables"`
}

type TypResource struct {
	Type       string         `json:"type"`
	Location   string         `json:"location"`
	SKU        typResourceSKU `json:"sku"`
	Properties struct {
		Template struct {
			Resources []TypResource `json:"resources"`
		} `json:"template"`
	} `json:"properties"`
}

type typResourceSKU struct {
	Name      string   `json:"name"`
	Tier      string   `json:"tier"`
	Size      string   `json:"size"`
	Family    string   `json:"family"`
	Capacity  int      `json:"capacity"`
	Locations []string `json:"locations"`
}

type TypParameter struct {
	Type          string      `json:"type"`
	DefaultValue  string      `json:"defaultValue"`
	AllowedValues []string    `json:"allowedValues"`
	MinValue      int         `json:"minValue"`
	MaxValue      int         `json:"maxValue"`
	MinLength     int         `json:"minLength"`
	MaxLength     int         `json:"maxLength"`
	Metadata      typMetadata `json:"metadata"`
}

type typVariable struct {
}

type typMetadata struct {
	Description string `json:"description"`
}

type typSummary struct {
	resource string
	details  []typSummaryDetails
}

type typSummaryDetails struct {
	location string
	count    int
}

type TypSummary struct {
	resource string
	sizes    []TypSizes
}

type TypSizes struct {
	size    string
	details []TypSummaryDetails
}

type TypSummaryDetails struct {
	location string
	count    int
}

type TypCloudResourceQuery struct {
	Provider string
	Location string
	SizeName string
	Type     string
}
type TypCloudResources struct {
	Cloud   string       `json:"cloud"`
	Resouce []TypResouce `json:"resources"`
}

type TypResouce struct {
	Type  string    `json:"type"`
	Sizes []TypSize `json:"sizes"`
}
type TypSize struct {
	Name    string `json:"Name"`
	Wattage string `json:"wattage"`
}
