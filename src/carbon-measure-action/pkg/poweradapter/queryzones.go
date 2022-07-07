package poweradapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func readJSON(jsonPath string) []TypCloudLocations {
	file, _ := ioutil.ReadFile(jsonPath)
	var cloudLoc []TypCloudLocations
	err := json.Unmarshal([]byte(file), &cloudLoc)
	if err != nil {
		fmt.Println(err.Error())
	}
	return cloudLoc
}

func GetLocation(qry TypCloudLocationQuery) (zone string) {
	cloudLoc := readJSON("references/locations.json")
	for _, c := range cloudLoc {
		if strings.ToLower(c.Cloud) == strings.ToLower(qry.Provider) {
			for _, l := range c.Locations {
				if strings.ToLower(l.AzureRegion) == strings.ToLower(qry.Location) {
					switch strings.ToLower(qry.Powerprovider) {
					case "electricitymap":
						zone = l.ElectricitymapZone
					case "watttime":
						zone = l.Watttime
					}
				}
			}
		}
	}
	return
}

type TypCloudLocationQuery struct {
	Provider      string
	Location      string
	Powerprovider string
}

type TypCloudLocations struct {
	Cloud     string         `json:"cloud"`
	Locations []TypLocations `json:"locations"`
}

type TypLocations struct {
	AzureRegion        string `json:"azureRegion"`
	ElectricitymapZone string `json:"electricitymapZone"`
	Watttime           string `json:"watttime"`
}
