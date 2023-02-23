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
    fmt.Println("Filepath: ", jsonPath)
	var arm TypARM
	err := json.Unmarshal([]byte(file), &arm)
	if err != nil {
		fmt.Println(err.Error())
	}
	return arm
}

func armSummary(filename string) []TypSummary {
    fmt.Println("Reading Json!")
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
			(*summary)[resIndex].Count++
			s := (*summary)[resIndex]
			sizeExists, sizeIndex := isExistingSize(&s.Sizes, getResourceSize(resource))
			if !sizeExists {
				addArmSizeToRes(summary, resIndex, resource)
			} else {
				sz := (&s).Sizes[sizeIndex]
				locExists, locIndex := isExistingLocation(&sz.Details, getValue(resource.Location))
				if !locExists {
					addArmLocToSize(summary, resIndex, sizeIndex, resource)
				} else {
					fmt.Println(sz.Details)
					fmt.Println(locIndex)
					d := sz.Details[locIndex]
					if d.Location == getValue(resource.Location) {
						(*summary)[resIndex].Sizes[sizeIndex].Details[locIndex].Count++
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
		if s.Resource == resource.Type {
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
		if s.Size == size {
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
		if s.Location == location {
			exists = true
			index = n
		}
	}
	return exists, index
}

func defDetails(resource *TypResource) (dtl []TypSummaryDetails) {
	dtl = append(dtl, TypSummaryDetails{Location: getValue(resource.Location), Count: 1})
	return
}

func addArmResToSummary(summary *[]TypSummary, resource *TypResource) {
	dtl := defDetails(resource)
	size := getResourceSize(resource)
	var sizes []TypSizes
	sizes = append(sizes, TypSizes{Size: size, Details: dtl})
	sum := TypSummary{Resource: resource.Type, Sizes: sizes, Count: 1}
	*summary = append(*summary, sum)
}

func addArmSizeToRes(summary *[]TypSummary, resIndex int, resource *TypResource) {
	dtl := defDetails(resource)
	size := TypSizes{Size: getResourceSize(resource), Details: dtl}
	(*summary)[resIndex].Sizes = append((*summary)[resIndex].Sizes, size)
}

func addArmLocToSize(summary *[]TypSummary, resIndex int, sizeIndex int, resource *TypResource) {
	dtl := TypSummaryDetails{Location: getValue(resource.Location), Count: 1}
	(*summary)[resIndex].Sizes[sizeIndex].Details = append((*summary)[resIndex].Sizes[sizeIndex].Details, dtl)
}

func getResourceSize(resource *TypResource) string {
	var size string
	switch resource.Type {
	case "Microsoft.Compute/virtualMachines":
		size = getValue(resource.Properties.HardwareProfile.VmSize)
	default:
		size = getValue(resource.SKU.Name)
	}
	return size
}

func getValue(inputValue string) string {

	if strings.Contains(inputValue, "[parameters(") {
		p := strings.Split(inputValue, "'")
		return data.Parameters[p[1]].DefaultValue
	} else if strings.Contains(inputValue, "[variables(") {
		p := strings.Split(inputValue, "'")
		return data.Variables[p[1]]
	}

	return inputValue
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
