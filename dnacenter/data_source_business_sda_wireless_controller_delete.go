package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceBusinessSdaWirelessControllerDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Fabric Wireless.

- Remove WLC from Fabric Domain
`,

		ReadContext: dataSourceBusinessSdaWirelessControllerDeleteRead,
		Schema: map[string]*schema.Schema{
			"device_ipaddress": &schema.Schema{
				Description: `deviceIPAddress query parameter. Device Management IP Address
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBusinessSdaWirelessControllerDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceIPAddress := d.Get("device_ipaddress")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: RemoveWLCFromFabricDomain")
		queryParams1 := dnacentersdkgo.RemoveWLCFromFabricDomainQueryParams{}

		queryParams1.DeviceIPAddress = vDeviceIPAddress.(string)

		response1, restyResp1, err := client.FabricWireless.RemoveWLCFromFabricDomain(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing RemoveWLCFromFabricDomain", err,
				"Failure at RemoveWLCFromFabricDomain, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFabricWirelessRemoveWLCFromFabricDomainItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RemoveWLCFromFabricDomain response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFabricWirelessRemoveWLCFromFabricDomainItem(item *dnacentersdkgo.ResponseFabricWirelessRemoveWLCFromFabricDomain) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
