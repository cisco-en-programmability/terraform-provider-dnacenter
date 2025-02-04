package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

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
							Description: `Virtual Network info retrieved successfully from SDA Fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_id": &schema.Schema{
							Description: `Execution Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_name": &schema.Schema{
							Description: `Fabric Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default_vn": &schema.Schema{
							Description: `Default VN
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_infra_vn": &schema.Schema{
							Description: `Infra VN
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Path of sda Fabric Site
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
						"virtual_network_context_id": &schema.Schema{
							Description: `Virtual Network Context Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_network_id": &schema.Schema{
							Description: `Virtual Network Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name
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

						"site_name_hierarchy": &schema.Schema{
							Description: `Path of sda Fabric Site
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name, that is created at Global level
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

func resourceSdaVirtualNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaVirtualNetworkAddVnInFabric(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vVirtualNetworkName := resourceItem["virtual_network_name"]
	vvVirtualNetworkName := interfaceToString(vVirtualNetworkName)
	vSiteNameHierarchy := resourceItem["site_name_hierarchy"]
	vvSiteNameHierarchy := interfaceToString(vSiteNameHierarchy)
	queryParamImport := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}
	queryParamImport.VirtualNetworkName = vvVirtualNetworkName
	queryParamImport.SiteNameHierarchy = vvSiteNameHierarchy
	item2, _, err := client.Sda.GetVnFromSdaFabric(&queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["virtual_network_name"] = item2.VirtualNetworkName
		resourceMap["site_name_hierarchy"] = item2.SiteNameHierarchy
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaVirtualNetworkRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddVnInFabric(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddVnInFabric", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddVnInFabric", err))
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
		for response2.Status == "IN_PROGRESS" {
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
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddVnInFabric", err))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}
	queryParamValidate.VirtualNetworkName = vvVirtualNetworkName
	queryParamValidate.SiteNameHierarchy = vvSiteNameHierarchy
	item3, _, err := client.Sda.GetVnFromSdaFabric(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddVnInFabric", err,
			"Failure at AddVnInFabric, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["virtual_network_name"] = item3.VirtualNetworkName
	resourceMap["site_name_hierarchy"] = item3.SiteNameHierarchy

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
		log.Printf("[DEBUG] Selected method: GetVnFromSdaFabric")
		queryParams1 := dnacentersdkgo.GetVnFromSdaFabricQueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName

		queryParams1.SiteNameHierarchy = vSiteNameHierarchy

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.GetVnFromSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
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

	queryParamDelete := dnacentersdkgo.DeleteVnFromSdaFabricQueryParams{}

	vvVirtualNetworkName := resourceMap["virtual_network_name"]

	vvSiteNameHierarchy := resourceMap["site_name_hierarchy"]
	queryParamDelete.VirtualNetworkName = vvVirtualNetworkName

	queryParamDelete.SiteNameHierarchy = vvSiteNameHierarchy

	response1, restyResp1, err := client.Sda.DeleteVnFromSdaFabric(&queryParamDelete)
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
		for response2.Status == "IN_PROGRESS" {
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
func expandRequestSdaVirtualNetworkAddVnInFabric(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaAddVnInFabric {
	request := dnacentersdkgo.RequestSdaAddVnInFabric{}
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
