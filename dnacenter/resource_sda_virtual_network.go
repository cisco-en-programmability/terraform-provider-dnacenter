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

func resourceSdaVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on SDA.

- Delete virtual network (VN) from SDA Fabric

- Add virtual network (VN) in SDA Fabric
`,

		CreateContext: resourceSdaVirtualNetworkCreate,
		ReadContext:   resourceSdaVirtualNetworkRead,
		UpdateContext: resourceSdaVirtualNetworkUpdate,
		DeleteContext: resourceSdaVirtualNetworkDelete,
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

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy should be a valid fabric site name hierarchy.( e.g. Global/USA/San Jose)
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name, that is created in Global level
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

func resourceSdaVirtualNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaVirtualNetworkAddVnInSdaFabric(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vVirtualNetworkName := resourceItem["virtual_network_name"]
	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvVirtualNetworkName := interfaceToString(vVirtualNetworkName)
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)

	queryParams1 := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}

	queryParams1.VirtualNetworkName = vvVirtualNetworkName

	queryParams1.SiteNameHierarchy = vvSiteNameHierarchy

	getResponse2, restyResp1, err := client.Sda.GetVnFromSdaFabric(&queryParams1)
	if err == nil && getResponse2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["virtual_network_name"] = vvVirtualNetworkName
		resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaVirtualNetworkRead(ctx, d, m)
	}
	response1, restyResp1, err := client.Sda.AddVnInSdaFabric(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddVnInSdaFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddVnInSdaFabric", err))
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
				"Failure when executing AddVnInSdaFabric", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["virtual_network_name"] = vvVirtualNetworkName
	resourceMap["site_name_hierarchy"] = vvSiteNameHierarchy
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaVirtualNetworkRead(ctx, d, m)
}

func resourceSdaVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVirtualNetworkName := resourceMap["virtual_network_name"]
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetVnFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		response1, restyResp1, err := client.Sda.GetVnFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetVnFromSdaFabric", err,
			// 	"Failure at GetVnFromSdaFabric, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetVnFromSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVnFromSdaFabric response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSdaVirtualNetworkUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSdaVirtualNetworkRead(ctx, d, m)
}

func resourceSdaVirtualNetworkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vVirtualNetworkName := resourceMap["virtual_network_name"]
	vSiteNameHierarchy := resourceMap["site_name_hierarchy"]

	queryParams1 := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}
	queryParams1.VirtualNetworkName = vVirtualNetworkName
	queryParams1.SiteNameHierarchy = vSiteNameHierarchy
	item, restyResp1, err := client.Sda.GetVnFromSdaFabric(&queryParams1)
	if err != nil || item == nil {
		//diags = append(diags, diagErrorWithAlt(
		//	"Failure when executing GetVNFromSDAFabric", err,
		//	"Failure at GetVNFromSDAFabric, unexpected response", ""))
		d.SetId("")
		return diags
	}

	queryParams2 := dnacentersdkgo.DeleteVnFromSdaFabricQueryParams{}
	queryParams2.VirtualNetworkName = vVirtualNetworkName
	queryParams2.SiteNameHierarchy = vSiteNameHierarchy

	response1, restyResp1, err := client.Sda.DeleteVnFromSdaFabric(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteVnFromSdaFabric", err, restyResp1.String(),
				"Failure at DeleteVnFromSdaFabric, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteVnFromSdaFabric", err,
			"Failure at DeleteVnFromSdaFabric, unexpected response", ""))
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
				"Failure when executing DeleteVnFromSdaFabric", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaVirtualNetworkAddVnInSdaFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddVnInSdaFabric {
	request := dnacentersdkgo.RequestSdaAddVnInSdaFabric{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
