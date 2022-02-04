package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSwimImageURL() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Software Image Management (SWIM).

- Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image
files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2
`,

		CreateContext: resourceSwimImageURLCreate,
		ReadContext:   resourceSwimImageURLRead,
		UpdateContext: resourceSwimImageURLUpdate,
		DeleteContext: resourceSwimImageURLDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSoftwareImageManagementSwimImportSoftwareImageViaURL`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"application_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"image_family": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"third_party": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSwimImageURLCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSwimImageURLImportSoftwareImageViaURL(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.SoftwareImageManagementSwim.ImportSoftwareImageViaURL(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ImportSoftwareImageViaURL", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ImportSoftwareImageViaURL", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceSwimImageURLRead(ctx, d, m)
}

func resourceSwimImageURLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

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

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSoftwareImageDetails")
		queryParams1 := dnacentersdkgo.GetSoftwareImageDetailsQueryParams{}

		if okImageUUID {
			queryParams1.ImageUUID = vImageUUID
		}
		if okName {
			queryParams1.Name = vName
		}
		if okFamily {
			queryParams1.Family = vFamily
		}
		if okApplicationType {
			queryParams1.ApplicationType = vApplicationType
		}
		if okImageIntegrityStatus {
			queryParams1.ImageIntegrityStatus = vImageIntegrityStatus
		}
		if okVersion {
			queryParams1.Version = vVersion
		}
		if okImageSeries {
			queryParams1.ImageSeries = vImageSeries
		}
		if okImageName {
			queryParams1.ImageName = vImageName
		}
		if okIsTaggedGolden {
			queryParams1.IsTaggedGolden = *stringToBooleanPtr(vIsTaggedGolden)
		}
		if okIsCCORecommended {
			queryParams1.IsCCORecommended = *stringToBooleanPtr(vIsCCORecommended)
		}
		if okIsCCOLatest {
			queryParams1.IsCCOLatest = *stringToBooleanPtr(vIsCCOLatest)
		}
		if okCreatedTime {
			queryParams1.CreatedTime = *stringToIntPtr(vCreatedTime)
		}
		if okImageSizeGreaterThan {
			queryParams1.ImageSizeGreaterThan = *stringToIntPtr(vImageSizeGreaterThan)
		}
		if okImageSizeLesserThan {
			queryParams1.ImageSizeLesserThan = *stringToIntPtr(vImageSizeLesserThan)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder
		}
		if okLimit {
			queryParams1.Limit = *stringToIntPtr(vLimit)
		}
		if okOffset {
			queryParams1.Offset = *stringToIntPtr(vOffset)
		}

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetSoftwareImageDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetSoftwareImageDetails", err,
				"Failure at GetSoftwareImageDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR DNAC

		vItem1 := flattenSoftwareImageManagementSwimGetSoftwareImageDetailsItems(response1)
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSoftwareImageDetails search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSwimImageURLUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSwimImageURLRead(ctx, d, m)
}

func resourceSwimImageURLDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete SwimImageURL on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestSwimImageURLImportSoftwareImageViaURL(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimImportSoftwareImageViaURL {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimImportSoftwareImageViaURL{}
	if v := expandRequestSwimImageURLImportSoftwareImageViaURLItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSwimImageURLImportSoftwareImageViaURLItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL {
	request := []dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestSwimImageURLImportSoftwareImageViaURLItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSwimImageURLImportSoftwareImageViaURLItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL {
	request := dnacentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_type")))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_family")))) {
		request.ImageFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_url")))) {
		request.SourceURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".third_party")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".third_party")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".third_party")))) {
		request.ThirdParty = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vendor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vendor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vendor")))) {
		request.Vendor = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchSoftwareImageManagementSwimGetSoftwareImageDetails(m interface{}, queryParams dnacentersdkgo.GetSoftwareImageDetailsQueryParams) (*dnacentersdkgo.ResponseItemSoftwareImageManagementSwimGetSoftwareImageDetails, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemSoftwareImageManagementSwimGetSoftwareImageDetails
	var ite *dnacentersdkgo.ResponseSoftwareImageManagementSwimGetSoftwareImageDetails
	ite, _, err = client.SoftwareImageManagementSwim.GetSoftwareImageDetails(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemSoftwareImageManagementSwimGetSoftwareImageDetails
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
