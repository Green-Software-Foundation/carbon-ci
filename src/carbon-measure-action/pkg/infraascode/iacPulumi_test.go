package infraascode

import (
	"testing"
)

func TestReturnResourceListFromStackData(t *testing.T) {
	resource := typResource{
		Type: "VM",
		Inputs: typInputs{
			Id:       "resource_id",
			Location: "southeastasia",
			SKU:      "vm_size",
			Name:     "resourcename",
			StorageOsDisk: struct {
				ManagedDiskType string "json:\"managedDiskType\""
			}{
				ManagedDiskType: "disktype",
			},
		},
	}
	stackData := typStack{
		Deployment: struct {
			Resources []typResource "json:\"resources\""
		}{
			Resources: []typResource{resource},
		},
	}

	result := processStackData(stackData)

	if result[0].Location != stackData.Deployment.Resources[0].Inputs.Location {
		t.Errorf("LOCATION: Output %q not equal to expected %q", result[0].Location, stackData.Deployment.Resources[0].Inputs.Location)
	} else if result[0].Name != stackData.Deployment.Resources[0].Inputs.Name {
		t.Errorf("NAME: Output %q not equal to expected %q", result[0].Name, stackData.Deployment.Resources[0].Inputs.Name)
	} else if result[0].SKU != stackData.Deployment.Resources[0].Inputs.SKU {
		t.Errorf("SKU: Output %q not equal to expected %q", result[0].SKU, stackData.Deployment.Resources[0].Inputs.SKU)
	} else if result[0].Type != stackData.Deployment.Resources[0].Type {
		t.Errorf("TYPE: Output %q not equal to expected %q", result[0].Type, stackData.Deployment.Resources[0].Type)
	}
}

func TestAddResourceToResourceListFromPreviewData(t *testing.T) {
	resource := typCloudResource{
		Name:     "vm1",
		Type:     "VM",
		SKU:      "vm_size",
		Location: "asia",
	}

	resourceList := []typCloudResource{resource}

	step := typStep{
		Operation: "create",
		State: struct {
			Type   string    "json:\"type\""
			Inputs typInputs "json:\"inputs\""
		}{
			Type: "Storage",
			Inputs: typInputs{
				Id:       "someid",
				Location: "asia",
				SKU:      "",
				Name:     "mystorage",
			},
		},
	}
	previewData := typPreview{
		Steps: []typStep{step},
	}

	processPreviewData(previewData, &resourceList)

	if len(resourceList) != 2 {
		t.Errorf("Length %v not equal to expected length which is 2", len(resourceList))
	}
}

func TestUpdateResourceWithDataFromPreview(t *testing.T) {
	resource := typCloudResource{
		Name:     "vm1",
		Type:     "VM",
		SKU:      "vm_size",
		Location: "asia",
	}

	resourceList := []typCloudResource{resource}

	step := typStep{
		Operation: "update",
		State: struct {
			Type   string    "json:\"type\""
			Inputs typInputs "json:\"inputs\""
		}{
			Type: "VM",
			Inputs: typInputs{
				Id:       "someid",
				Location: "asia",
				SKU:      "vm_size1",
				Name:     "vm1",
			},
		},
	}
	previewData := typPreview{
		Steps: []typStep{step},
	}

	processPreviewData(previewData, &resourceList)

	if len(resourceList) != 1 {
		t.Errorf("Length %v not equal to expected length which is 1", len(resourceList))
	} else if resourceList[0].SKU != previewData.Steps[0].State.Inputs.SKU {
		t.Errorf("Output %v not equal to expected %v", resourceList[0].SKU, previewData.Steps[0].State.Inputs.SKU)
	}
}

func TestGetCloudProviderFromStackData(t *testing.T) {
	stackData := typStack{
		Deployment: struct {
			Resources []typResource "json:\"resources\""
		}{
			Resources: []typResource{
				{
					Type: "pulumi:providers:azure",
				},
			},
		},
	}

	previewData := typPreview{}

	result := getCloudProvider(&stackData, &previewData)

	if result != "azure" {
		t.Errorf("Output %v is not equal to expected value which is azure", result)
	}
}

func TestGetCloudProviderFromPreviewData(t *testing.T) {
	stackData := typStack{}

	previewData := typPreview{
		Steps: []typStep{
			{
				State: struct {
					Type   string    "json:\"type\""
					Inputs typInputs "json:\"inputs\""
				}{
					Type: "azure:compute:virtualmachine",
				},
			},
		},
	}

	result := getCloudProvider(&stackData, &previewData)

	if result != "azure" {
		t.Errorf("Output %v is not equal to expected value which is azure", result)
	}
}

func TestCloudProviderNotFoundOnStackAndPreviewData(t *testing.T) {
	stackData := typStack{}
	previewData := typPreview{}

	result := getCloudProvider(&stackData, &previewData)

	if result != "" {
		t.Errorf("Output %v is not equal to expected value", result)
	}

}

func TestReturnResourceType(t *testing.T) {
	resourceTypes := []typResourceTypesReference{
		{
			CloudResourceType:  "cloudresourcetype",
			PulumiResourceType: "pulumiresourcetype",
		},
	}

	result := getCloudProviderResourceType("pulumiresourcetype", &resourceTypes)

	if result != resourceTypes[0].CloudResourceType {
		t.Errorf("Output %v is not equal to excpected value which is %v", result, resourceTypes[0].CloudResourceType)
	}
}

func TestReturnProperTypSummaryDetails(t *testing.T) {
	resource := typCloudResource{
		Name:     "resourcename",
		Type:     "VM",
		SKU:      "vm_size",
		Location: "asia",
	}

	result := defineTypSummaryDetails(&resource)

	if result.Location != resource.Location || result.Count != 1 {
		t.Errorf("Output %v is not equal to what is expected.", result)
	}
}

func TestReturnProperTypSize(t *testing.T) {
	resource := typCloudResource{
		Name:     "resourcename",
		Type:     "VM",
		SKU:      "vm_size",
		Location: "asia",
	}

	result := defineTypSize(&resource)

	if result.Size != resource.SKU || result.Details[0].Count != 1 || result.Details[0].Location != resource.Location {
		t.Errorf("Output %v is not equal to what is expected.", result)
	}
}

func TestReturnProperTypSummary(t *testing.T) {
	resource := typCloudResource{
		Name:     "resourcename",
		Type:     "VM",
		SKU:      "vm_size",
		Location: "asia",
	}

	result := defineTypSummary(&resource)

	if result.Resource != resource.Type && result.Count != 1 {
		t.Errorf("Output %v is not equal to what is expected.", result)
	}
}

func TestExistingPulumiResourceReturnsTrue(t *testing.T) {
	summary := []TypSummary{
		{
			Resource: "VM",
		},
	}

	exists, index := isExistingPulumiResource(&summary, "VM")

	if !exists || index != 0 {
		t.Errorf("Output %v, %v is not equal to what is expected.", exists, index)
	}
}

func TestExistingPulumiResourceReturnsFalse(t *testing.T) {
	summary := []TypSummary{
		{
			Resource: "VM",
		},
	}

	exists, index := isExistingPulumiResource(&summary, "Storage")

	if exists || index != -1 {
		t.Errorf("Output %v, %v is not equal to what is expected.", exists, index)
	}
}
