package infraascode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func readPulumiJSON(stackJson string, previewJson string) (typStack, typPreview) {
	stackFile, _ := ioutil.ReadFile(stackJson)
	previewFile, _ := ioutil.ReadFile(previewJson)
	var stack typStack
	var preview typPreview
	stackErr := json.Unmarshal([]byte(stackFile), &stack)
	previewErr := json.Unmarshal([]byte(previewFile), &preview)

	if stackErr != nil {
		fmt.Println(stackErr.Error())
	}

	if previewErr != nil {
		fmt.Println(previewErr.Error())
	}

	return stack, preview
}

func processStackData(stack typStack) []typCloudResource {
	var resourceList []typCloudResource

	for _, resource := range stack.Deployment.Resources {
		if strings.Split(resource.Type, ":")[0] != "pulumi" && !strings.Contains(resource.Type, "resourceGroup") {
			cloudRsc := typCloudResource{
				Location: resource.Inputs.Location,
				Name:     resource.Inputs.Name,
				SKU:      resource.Inputs.SKU,
				Type:     resource.Type,
			}
			resourceList = append(resourceList, cloudRsc)
		}
	}
	return resourceList
}

func processPreviewData(preview typPreview, resourceList *[]typCloudResource) {
	for _, step := range preview.Steps {
		if strings.Split(step.State.Type, ":")[0] != "pulumi" && !strings.Contains(step.State.Type, "resourceGroup") {
			if step.Operation == "create" {
				cloudRsc := typCloudResource{
					Location: step.State.Inputs.Location,
					Name:     step.State.Inputs.Name,
					SKU:      step.State.Inputs.SKU,
					Type:     step.State.Type,
				}
				(*resourceList) = append((*resourceList), cloudRsc)
			} else if step.Operation == "update" {
				for i := range *resourceList {
					if (*resourceList)[i].Name == step.State.Inputs.Name {
						(*resourceList)[i].SKU = step.State.Inputs.SKU
						(*resourceList)[i].Location = step.State.Inputs.Location
						break
					}
				}
			}
		}
	}
}

func getCloudProvider(stack *typStack, preview *typPreview) string {

	// Check for cloud provider on stack data
	for _, resource := range stack.Deployment.Resources {
		if strings.Contains(resource.Type, "pulumi:providers") {
			parts := strings.Split(resource.Type, ":")
			return parts[2]
		}
	}

	// If there is no stack data, look for cloud provider on preview data
	for _, resource := range preview.Steps {
		if !strings.Contains(resource.State.Type, "pulumi:pulumi") {
			parts := strings.Split(resource.State.Type, ":")
			return parts[0]
		}
	}

	// If no stack data or preview data is available, cloud provider cannot be determined
	return ""
}

func openResourceTypeRefenrece() []typResourceTypesReference {
	ref, _ := ioutil.ReadFile("./references/resourceTypes.json")

	var resourceTypeRef []typResourceTypesReference

	err := json.Unmarshal([]byte(ref), &resourceTypeRef)

	if err != nil {
		fmt.Println(err.Error())
	}

	return resourceTypeRef
}

func getCloudProviderResourceType(pulumiResourceType string, ref *[]typResourceTypesReference) string {
	for _, resourceType := range *ref {
		if resourceType.PulumiResourceType == pulumiResourceType {
			return resourceType.CloudResourceType
		}
	}
	//return pulumi type if not found on the reference
	return pulumiResourceType
}

func pulumiSummary(stackJson string, previewJson string) []TypSummary {
	var resourceList []typCloudResource
	var summary []TypSummary

	stack, preview := readPulumiJSON(stackJson, previewJson)
	resourceList = processStackData(stack)
	processPreviewData(preview, &resourceList)

	cloudProvider := getCloudProvider(&stack, &preview)
	fmt.Printf("Cloud provider is %v.\n", cloudProvider)

	resourceTypeRef := openResourceTypeRefenrece()

	for _, resource := range resourceList {
		resource.Type = getCloudProviderResourceType(resource.Type, &resourceTypeRef)
		resourceExists, resourceIndex := isExistingPulumiResource(&summary, resource.Type)
		if resourceExists {
			sizeExists, sizeIndex := isExistingSize(&summary[resourceIndex].Sizes, resource.SKU)
			if sizeExists {
				locationExists, locationIndex := isExistingLocation(&summary[resourceIndex].Sizes[sizeIndex].Details, resource.Location)
				if locationExists {
					summary[resourceIndex].Sizes[sizeIndex].Details[locationIndex].Count += 1
				} else {
					summary[resourceIndex].Sizes[sizeIndex].Details = append(summary[resourceIndex].Sizes[sizeIndex].Details, defineTypSummaryDetails(&resource))
				}
			} else {
				summary[resourceIndex].Sizes = append(summary[resourceIndex].Sizes, defineTypSize(&resource))
			}
		} else {
			summary = append(summary, defineTypSummary(&resource))
		}
	}
	return summary
}

func defineTypSummaryDetails(resource *typCloudResource) TypSummaryDetails {
	return TypSummaryDetails{
		Location: (*resource).Location,
		Count:    1,
	}
}

func defineTypSize(resource *typCloudResource) TypSizes {
	return TypSizes{
		Size:    (*resource).SKU,
		Details: []TypSummaryDetails{defineTypSummaryDetails(resource)},
	}
}

func defineTypSummary(resource *typCloudResource) TypSummary {
	return TypSummary{
		Resource: (*resource).Type,
		Sizes:    []TypSizes{defineTypSize(resource)},
	}
}

func isExistingPulumiResource(summary *[]TypSummary, resource string) (bool, int) {
	exists := false
	index := 0
	for n, s := range *summary {
		if s.Resource == resource {
			exists = true
			index = n
		}
	}
	return exists, index
}

type typStack struct {
	Deployment struct {
		Resources []typResource `json:"resources"`
	} `json:"deployment"`
}

type typResource struct {
	Type   string    `json:"type"`
	Inputs typInputs `json:"inputs"`
}

type typInputs struct {
	Id            string `json:"id"`
	Location      string `json:"location"`
	SKU           string `json:"vmSize"`
	Name          string `json:"name"`
	StorageOsDisk struct {
		ManagedDiskType string `json:"managedDiskType"`
	} `json:"storageOsDisk"`
}

type typPreview struct {
	Steps []typStep `json:"steps"`
}

type typStep struct {
	Operation string `json:"op"`
	State     struct {
		Type   string    `json:"type"`
		Inputs typInputs `json:"inputs"`
	} `json:"newState"`
}

type typCloudResource struct {
	Name     string
	Type     string
	SKU      string
	Location string
}

type typResourceTypesReference struct {
	CloudResourceType  string `json:"cloudResourceType"`
	PulumiResourceType string `json:"pulumiResourceType"`
}
