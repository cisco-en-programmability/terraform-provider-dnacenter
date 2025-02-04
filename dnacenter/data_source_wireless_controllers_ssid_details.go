package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersSSIDDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieves all details of SSIDs associated with the specific Wireless Controller.
`,

		ReadContext: dataSourceWirelessControllersSSIDDetailsRead,
		Schema: map[string]*schema.Schema{
			"admin_status": &schema.Schema{
				Description: `adminStatus query parameter. Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"managed": &schema.Schema{
				Description: `managed query parameter. If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"ssid_name": &schema.Schema{
				Description: `ssidName query parameter. Employ this query parameter to obtain the details of the SSID corresponding to the provided SSID name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_status": &schema.Schema{
							Description: `Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"l2_security": &schema.Schema{
							Description: `This represents the identifier for the Layer 2 authentication type. The authentication types supported include wpa2_enterprise, wpa2_personal, open, wpa3_enterprise, wpa3_personal, wpa2_wpa3_personal, wpa2_wpa3_enterprise, and open-secured.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"l3_security": &schema.Schema{
							Description: `This represents the identifier for the Layer 3 authentication type. The authentication types supported are 'open' and 'webauth'.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed": &schema.Schema{
							Description: `If the value is 'true,' the SSID is configured through design; if 'false,' it indicates out-of-band configuration on the Wireless LAN Controller.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"radio_policy": &schema.Schema{
							Description: `This represents the identifier for the radio policy. The policies supported include 2.4GHz, 5GHz, and 6GHz.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_name": &schema.Schema{
							Description: `Name of the SSID.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlan_id": &schema.Schema{
							Description: `WLAN ID.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"wlan_profile_name": &schema.Schema{
							Description: `WLAN Profile Name.
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

func dataSourceWirelessControllersSSIDDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vSSIDName, okSSIDName := d.GetOk("ssid_name")
	vAdminStatus, okAdminStatus := d.GetOk("admin_status")
	vManaged, okManaged := d.GetOk("managed")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSSIDDetailsForSpecificWirelessController")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := dnacentersdkgo.GetSSIDDetailsForSpecificWirelessControllerQueryParams{}

		if okSSIDName {
			queryParams1.SSIDName = vSSIDName.(string)
		}
		if okAdminStatus {
			queryParams1.AdminStatus = vAdminStatus.(bool)
		}
		if okManaged {
			queryParams1.Managed = vManaged.(bool)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.Wireless.GetSSIDDetailsForSpecificWirelessController(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSSIDDetailsForSpecificWirelessController", err,
				"Failure at GetSSIDDetailsForSpecificWirelessController, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetSSIDDetailsForSpecificWirelessControllerItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDDetailsForSpecificWirelessController response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetSSIDDetailsForSpecificWirelessControllerItems(items *[]dnacentersdkgo.ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssid_name"] = item.SSIDName
		respItem["wlan_id"] = item.WLANID
		respItem["wlan_profile_name"] = item.WLANProfileName
		respItem["l2_security"] = item.L2Security
		respItem["l3_security"] = item.L3Security
		respItem["radio_policy"] = item.RadioPolicy
		respItem["admin_status"] = boolPtrToString(item.AdminStatus)
		respItem["managed"] = boolPtrToString(item.Managed)
		respItems = append(respItems, respItem)
	}
	return respItems
}
