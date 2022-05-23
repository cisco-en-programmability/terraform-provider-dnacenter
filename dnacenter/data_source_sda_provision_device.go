package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaProvisionDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get Provisioned Wired Device
`,

		ReadContext: dataSourceSdaProvisionDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_management_ip_address": &schema.Schema{
				Description: `deviceManagementIpAddress query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_management_ip_address": &schema.Schema{
							Description: `Device Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"task_id": &schema.Schema{
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"task_status_url": &schema.Schema{
							Description: `Task Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaProvisionDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceManagementIPAddress := d.Get("device_management_ip_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetProvisionedWiredDevice")
		queryParams1 := dnacentersdkgo.GetProvisionedWiredDeviceQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		response1, restyResp1, err := client.Sda.GetProvisionedWiredDevice(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetProvisionedWiredDevice", err,
				"Failure at GetProvisionedWiredDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetProvisionedWiredDeviceItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProvisionedWiredDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetProvisionedWiredDeviceItem(item *dnacentersdkgo.ResponseSdaGetProvisionedWiredDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_management_ip_address"] = item.DeviceManagementIPAddress
	respItem["site_name_hierarchy"] = item.SiteNameHierarchy
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["task_id"] = item.TaskID
	respItem["task_status_url"] = item.TaskStatusURL
	respItem["execution_status_url"] = item.ExecutionStatusURL
	return []map[string]interface{}{
		respItem,
	}
}
