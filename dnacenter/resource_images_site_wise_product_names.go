package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImagesSiteWiseProductNames() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Software Image Management (SWIM).

- Assign network device product name and sites for the given image identifier. Refer */dna/intent/api/v1/images* API for
obtaining imageId

- This resource unassigns the network device product name from all the sites for the given software image.
        Refer to */dna/intent/api/v1/images* and */dna/intent/api/v1/images/{imageId}/siteWiseProductNames* GET APIs for
obtaining  *imageId* and *productNameOrdinal* respectively.

- Update the list of sites for the network device product name assigned to the software image. Refer to
*/dna/intent/api/v1/images* and */dna/intent/api/v1/images/{imageId}/siteWiseProductNames* GET APIs for obtaining
*imageId* and *productNameOrdinal* respectively.
`,

		CreateContext: resourceImagesSiteWiseProductNamesCreate,
		ReadContext:   resourceImagesSiteWiseProductNamesRead,
		UpdateContext: resourceImagesSiteWiseProductNamesUpdate,
		DeleteContext: resourceImagesSiteWiseProductNamesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Product name ordinal is unique value for each network device product
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_ids": &schema.Schema{
							Description: `Supported PIDs
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"product_name": &schema.Schema{
							Description: `Network device product name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_name_ordinal": &schema.Schema{
							Description: `Product name ordinal is unique value for each network device product
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"recommended": &schema.Schema{
							Description: `If 'CISCO' network device product recommandation came from Cisco.com and 'USER' manually assigned
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_ids": &schema.Schema{
							Description: `Sites where all  this image is assigned
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"image_id": &schema.Schema{
							Description: `imageId path parameter. Software image identifier. Refer */dna/intent/api/v1/images* API for obtaining *imageId*
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"product_name_ordinal": &schema.Schema{
							Description: `Product name ordinal is unique value for each network device product
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"site_ids": &schema.Schema{
							Description: `Sites where this image needs to be assigned. Ref https://developer.cisco.com/docs/dna-center/#!sites
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceImagesSiteWiseProductNamesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestImagesSiteWiseProductNamesAssignNetworkDeviceProductNameToTheGivenSoftwareImage(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vProductNameOrdinal := resourceItem["product_name_ordinal"]

	vvProductNameOrdinal := interfaceToString(vProductNameOrdinal)

	vImageID := resourceItem["image_id"]

	vvImageID := interfaceToString(vImageID)
	queryParamImport := dnacentersdkgo.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams{}

	item2, err := searchSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(m, queryParamImport, vvImageID, vvProductNameOrdinal)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["product_name_ordinal"] = fmt.Sprintf("%f", *item2.ProductNameOrdinal)
		resourceMap["image_id"] = vvImageID
		d.SetId(joinResourceID(resourceMap))
		return resourceImagesSiteWiseProductNamesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.SoftwareImageManagementSwim.AssignNetworkDeviceProductNameToTheGivenSoftwareImage(vvImageID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AssignNetworkDeviceProductNameToTheGivenSoftwareImage", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AssignNetworkDeviceProductNameToTheGivenSoftwareImage", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignNetworkDeviceProductNameToTheGivenSoftwareImage", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing AssignNetworkDeviceProductNameToTheGivenSoftwareImage", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams{}

	item3, err := searchSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(m, queryParamValidate, vvImageID, vvProductNameOrdinal)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AssignNetworkDeviceProductNameToTheGivenSoftwareImage", err,
			"Failure at AssignNetworkDeviceProductNameToTheGivenSoftwareImage, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["product_name_ordinal"] = fmt.Sprintf("%f", *item3.ProductNameOrdinal)
	resourceMap["image_id"] = vvImageID
	d.SetId(joinResourceID(resourceMap))
	return resourceImagesSiteWiseProductNamesRead(ctx, d, m)
}

func resourceImagesSiteWiseProductNamesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vImageID := resourceMap["image_id"]
	vvProductNameOrdinal := resourceMap["product_name_ordinal"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage")
		vvImageID := vImageID
		queryParams1 := dnacentersdkgo.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams{}
		item1, err := searchSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(m, queryParams1, vvImageID, vvProductNameOrdinal)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		items := []dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageResponse{
			*item1,
		}

		// Review flatten function used
		vItem1 := flattenSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceImagesSiteWiseProductNamesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvImageID := resourceMap["image_id"]
	vvProductNameOrdinal, _ := strconv.ParseFloat(resourceMap["product_name_ordinal"], 64)
	if d.HasChange("parameters") {
		request1 := expandRequestImagesSiteWiseProductNamesUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SoftwareImageManagementSwim.UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage(vvImageID, vvProductNameOrdinal, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage", err, restyResp1.String(),
					"Failure at UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage", err,
				"Failure at UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage", err1))
				return diags
			}
		}

	}

	return resourceImagesSiteWiseProductNamesRead(ctx, d, m)
}

func resourceImagesSiteWiseProductNamesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvProductNameOrdinal, _ := strconv.ParseFloat(resourceMap["product_name_ordinal"], 64)
	vvImageID := resourceMap["image_id"]
	response1, restyResp1, err := client.SoftwareImageManagementSwim.UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage(vvImageID, vvProductNameOrdinal)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage", err, restyResp1.String(),
				"Failure at UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage", err,
			"Failure at UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestImagesSiteWiseProductNamesAssignNetworkDeviceProductNameToTheGivenSoftwareImage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_name_ordinal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_name_ordinal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_name_ordinal")))) {
		request.ProductNameOrdinal = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestImagesSiteWiseProductNamesUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage {
	request := dnacentersdkgo.RequestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(m interface{}, queryParams dnacentersdkgo.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams, vImageID string, vProductNameOrdinal string) (*dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageResponse
	// var ite *dnacentersdkgo.ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage
	if vProductNameOrdinal != "" {

		vvProductNameOrdinal, _ := strconv.ParseFloat(vProductNameOrdinal, 64)
		queryParams.Offset = 1
		nResponse, _, err := client.SoftwareImageManagementSwim.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(vImageID, nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vvProductNameOrdinal == *item.ProductNameOrdinal {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.SoftwareImageManagementSwim.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(vImageID, &queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	}
	return foundItem, err
}
