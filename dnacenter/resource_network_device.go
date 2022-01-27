package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and delete operations on Devices.

- Deletes the network device for the given Id
`,

		CreateContext: resourceNetworkDeviceCreate,
		ReadContext:   resourceNetworkDeviceRead,
		UpdateContext: resourceNetworkDeviceUpdate,
		DeleteContext: resourceNetworkDeviceDelete,
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

						"ap_manager_interface_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"associated_wlc_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"boot_date_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_interval": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_status_detail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_failure_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"series": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_contact": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_udp_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"up_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"waas_device_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
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

						"id": &schema.Schema{
							Description: `id path parameter. Device ID
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceNetworkDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceByID")
		vvID := vID

		response1, restyResp1, err := client.Devices.GetDeviceByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing GetDeviceByID", err,
			// 	"Failure at GetDeviceByID, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDeviceByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNetworkDeviceRead(ctx, d, m)
}

func resourceNetworkDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	selectedMethod := 1
	var vvID string
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.Devices.GetDeviceByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	queryParams1 := dnacentersdkgo.DeleteDeviceByIDQueryParams{}
	queryParams1.CleanConfig = true
	response1, restyResp1, err := client.Devices.DeleteDeviceByID(vvID, &queryParams1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceByID", err, restyResp1.String(),
				"Failure at DeleteDeviceByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceByID", err,
			"Failure at DeleteDeviceByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
