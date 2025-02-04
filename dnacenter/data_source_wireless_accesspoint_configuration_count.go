package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessAccesspointConfigurationCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Get Access Point Configuration Count
`,

		ReadContext: dataSourceWirelessAccesspointConfigurationCountRead,
		Schema: map[string]*schema.Schema{
			"ap_mode": &schema.Schema{
				Description: `apMode query parameter. AP Mode. Allowed values are Local, Bridge, Monitor, FlexConnect, Sniffer, Rogue Detector, SE-Connect, Flex+Bridge, Sensor.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_model": &schema.Schema{
				Description: `apModel query parameter. AP Model
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"mesh_role": &schema.Schema{
				Description: `meshRole query parameter. Mesh Role. Allowed values are RAP or MAP
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"provisioned": &schema.Schema{
				Description: `provisioned query parameter. Indicate whether AP provisioned or not. Allowed values are True or False
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"wlc_ip_address": &schema.Schema{
				Description: `wlcIpAddress query parameter. WLC IP Address
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count of the requested resource
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessAccesspointConfigurationCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vWlcIPAddress, okWlcIPAddress := d.GetOk("wlc_ip_address")
	vApMode, okApMode := d.GetOk("ap_mode")
	vApModel, okApModel := d.GetOk("ap_model")
	vMeshRole, okMeshRole := d.GetOk("mesh_role")
	vProvisioned, okProvisioned := d.GetOk("provisioned")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointConfigurationCount")
		queryParams1 := dnacentersdkgo.GetAccessPointConfigurationCountQueryParams{}

		if okWlcIPAddress {
			queryParams1.WlcIPAddress = vWlcIPAddress.(string)
		}
		if okApMode {
			queryParams1.ApMode = vApMode.(string)
		}
		if okApModel {
			queryParams1.ApModel = vApModel.(string)
		}
		if okMeshRole {
			queryParams1.MeshRole = vMeshRole.(string)
		}
		if okProvisioned {
			queryParams1.Provisioned = vProvisioned.(string)
		}

		response1, restyResp1, err := client.Wireless.GetAccessPointConfigurationCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointConfigurationCount", err,
				"Failure at GetAccessPointConfigurationCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAccessPointConfigurationCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointConfigurationCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAccessPointConfigurationCountItem(item *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
