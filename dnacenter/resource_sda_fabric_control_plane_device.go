package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

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
							Description: `Control plane device info retrieved successfully in sda fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the Device which is provisioned successfully
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_name": &schema.Schema{
							Description: `Device Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": &schema.Schema{
							Description: `Assigned roles
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_distribution_protocol": &schema.Schema{
							Description: `routeDistributionProtocol
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
							Computed: true,
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

						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the Device which is provisioned successfully
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_distribution_protocol": &schema.Schema{
							Description: `Route Distribution Protocol for Control Plane Device. Allowed values are "LISP_BGP" or "LISP_PUB_SUB". Default value is "LISP_BGP"
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `siteNameHierarchy of the Provisioned Device(site should be part of Fabric Site)
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vDeviceManagementIPAddress := resourceItem["device_management_ip_address"]
	vvDeviceManagementIPAddress := interfaceToString(vDeviceManagementIPAddress)
	queryParamImport := dnacentersdkgo.GetControlPlaneDeviceFromSdaFabricQueryParams{}
	queryParamImport.DeviceManagementIPAddress = vvDeviceManagementIPAddress
	item2, _, err := client.Sda.GetControlPlaneDeviceFromSdaFabric(&queryParamImport)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["device_management_ip_address"] = item2.DeviceManagementIPAddress
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricControlPlaneDeviceRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddControlPlaneDeviceInSdaFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddControlPlaneDeviceInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddControlPlaneDeviceInSdaFabric", err))
		return diags
	}
	executionId := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddControlPlaneDeviceInSdaFabric", err))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetControlPlaneDeviceFromSdaFabricQueryParams{}
	queryParamValidate.DeviceManagementIPAddress = vvDeviceManagementIPAddress
	item3, _, err := client.Sda.GetControlPlaneDeviceFromSdaFabric(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddControlPlaneDeviceInSdaFabric", err,
			"Failure at AddControlPlaneDeviceInSdaFabric, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["device_management_ip_address"] = item3.DeviceManagementIPAddress

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
		log.Printf("[DEBUG] Selected method: GetControlPlaneDeviceFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetControlPlaneDeviceFromSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

		response1, restyResp1, err := client.Sda.GetControlPlaneDeviceFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
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
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SdaFabricControlPlaneDeviceUpdate", err, "Update method is not supported",
		"Failure at SdaFabricControlPlaneDeviceUpdate, unexpected response", ""))

	return diags
}

func resourceSdaFabricControlPlaneDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := dnacentersdkgo.DeleteControlPlaneDeviceInSdaFabricQueryParams{}

	vvDeviceManagementIPAddress := resourceMap["device_management_ip_address"]
	queryParamDelete.DeviceManagementIPAddress = vvDeviceManagementIPAddress

	response1, restyResp1, err := client.Sda.DeleteControlPlaneDeviceInSdaFabric(&queryParamDelete)
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
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
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
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".route_distribution_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".route_distribution_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".route_distribution_protocol")))) {
		request.RouteDistributionProtocol = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
