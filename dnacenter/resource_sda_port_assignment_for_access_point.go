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

func resourceSdaPortAssignmentForAccessPoint() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Add Port assignment for access point in SDA Fabric

- Delete Port assignment for access point in SDA Fabric
`,

		CreateContext: resourceSdaPortAssignmentForAccessPointCreate,
		ReadContext:   resourceSdaPortAssignmentForAccessPointRead,
		UpdateContext: resourceSdaPortAssignmentForAccessPointUpdate,
		DeleteContext: resourceSdaPortAssignmentForAccessPointDelete,
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

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate Template Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"data_ip_address_pool_name": &schema.Schema{
							Description: `Data Ip Address Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

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

						"interface_name": &schema.Schema{
							Description: `Interface Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"scalable_group_name": &schema.Schema{
							Description: `Scalable Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"voice_ip_address_pool_name": &schema.Schema{
							Description: `Voice Ip Address Pool Name`,
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

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate TemplateName associated to siteNameHierarchy.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"data_ip_address_pool_name": &schema.Schema{
							Description: `Ip Pool Name, that is assigned to INFRA_VN  
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the edge device 
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_description": &schema.Schema{
							Description: `Details or note of interface assignment
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface_name": &schema.Schema{
							Description: `Interface Name of the edge device 
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy should be a valid fabric site name hierarchy. e.g Global/USA/San Jose
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

func resourceSdaPortAssignmentForAccessPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaPortAssignmentForAccessPointAddPortAssignmentForAccessPointInSdaFabric(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vDeviceManagementIPAddress := resourceItem["device_management_ip_address"]
	vInterfaceName := resourceItem["interface_name"]
	vvDeviceManagementIPAddress := interfaceToString(vDeviceManagementIPAddress)
	vvInterfaceName := interfaceToString(vInterfaceName)

	queryParams1 := dnacentersdkgo.GetPortAssignmentForAccessPointInSdaFabricQueryParams{}

	queryParams1.DeviceManagementIPAddress = vvDeviceManagementIPAddress

	queryParams1.InterfaceName = vvInterfaceName

	getResponse2, _, err := client.Sda.GetPortAssignmentForAccessPointInSdaFabric(&queryParams1)
	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["device_management_ip_address"] = vvDeviceManagementIPAddress
		resourceMap["interface_name"] = vvInterfaceName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaPortAssignmentForAccessPointRead(ctx, d, m)
	}

	response1, restyResp1, err := client.Sda.AddPortAssignmentForAccessPointInSdaFabric(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddPortAssignmentForAccessPointInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddPortAssignmentForAccessPointInSdaFabric", err))
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
				"Failure when executing AddPortAssignmentForAccessPointInSdaFabric", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["device_management_ip_address"] = vvDeviceManagementIPAddress
	resourceMap["interface_name"] = vvInterfaceName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaPortAssignmentForAccessPointRead(ctx, d, m)
}

func resourceSdaPortAssignmentForAccessPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]
	vInterfaceName := resourceMap["interface_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPortAssignmentForAccessPointInSdaFabric")
		queryParams1 := dnacentersdkgo.GetPortAssignmentForAccessPointInSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

		queryParams1.InterfaceName = vInterfaceName

		response1, restyResp1, err := client.Sda.GetPortAssignmentForAccessPointInSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetPortAssignmentForAccessPointInSdaFabric", err,
			// 	"Failure at GetPortAssignmentForAccessPointInSdaFabric, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetPortAssignmentForAccessPointInSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortAssignmentForAccessPointInSdaFabric response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaPortAssignmentForAccessPointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaPortAssignmentForAccessPointRead(ctx, d, m)
}

func resourceSdaPortAssignmentForAccessPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]
	vInterfaceName := resourceMap["interface_name"]

	queryParams1 := dnacentersdkgo.GetPortAssignmentForAccessPointInSdaFabricQueryParams{}
	queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress
	queryParams1.InterfaceName = vInterfaceName
	item, restyResp1, err := client.Sda.GetPortAssignmentForAccessPointInSdaFabric(&queryParams1)
	if err != nil || item == nil {
		/*diags = append(diags, diagErrorWithAlt(
		"Failure when executing GetPortAssignmentForAccessPointInSDAFabric", err,
		"Failure at GetPortAssignmentForAccessPointInSDAFabric, unexpected response", ""))*/
		d.SetId("")
		return diags
	}

	queryParams2 := dnacentersdkgo.DeletePortAssignmentForAccessPointInSdaFabricQueryParams{}
	queryParams2.DeviceManagementIPAddress = vDeviceManagementIPAddress
	queryParams2.InterfaceName = vInterfaceName

	response1, restyResp1, err := client.Sda.DeletePortAssignmentForAccessPointInSdaFabric(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletePortAssignmentForAccessPointInSdaFabric", err, restyResp1.String(),
				"Failure at DeletePortAssignmentForAccessPointInSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletePortAssignmentForAccessPointInSdaFabric", err,
			"Failure at DeletePortAssignmentForAccessPointInSdaFabric, unexpected response", ""))
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
				"Failure when executing DeletePortAssignmentForAccessPointInSdaFabric", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaPortAssignmentForAccessPointAddPortAssignmentForAccessPointInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddPortAssignmentForAccessPointInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddPortAssignmentForAccessPointInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_ip_address_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_ip_address_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_ip_address_pool_name")))) {
		request.DataIPAddressPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authenticate_template_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authenticate_template_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authenticate_template_name")))) {
		request.AuthenticateTemplateName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_description")))) {
		request.InterfaceDescription = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
