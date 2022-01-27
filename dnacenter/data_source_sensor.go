package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSensor() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Intent API to get a list of SENSOR devices
`,

		ReadContext: dataSourceSensorRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"backhaul_type": &schema.Schema{
							Description: `Backhaul Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ethernet_mac_address": &schema.Schema{
							Description: `Ethernet Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ip_address": &schema.Schema{
							Description: `Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_led_enabled": &schema.Schema{
							Description: `Is L E D Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_seen": &schema.Schema{
							Description: `Last Seen`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"location": &schema.Schema{
							Description: `Location`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"radio_mac_address": &schema.Schema{
							Description: `Radio Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"ssh_config": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_password": &schema.Schema{
										Description: `Enable Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ssh_password": &schema.Schema{
										Description: `Ssh Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ssh_state": &schema.Schema{
										Description: `Ssh State`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"ssh_user_name": &schema.Schema{
										Description: `Ssh User Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSensorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Sensors")
		queryParams1 := dnacentersdkgo.SensorsQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		response1, restyResp1, err := client.Sensors.Sensors(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Sensors", err,
				"Failure at Sensors, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSensorsSensorsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Sensors response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsSensorsItems(items *[]dnacentersdkgo.ResponseSensorsSensorsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["status"] = item.Status
		respItem["radio_mac_address"] = item.RadioMacAddress
		respItem["ethernet_mac_address"] = item.EthernetMacAddress
		respItem["location"] = item.Location
		respItem["backhaul_type"] = item.BackhaulType
		respItem["serial_number"] = item.SerialNumber
		respItem["ip_address"] = item.IPAddress
		respItem["version"] = item.Version
		respItem["last_seen"] = item.LastSeen
		respItem["type"] = item.Type
		respItem["ssh_config"] = flattenSensorsSensorsItemsSSHConfig(item.SSHConfig)
		respItem["is_led_enabled"] = boolPtrToString(item.IsLEDEnabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsSensorsItemsSSHConfig(item *dnacentersdkgo.ResponseSensorsSensorsResponseSSHConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ssh_state"] = item.SSHState
	respItem["ssh_user_name"] = item.SSHUserName
	respItem["ssh_password"] = item.SSHPassword
	respItem["enable_password"] = item.EnablePassword

	return []map[string]interface{}{
		respItem,
	}

}
