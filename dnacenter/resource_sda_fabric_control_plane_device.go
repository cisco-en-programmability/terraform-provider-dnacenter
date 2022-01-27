package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaFabricControlPlaneDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Delete control plane device in SDA Fabric

- Add control plane device in SDA Fabric
`,

		CreateContext: resourceSdaFabricControlPlaneDeviceCreate,
		ReadContext:   resourceSdaFabricControlPlaneDeviceRead,
		UpdateContext: resourceSdaFabricControlPlaneDeviceUpdate,
		DeleteContext: resourceSdaFabricControlPlaneDeviceDelete,
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

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_management_ip_address": &schema.Schema{
							Description: `Device Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"roles": &schema.Schema{
							Description: `Roles`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the Device which is provisioned successfully
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy of provisioned Device(site should be fabric site)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaFabricControlPlaneDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaFabricControlPlaneDeviceAddControlPlaneDeviceInSdaFabric(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vDeviceManagementIPAddress := resourceItem["device_management_ip_address"]
	vvDeviceManagementIPAddress := interfaceToString(vDeviceManagementIPAddress)

	queryParams1 := dnacentersdkgo.GetControlPlaneDeviceFromSdaFabricQueryParams{}

	queryParams1.DeviceManagementIPAddress = vvDeviceManagementIPAddress

	getResponse2, _, err := client.Sda.GetControlPlaneDeviceFromSdaFabric(&queryParams1)

	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["device_management_ip_address"] = vvDeviceManagementIPAddress
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricControlPlaneDeviceRead(ctx, d, m)
	}

	response1, restyResp1, err := client.Sda.AddControlPlaneDeviceInSdaFabric(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddControlPlaneDeviceInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddControlPlaneDeviceInSdaFabric", err))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddControlPlaneDeviceInSdaFabric", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["device_management_ip_address"] = vvDeviceManagementIPAddress
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricControlPlaneDeviceRead(ctx, d, m)
}

func resourceSdaFabricControlPlaneDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetControlPlaneDeviceFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetControlPlaneDeviceFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

		response1, restyResp1, err := client.Sda.GetControlPlaneDeviceFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetControlPlaneDeviceFromSdaFabric", err,
			// 	"Failure at GetControlPlaneDeviceFromSdaFabric, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetControlPlaneDeviceFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetControlPlaneDeviceFromSdaFabric response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaFabricControlPlaneDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaFabricControlPlaneDeviceRead(ctx, d, m)
}

func resourceSdaFabricControlPlaneDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	queryParams1 := dnacentersdkgo.GetControlPlaneDeviceFromSdaFabricQueryParams{}
	queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress
	item, _, err := client.Sda.GetControlPlaneDeviceFromSdaFabric(&queryParams1)
	if err != nil || item == nil {
		/*diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetControlPlaneDeviceFromSDAFabric", err,
		"Failure at GetControlPlaneDeviceFromSDAFabric, unexpected response", ""))*/
		d.SetId("")
		return diags
	}

	// REVIEW: Add getAllItems and search function to get missing params

	queryParams2 := dnacentersdkgo.DeleteControlPlaneDeviceInSdaFabricQueryParams{}
	queryParams2.DeviceManagementIPAddress = vDeviceManagementIPAddress
	response1, restyResp1, err := client.Sda.DeleteControlPlaneDeviceInSdaFabric(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteControlPlaneDeviceInSdaFabric", err, restyResp1.String(),
				"Failure at DeleteControlPlaneDeviceInSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteControlPlaneDeviceInSdaFabric", err,
			"Failure at DeleteControlPlaneDeviceInSdaFabric, unexpected response", ""))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteControlPlaneDeviceInSdaFabric", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaFabricControlPlaneDeviceAddControlPlaneDeviceInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddControlPlaneDeviceInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddControlPlaneDeviceInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
