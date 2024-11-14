package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

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
							Description: `Backhall type: WIRED, WIRELESS
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ethernet_mac_address": &schema.Schema{
							Description: `Sensor device's ethernet MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Description: `IP Address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_seen": &schema.Schema{
							Description: `Last seen timestamp
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"led": &schema.Schema{
							Description: `Is LED Enabled
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Description: `Site name in hierarchy form
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `The sensor device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"radio_mac_address": &schema.Schema{
							Description: `Sensor device's radio MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial number
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssh": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enable_password": &schema.Schema{
										Description: `Enable password
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssh_password": &schema.Schema{
										Description: `SSH password
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssh_state": &schema.Schema{
										Description: `SSH state
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssh_user_name": &schema.Schema{
										Description: `SSH user name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"status": &schema.Schema{
							Description: `Status of sensor device (REACHABLE, UNREACHABLE, DELETED, RUNNING, IDLE, UCLAIMED)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Sensor version
`,
							Type:     schema.TypeString,
							Computed: true,
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
		log.Printf("[DEBUG] Selected method: Sensors")
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
				"Failure when executing 2 Sensors", err,
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
		respItem["ssh"] = flattenSensorsSensorsItemsSSH(item.SSH)
		respItem["led"] = boolPtrToString(item.Led)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsSensorsItemsSSH(item *dnacentersdkgo.ResponseSensorsSensorsResponseSSH) []map[string]interface{} {
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
