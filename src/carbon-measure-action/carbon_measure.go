package main

import (
	"fmt"
	"os"
)

func main() {
	githubNoticeMessage("Starting carbon measure action.")
	infraFileType := os.Getenv("file_type")
	// TODO: For terraform, we might need to accept a list of multiple files
	infraFileName := os.Getenv("file_location")
	electricityMapZoneKey := os.Getenv("zone_key")
	totalKwh := iterateOverFile(infraFileName, infraFileType)
	averageKwh := getCarbonIntensity(electricityMapZoneKey)
	// this comes from electricity map
	carbonIntensity := float64(averageKwh) * totalKwh
	gitHubOutputVariable("grams_carbon_equivalent_per_kwh", fmt.Sprint(averageKwh))
	gitHubOutputVariable("grams_emitted_over_24h", fmt.Sprint(carbonIntensity))
	githubNoticeMessage("Successfully ran carbon measure action.")
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
