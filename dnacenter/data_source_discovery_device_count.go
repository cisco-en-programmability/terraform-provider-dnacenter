package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryDeviceCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns the count of network devices discovered in the given discovery. Discovery ID can be obtained using the "Get
Discoveries by range" API.
`,

		ReadContext: dataSourceDiscoveryDeviceCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Discovery ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"task_id": &schema.Schema{
				Description: `taskId query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vTaskID, okTaskID := d.GetOk("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDevicesDiscoveredByID")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetDevicesDiscoveredByIDQueryParams{}

		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}

		response1, restyResp1, err := client.Discovery.GetDevicesDiscoveredByID(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDevicesDiscoveredByID", err,
				"Failure at GetDevicesDiscoveredByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryGetDevicesDiscoveredByIDItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDevicesDiscoveredByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetDevicesDiscoveredByIDItem(item *dnacentersdkgo.ResponseDiscoveryGetDevicesDiscoveredByID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
