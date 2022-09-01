package infraascode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var terraformData TypTerraform

func terraformSummary(filename string) []TypSummary {
	var summary []TypSummary
	var tfResources []TypTfResource

	fmt.Println("Reading terraform json")
	terraformData = readTerraformJSON(filename)
	fmt.Println("Terraform json read successfully")

	resourceTypeRef := openTfResourceTypeReference()

	// First, process the main module
	for _, resource := range terraformData.PlannedValues.RootModule.Resources {
		resource.Type = getCloudProviderResourceTypeFromTfResourceType(resource.Type, &resourceTypeRef)
		tfResources = append(tfResources, resource)
	}

	// Then, process the child modules
	for _, childModule := range terraformData.PlannedValues.RootModule.ChildModules {
		for _, resource := range childModule.Resources {
			resource.Type = getCloudProviderResourceTypeFromTfResourceType(resource.Type, &resourceTypeRef)
			tfResources = append(tfResources, resource)
		}
	}

	// Now, get together all the resources and change the resource type. Then add them to the summary
	for _, resource := range tfResources {
		resource.Type = getCloudProviderResourceTypeFromTfResourceType(resource.Type, &resourceTypeRef)

		processTfResourceIntoSummary(&resource, &summary)
	}

	return summary
}

func processTfResourceIntoSummary(resource *TypTfResource, summary *[]TypSummary) {
	if len(*summary) == 0 {
		// At the beginning, when the summary is still empty, just add the terraform resource
		addTfResToSummary(summary, resource)
	} else {
		// Get if the resource exists and its position
		resourceExists, resIndex := isExistingTfResource(summary, resource)
		if !resourceExists {
			// This is a new resource, add it to the summary
			addTfResToSummary(summary, resource)
		} else {
			// This is not a new resource. Increase the count for the existing resource
			(*summary)[resIndex].Count++

			// Get the item in the summary
			s := (*summary)[resIndex]

			// Get if the size exists and its position
			sizeExists, sizeIndex := isExistingSize(&s.Sizes, getTfResourceSize(resource))

			if !sizeExists {
				// This is a new size, add it to the summary
				addTfSizeToRes(summary, resIndex, resource)
			} else {
				// This is not a new size. Get the item in the size
				sz := (&s).Sizes[sizeIndex]

				// Get if the location exists and its position
				locExists, locIndex := isExistingLocation(&sz.Details, resource.Values.Location)

				if !locExists {
					// This is a new location, add it to the size
					addTfLocToSize(summary, resIndex, sizeIndex, resource)
				} else {
					// This is not a new location. Increase its count
					(*summary)[resIndex].Sizes[sizeIndex].Details[locIndex].Count++
				}
			}
		}
	}
}

func addTfResToSummary(summary *[]TypSummary, resource *TypTfResource) {
	dtl := defineDetailsTfResource(resource)
	size := getTfResourceSize(resource)
	var sizes []TypSizes
	sizes = append(sizes, TypSizes{Size: size, Details: dtl})
	sum := TypSummary{Resource: resource.Type, Sizes: sizes, Count: 1}
	*summary = append(*summary, sum)
}

func defineDetailsTfResource(resource *TypTfResource) (dtl []TypSummaryDetails) {
	dtl = append(dtl, TypSummaryDetails{Location: getValue(resource.Values.Location), Count: 1})
	return
}

func getTfResourceSize(resource *TypTfResource) string {
	var size string
	switch resource.Type {
	case "Microsoft.Compute/virtualMachines":
		if len(resource.Values.Size) > 0 {
			// Use Size, because it's defined
			size = resource.Values.Size
		} else if len(resource.Values.VMSize) > 0 {
			// Use VMSize otherwise
			size = resource.Values.VMSize
		}
	default:
		size = getValue(resource.Values.SKU)
	}

	return size
}

func isExistingTfResource(summary *[]TypSummary, resource *TypTfResource) (bool, int) {
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

// Open the reference file used to map terraform resource type to the cloud provider resource type
func openTfResourceTypeReference() []TypTfResourceTypesReference {
	ref, _ := ioutil.ReadFile("./references/resourceTypes.json")

	var resourceTypeRef []TypTfResourceTypesReference

	err := json.Unmarshal([]byte(ref), &resourceTypeRef)

	if err != nil {
		fmt.Println(err.Error())
	}

	return resourceTypeRef
}

func addTfSizeToRes(summary *[]TypSummary, resIndex int, resource *TypTfResource) {
	dtl := defineDetailsTfResource(resource)
	size := TypSizes{Size: getTfResourceSize(resource), Details: dtl}
	(*summary)[resIndex].Sizes = append((*summary)[resIndex].Sizes, size)
}

func addTfLocToSize(summary *[]TypSummary, resIndex int, sizeIndex int, resource *TypTfResource) {
	dtl := TypSummaryDetails{Location: resource.Values.Location, Count: 1}
	(*summary)[resIndex].Sizes[sizeIndex].Details = append((*summary)[resIndex].Sizes[sizeIndex].Details, dtl)
}

func readTerraformJSON(jsonPath string) TypTerraform {
	file, _ := ioutil.ReadFile(jsonPath)
	var terraform TypTerraform
	err := json.Unmarshal([]byte(file), &terraform)
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot read the terraform json.")
	}
	return terraform
}

// Returns the corresponding cloud provider resource type given the terraform resource type
func getCloudProviderResourceTypeFromTfResourceType(terraformResourceType string, ref *[]TypTfResourceTypesReference) string {
	for _, resourceType := range *ref {
		if resourceType.TerraformResourceType == terraformResourceType {
			return resourceType.CloudResourceType
		}
	}

	//return terraform type if not found on the reference
	return terraformResourceType
}

type TypTerraform struct {
	PlannedValues TypTfPlannedValues `json:"planned_values"`
}

type TypTfPlannedValues struct {
	RootModule struct {
		Resources    []TypTfResource `json:"resources"`
		ChildModules []struct {
			Resources []TypTfResource `json:"resources"`
		} `json:"child_modules"`
	} `json:"root_module"`
}

type TypTfResource struct {
	Type   string      `json:"type"`
	Values TypTfValues `json:"values"`
}

type TypTfValues struct {
	Size     string `json:"size"`
	VMSize   string `json:"vm_size"`
	Location string `json:"location"`
	SKU      string `json:"sku"`
}

type TypTfResourceTypesReference struct {
	CloudResourceType     string `json:"cloudResourceType"`
	TerraformResourceType string `json:"terraformResourceType"`
}
