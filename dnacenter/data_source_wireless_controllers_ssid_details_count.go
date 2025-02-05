package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersSSIDDetailsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieves the count of SSIDs associated with the specific Wireless Controller.
`,

		ReadContext: dataSourceWirelessControllersSSIDDetailsCountRead,
		Schema: map[string]*schema.Schema{
			"admin_status": &schema.Schema{
				Description: `adminStatus query parameter. Utilize this query parameter to obtain the number of SSIDs according to their administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
`,
				Type:     schema.TypeBool,
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

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The count of the SSIDs.
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

func dataSourceWirelessControllersSSIDDetailsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vAdminStatus, okAdminStatus := d.GetOk("admin_status")
	vManaged, okManaged := d.GetOk("managed")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSSIDCountForSpecificWirelessController")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := dnacentersdkgo.GetSSIDCountForSpecificWirelessControllerQueryParams{}

		if okAdminStatus {
			queryParams1.AdminStatus = vAdminStatus.(bool)
		}
		if okManaged {
			queryParams1.Managed = vManaged.(bool)
		}

		response1, restyResp1, err := client.Wireless.GetSSIDCountForSpecificWirelessController(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSSIDCountForSpecificWirelessController", err,
				"Failure at GetSSIDCountForSpecificWirelessController, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetSSIDCountForSpecificWirelessControllerItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDCountForSpecificWirelessController response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetSSIDCountForSpecificWirelessControllerItem(item *dnacentersdkgo.ResponseWirelessGetSSIDCountForSpecificWirelessControllerResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
