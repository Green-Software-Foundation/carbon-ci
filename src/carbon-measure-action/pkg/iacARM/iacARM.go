package infraascode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadJSON(jsonPath string) TypARM { //file []byte
	file, _ := ioutil.ReadFile(jsonPath)
	var arm TypARM
	err := json.Unmarshal([]byte(file), &arm)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(arm)
	return arm
}

func SummarizeData(data *TypARM) []typSummary {
	var summary []typSummary
	for _, resource := range data.Resources {
		if resource.Type == "Microsoft.Resources/deployments" {
			// fmt.Println("Deployment type. Check sub Resources.")
			for _, depResource := range resource.Properties.Template.Resources {
				processSummary(&summary, &depResource)
			}
		} else {
			processSummary(&summary, &resource)
		}
	}
	// fmt.Printf("\n\nSUMMARY: \n%+v\n\n", summary)
	return summary
}

func processSummary(summary *[]typSummary, resource *TypResource) {
	// fmt.Printf("%+v\n", v)
	if len(*summary) == 0 {
		// fmt.Printf("Summary is empty. Adding resource %v with count 1.\n\n", resource.Type)
		addResourceToSummary(resource, summary)
	} else {
		resourceExists := false
		for i, s := range *summary {
			// fmt.Printf("if %v == %v\n", s.resource, resource.Type)
			if s.resource == resource.Type {
				locationExists := false
				for i2, d := range s.details {
					// fmt.Printf("if %v == %v\n", d.location, resource.Location)
					// fmt.Printf("Parameter value of %v is %v\n", resource.Location, getParameterValue(resource.Location, data))
					if d.location == resource.Location { // add count to existing location
						// fmt.Printf("Location exists. Adding 1 to count %v.\n\n", d.count)
						// fmt.Println(summary[i])
						(*summary)[i].details[i2].count++
						// fmt.Println(summary[i])
						locationExists = true
					}
				}
				if !locationExists { // add new location
					// fmt.Printf("Location %v does not exist. Adding with count 1.\n\n", resource.Location)
					det := typSummaryDetails{location: resource.Location, count: 1}
					(*summary)[i].details = append((*summary)[i].details, det)
				}
				resourceExists = true
			}
		}
		if !resourceExists {
			// fmt.Printf("Resource %v does not exist. Adding with count 1.\n\n", resource.Type)
			addResourceToSummary(resource, summary)
		}
	}
}

func addResourceToSummary(resource *TypResource, summary *[]typSummary) {
	var det []typSummaryDetails
	det = append(det, typSummaryDetails{location: resource.Location, count: 1})
	sum := typSummary{resource: resource.Type, details: det}
	*summary = append(*summary, sum)
}

func getParameterValue(param string, data *TypARM) string {
	p := strings.Split(param, "'")
	// fmt.Printf("Parameter on index 1 is %v\n", p[1])
	return data.Parameters[p[1]].DefaultValue

}

func PrintSummary(summary *[]typSummary, data *TypARM) {

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("*****************************************")
	fmt.Println("*** L I S T   O F   R E S O U R C E S ***")
	fmt.Println("*****************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println()

	total := 0
	for _, s := range *summary {
		count := 0
		fmt.Println(s.resource)
		for _, d := range s.details {
			fmt.Printf("- %v in %v\n", d.count, getParameterValue(d.location, data))
			count += d.count
			total += d.count
		}
		fmt.Printf("TOTAL: %v\n", count)
		fmt.Println()
	}
	fmt.Printf("Total number of resources: %v\n\n", total)
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
