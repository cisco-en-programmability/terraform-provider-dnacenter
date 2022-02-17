package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSwimImageFile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Software Image Management (SWIM).

- Fetches a software image from local file system and uploads to DNA Center. Supported software image files extensions
are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2.
Upload the file to the **file** form data field
`,

		CreateContext: resourceSwimImageFileCreate,
		ReadContext:   resourceSwimImageFileRead,
		UpdateContext: resourceSwimImageFileUpdate,
		DeleteContext: resourceSwimImageFileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSwimImageFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	/*
		client := m.(*dnacentersdkgo.Client)

		var diags diag.Diagnostics
		resourceItem := *getResourceItem(d.Get("parameters"))
		request1 := expandRequestSwimImageFileImportLocalSoftwareImage(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

		resp1, restyResp1, err := client.SoftwareImageManagementSwim.ImportLocalSoftwareImage(request1)
		if err != nil || resp1 == nil {
			if restyResp1 != nil {
				diags = append(diags, diagErrorWithResponse(
					"Failure when executing ImportLocalSoftwareImage", err, restyResp1.String()))
				return diags
			}
			diags = append(diags, diagError(
				"Failure when executing ImportLocalSoftwareImage", err))
			return diags
		}
	*/
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceSwimImageFileRead(ctx, d, m)
}

func resourceSwimImageFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	/*
		resourceID := d.Id()
		resourceMap := separateResourceID(resourceID)
		vImageUUID := resourceMap["image_uuid"]
		vName := resourceMap["name"]
		vFamily := resourceMap["family"]
		vApplicationType := resourceMap["application_type"]
		vImageIntegrityStatus := resourceMap["image_integrity_status"]
		vVersion := resourceMap["version"]
		vImageSeries := resourceMap["image_series"]
		vImageName := resourceMap["image_name"]
		vIsTaggedGolden := resourceMap["is_tagged_golden"]
		vIsCCORecommended := resourceMap["is_cco_recommended"]
		vIsCCOLatest := resourceMap["is_cco_latest"]
		vCreatedTime := resourceMap["created_time"]
		vImageSizeGreaterThan := resourceMap["image_size_greater_than"]
		vImageSizeLesserThan := resourceMap["image_size_lesser_than"]
		vSortBy := resourceMap["sort_by"]
		vSortOrder := resourceMap["sort_order"]
		vLimit := resourceMap["limit"]
		vOffset := resourceMap["offset"]
	*/
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSoftwareImageDetails")
		queryParams1 := dnacentersdkgo.GetSoftwareImageDetailsQueryParams{}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetSoftwareImageDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC
		/*
			vItem1 := flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItems(response1)
			if err := d.Set("parameters", vItem1); err != nil {
				diags = append(diags, diagError(
					"Failure when setting GetSoftwareImageDetails search response",
					err))
				return diags
			}
		*/
	}
	return diags
}

func resourceSwimImageFileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSwimImageFileRead(ctx, d, m)
}

func resourceSwimImageFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SwimImageFile on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}

func searchSoftwareImageManagementSwimGetSoftwareImageDetailsFile(m interface{}, queryParams dnacentersdkgo.GetSoftwareImageDetailsQueryParams) (*dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse
	var ite *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetails
	ite, _, err = client.SoftwareImageManagementSwim.GetSoftwareImageDetails(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite.Response
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
