package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricsFabricIDWirelessMulticast() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Fabric Wireless.

- Retrieves the current Software-Defined Access (SDA) Wireless Multicast setting for a specified fabric site. The
setting indicates whether multicast is enabled (true) or disabled (false). For optimal performance, ensure wired
multicast is also enabled.
`,

		ReadContext: dataSourceFabricsFabricIDWirelessMulticastRead,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId path parameter. The unique identifier of the fabric site for which the multicast setting is being requested. The identifier should be in the format of a UUID. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"multicast_enabled": &schema.Schema{
							Description: `The setting indicates whether multicast is enabled (true) or disabled (false).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFabricsFabricIDWirelessMulticastRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID := d.Get("fabric_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSdaWirelessMulticast")
		vvFabricID := vFabricID.(string)

		response1, restyResp1, err := client.FabricWireless.GetSdaWirelessMulticast(vvFabricID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSdaWirelessMulticast", err,
				"Failure at GetSdaWirelessMulticast, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFabricWirelessGetSdaWirelessMulticastItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSdaWirelessMulticast response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessGetSdaWirelessMulticastItem(item *dnacentersdkgo.ResponseFabricWirelessGetSdaWirelessMulticastResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["multicast_enabled"] = boolPtrToString(item.MulticastEnabled)
	return []map[string]interface{}{
		respItem,
	}
}
