package main

import (
	"fmt"
	"os"
)

func main() {
	infrastructureAsCodeFileType := os.Getenv("file_type")
	// TODO: For terraform, we might need to accept a list of multiple files
	infrastructureAsCodeFileName := os.Getenv("file_location")
	infrastructureAsCodeZoneKey := os.Getenv("zone_key")
	fmt.Printf("Unused: %v %v %v\n", infrastructureAsCodeFileName, infrastructureAsCodeFileType, infrastructureAsCodeZoneKey)
	// this comes from electricity map
	gitHubOutputVariable("grams_carbon_equivalent_per_kwh", "200")
	gitHubOutputVariable("grams_over_24h", "4800")
	githubNoticeMessage("Successfully ran carbon measure action.")
}
