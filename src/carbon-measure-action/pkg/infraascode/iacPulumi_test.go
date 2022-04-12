package infraascode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReturnResourceListFromStackData(t *testing.T) {
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

	assert.Equal(t, result[0].Location, stackData.Deployment.Resources[0].Inputs.Location)
	assert.Equal(t, result[0].Name, stackData.Deployment.Resources[0].Inputs.Name)
	assert.Equal(t, result[0].SKU, stackData.Deployment.Resources[0].Inputs.SKU)
	assert.Equal(t, result[0].Type, stackData.Deployment.Resources[0].Type)
}

func Test_AddResourceToResourceListFromPreviewData(t *testing.T) {
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

	assert.Equal(t, len(resourceList), 2)
}

func Test_UpdateResourceWithDataFromPreview(t *testing.T) {
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

	assert.Equal(t, len(resourceList), 1)
	assert.Equal(t, resourceList[0].SKU, previewData.Steps[0].State.Inputs.SKU)
}

func Test_GetCloudProviderFromStackData(t *testing.T) {
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

	assert.Equal(t, result, "azure")
}

func Test_GetCloudProviderFromPreviewData(t *testing.T) {
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

	assert.Equal(t, result, "azure")
}

func Test_CloudProviderNotFoundOnStackAndPreviewData(t *testing.T) {
	stackData := typStack{}
	previewData := typPreview{}

	result := getCloudProvider(&stackData, &previewData)

	assert.Equal(t, result, "")
}

func Test_ReturnResourceType(t *testing.T) {
	resourceTypes := []typResourceTypesReference{
		{
			CloudResourceType:  "cloudresourcetype",
			PulumiResourceType: "pulumiresourcetype",
		},
	}

	result := getCloudProviderResourceType("pulumiresourcetype", &resourceTypes)

	assert.Equal(t, result, resourceTypes[0].CloudResourceType)
}

func Test_ReturnProperTypSummaryDetails(t *testing.T) {
	resource := typCloudResource{
		Name:     "resourcename",
		Type:     "VM",
		SKU:      "vm_size",
		Location: "asia",
	}

	result := defineTypSummaryDetails(&resource)

	assert.Equal(t, result.Location, resource.Location)
	assert.Equal(t, result.Count, 1)
}

func Test_ReturnProperTypSize(t *testing.T) {
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

func Test_ReturnProperTypSummary(t *testing.T) {
	resource := typCloudResource{
		Name:     "resourcename",
		Type:     "VM",
		SKU:      "vm_size",
		Location: "asia",
	}

	result := defineTypSummary(&resource)

	assert.Equal(t, result.Resource, resource.Type)
	assert.Equal(t, result.Count, 1)
}

func Test_ExistingPulumiResourceReturnsTrue(t *testing.T) {
	summary := []TypSummary{
		{
			Resource: "VM",
		},
	}

	exists, index := isExistingPulumiResource(&summary, "VM")

	assert.True(t, exists)
	assert.Equal(t, index, 0)
}

func Test_ExistingPulumiResourceReturnsFalse(t *testing.T) {
	summary := []TypSummary{
		{
			Resource: "VM",
		},
	}

	exists, index := isExistingPulumiResource(&summary, "Storage")

	assert.False(t, exists)
	assert.Equal(t, index, -1)
}
