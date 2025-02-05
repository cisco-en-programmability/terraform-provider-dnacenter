package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsNetworkDevicesNetworkDeviceIDNoticesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of field notices affecting the network device
`,

		ReadContext: dataSourceFieldNoticesResultsNetworkDevicesNetworkDeviceIDNoticesCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id query parameter. Id of the field notice
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Return field notices with this type. Available values : SOFTWARE, HARDWARE
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
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFieldNoticesResultsNetworkDevicesNetworkDeviceIDNoticesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vID, okID := d.GetOk("id")
	vType, okType := d.GetOk("type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfFieldNoticesAffectingTheNetworkDevice")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := dnacentersdkgo.GetCountOfFieldNoticesAffectingTheNetworkDeviceQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}

		response1, restyResp1, err := client.Compliance.GetCountOfFieldNoticesAffectingTheNetworkDevice(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfFieldNoticesAffectingTheNetworkDevice", err,
				"Failure at GetCountOfFieldNoticesAffectingTheNetworkDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfFieldNoticesAffectingTheNetworkDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceItem(item *dnacentersdkgo.ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
