package main

import (
	"fmt"
	iac "main/pkg/iacARM"
	"os"
)

func main() {
	githubNoticeMessage("Starting carbon measure action.")
	infraFileType := os.Getenv("IACType")
	infraFileName := os.Getenv("IACFile")

	// TODO: For terraform, we might need to accept a list of multiple files

	if infraFileType == "arm" {
		processARMData(infraFileName)
	}

	electricityMapZoneKey := os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")
	totalKwh := iterateOverFile(infraFileName, infraFileType)
	averageKwh := getCarbonIntensity(electricityMapZoneKey)
	// this comes from electricity map
	carbonIntensity := float64(averageKwh) * totalKwh
	gitHubOutputVariable("grams_carbon_equivalent_per_kwh", fmt.Sprint(averageKwh))
	gitHubOutputVariable("grams_emitted_over_24h", fmt.Sprint(carbonIntensity))
	githubNoticeMessage("Successfully ran carbon measure action.")
}

func processARMData(file string) {
	data := iac.ReadJSON(file)

	// Summarize ARM JSON file to resource and location
	// TODO: Need to retrieve Variables and Parameters values inside resource type "Microsoft.Resources/deployments"
	summary := iac.SummarizeData(&data)

	// Print out summarized ARM data
	iac.PrintSummary(&summary, &data)
}

func getCarbonIntensity(zoneKey string) int {
	return 200
	// TODO: Get the carbon intensity over 24 hours rather than just the current intensity
}

func getKwhForComponent(componentName string) float64 {
	return 2.6
}

func iterateOverFile(fileName string, infraFileType string) float64 {
	// TODO: Get kwh for each component and return a summed float
	// TODO: Call a different iterator depending on if it is arm, bicep, terraform, pulumi, etc
	return getKwhForComponent("component1") + getKwhForComponent("component2")
}
