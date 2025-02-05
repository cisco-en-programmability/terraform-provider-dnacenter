package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSettingsDeviceDeploymentsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Returns the count of device deployment status(s) based on filter criteria. For detailed information about the usage of
the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSettingsDeviceDeploymentsCountRead,
		Schema: map[string]*schema.Schema{
			"deploy_activity_id": &schema.Schema{
				Description: `deployActivityId query parameter. activity from the /deploy task response
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_ids": &schema.Schema{
				Description: `networkDeviceIds query parameter. device ids, retrievable from the id attribute in intent/api/v1/network-device
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
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapSettingsDeviceDeploymentsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeployActivityID, okDeployActivityID := d.GetOk("deploy_activity_id")
	vNetworkDeviceIDs, okNetworkDeviceIDs := d.GetOk("network_device_ids")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceDeploymentStatusCount")
		queryParams1 := dnacentersdkgo.GetDeviceDeploymentStatusCountQueryParams{}

		if okDeployActivityID {
			queryParams1.DeployActivityID = vDeployActivityID.(string)
		}
		if okNetworkDeviceIDs {
			queryParams1.NetworkDeviceIDs = vNetworkDeviceIDs.(string)
		}

		response1, restyResp1, err := client.Sensors.GetDeviceDeploymentStatusCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceDeploymentStatusCount", err,
				"Failure at GetDeviceDeploymentStatusCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsGetDeviceDeploymentStatusCountItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceDeploymentStatusCount response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsGetDeviceDeploymentStatusCountItem(item *dnacentersdkgo.ResponseSensorsGetDeviceDeploymentStatusCountResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
