package infraascode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var data TypARM

func readJSON(jsonPath string) TypARM {
	file, _ := ioutil.ReadFile(jsonPath)
	var arm TypARM
	err := json.Unmarshal([]byte(file), &arm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return arm
}

func armSummary(filename string) []TypSummary {
	data = readJSON(filename)
	var summary []TypSummary
	for _, resource := range data.Resources {
		if resource.Type == "Microsoft.Resources/deployments" {
			for _, depResource := range resource.Properties.Template.Resources {
				processArmSummary(&summary, &depResource)
			}
		} else {
			processArmSummary(&summary, &resource)
		}
	}
	return summary
}

func processArmSummary(summary *[]TypSummary, resource *TypResource) {

	if len(*summary) == 0 {
		addArmResToSummary(summary, resource)
	} else {
		resourceExists, resIndex := isExistingResource(summary, resource)
		if !resourceExists {
			addArmResToSummary(summary, resource)
		} else {
			s := (*summary)[resIndex]
			sizeExists, sizeIndex := isExistingSize(&s.sizes, getResourceSize(resource))
			if !sizeExists {
				addArmSizeToRes(summary, resIndex, resource)
			} else {
				sz := (&s).sizes[sizeIndex]
				locExists, locIndex := isExistingLocation(&sz.details, getParameterValue(resource.Location))
				if !locExists {
					addArmLocToSize(summary, resIndex, sizeIndex, resource)
				} else {
					fmt.Println(sz.details)
					fmt.Println(locIndex)
					d := sz.details[locIndex]
					if d.location == getParameterValue(resource.Location) {
						(*summary)[resIndex].sizes[sizeIndex].details[locIndex].count++
					}
				}
			}
		}
	}
}

func isExistingResource(summary *[]TypSummary, resource *TypResource) (bool, int) {
	exists := false
	index := 0
	for n, s := range *summary {
		if s.resource == resource.Type {
			exists = true
			index = n
		}
	}
	return exists, index
}

func isExistingSize(sizes *[]TypSizes, size string) (bool, int) {
	exists := false
	index := 0
	for n, s := range *sizes {
		if s.size == size {
			exists = true
			index = n
		}
	}
	return exists, index
}

func isExistingLocation(details *[]TypSummaryDetails, location string) (bool, int) {
	exists := false
	index := 0
	for n, s := range *details {
		if s.location == location {
			exists = true
			index = n
		}
	}
	return exists, index
}

func defDetails(resource *TypResource) (dtl []TypSummaryDetails) {
	dtl = append(dtl, TypSummaryDetails{location: getParameterValue(resource.Location), count: 1})
	return
}

func addArmResToSummary(summary *[]TypSummary, resource *TypResource) {
	dtl := defDetails(resource)
	size := getResourceSize(resource)
	var sizes []TypSizes
	sizes = append(sizes, TypSizes{size: size, details: dtl})
	sum := TypSummary{resource: resource.Type, sizes: sizes}
	*summary = append(*summary, sum)
}

func addArmSizeToRes(summary *[]TypSummary, resIndex int, resource *TypResource) {
	dtl := defDetails(resource)
	size := TypSizes{size: getResourceSize(resource), details: dtl}
	(*summary)[resIndex].sizes = append((*summary)[resIndex].sizes, size)
}

func addArmLocToSize(summary *[]TypSummary, resIndex int, sizeIndex int, resource *TypResource) {
	dtl := TypSummaryDetails{location: getParameterValue(resource.Location), count: 1}
	(*summary)[resIndex].sizes[sizeIndex].details = append((*summary)[resIndex].sizes[sizeIndex].details, dtl)
}

func getResourceSize(resource *TypResource) string {
	var size string
	switch resource.Type {
	case "Microsoft.Compute/virtualMachines":
		size = resource.Properties.HardwareProfile.VmSize
	default:
		size = resource.SKU.Name
	}
	return size
}

func getParameterValue(param string) string {
	p := strings.Split(param, "'")
	return data.Parameters[p[1]].DefaultValue
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
		HardwareProfile struct {
			VmSize string `json:"vmSize"`
		} `json:"hardwareProfile"`
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
