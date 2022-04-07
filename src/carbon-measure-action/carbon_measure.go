package main

import (
	"fmt"
	EM "main/pkg/electricitymap"
	iac "main/pkg/infraascode"

	//elec "main/pkg/electricitymap"
	"os"
)

//SgithubNoticeMessage("Starting carbon measure action.")
// wattTimeUser := os.Getenv("WATT_TIME_USER")
// wattTimePass := os.Getenv("WATT_TIME_PASS")
func main() {
	//var grandaverageKwh int
	//("Starting carbon measure action.")
	//os.Setenv("IACType", "arm")
	//os.Setenv("IACTemplateFile", "vm")
	//os.Setenv("ELECTRICITY_MAP_AUTH_TOKEN", "3bhtgXSayVvgmuwEHry6zYYr")
	infraFileType := os.Getenv("IACType")
	infraFileName := os.Getenv("IACTemplateFile")
	electricityMapZoneKey := os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")

	// TODO: For terraform, we might need to accept a list of multiple files

	//sumary := iac.GetIACSummary(iac.TypIACQuery{Filetype: infraFileType, Filename: infraFileName})

	em := EM.New(electricityMapZoneKey)
	em.LiveCarbonIntensity(EM.TypAPIParams{Zone: "US-CAL-CISCO"})

	totalKwh := iterateOverFile(infraFileName, infraFileType)
	averageKwh := getCarbonIntensity(electricityMapZoneKey)
	// this comes from electricity map
	carbonIntensity := float64(averageKwh) * totalKwh
	fmt.Println("grams_carbon_equivalent_per_kwh", averageKwh)
	fmt.Println("grams_emitted_over_24h", carbonIntensity)
	fmt.Println("Successfully ran carbon measure action.")

	//gitHubOutputVariable("grams_carbon_equivalent_per_kwh", fmt.Sprint(averageKwh))
	//gitHubOutputVariable("grams_emitted_over_24h", fmt.Sprint(carbonIntensity))
	//githubNoticeMessage("Successfully ran carbon measure action.")
}

func getCarbonIntensity(zoneKey string) int {

	//cc := new(electmap.TypAPIParams)
	//cc.Zone = "US-CAL-CISO"
	//ccmap, _ := electmap.New(zoneKey).RecentCarbonIntensity(*cc)
	//return ccmap.History[len(ccmap.History)-1].CarbonIntensity //200
	// var x []TypResource
	// for s := range x {

	// 	for resourcedetail := range x[s].Location {

	// 		//	carbontotal = carbontotal + 1
	// 		//	totalwatt = totalwatt + 1
	// 		//	procesorcount = procesorcount + 1
	// 	}

	// }
	// TODO: Get the carbon intensity over 24 hours rather than just the current intensity
	return 200
}

func getKwhForComponent(componentName string) float64 {
	return 2.6
}

func iterateOverFile(fileName string, infraFileType string) float64 {
	// TODO: Get kwh for each component and return a summed float
	// TODO: Call a different iterator depending on if it is arm, bicep, terraform, pulumi, etc
	//var summary []TypSummary
	println(fileName)
	println(infraFileType)
	var c int
	summary := iac.GetIACSummary(iac.TypIACQuery{Filetype: infraFileType, Filename: fileName})
	for _, ts := range summary {
		c = ts.Count
	}

	return getKwhForComponent("component1") * float64(c)
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
