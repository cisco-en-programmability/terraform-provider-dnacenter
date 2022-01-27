package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigurationTemplateDeployStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- API to retrieve the status of template deployment.
`,

		ReadContext: dataSourceConfigurationTemplateDeployStatusRead,
		Schema: map[string]*schema.Schema{
			"deployment_id": &schema.Schema{
				Description: `deploymentId path parameter. UUID of deployment to retrieve template deployment status
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"deployment_id": &schema.Schema{
							Description: `UUID of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"deployment_name": &schema.Schema{
							Description: `Name of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"detailed_status_message": &schema.Schema{
										Description: `Device detailed status message
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_id": &schema.Schema{
										Description: `UUID of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"duration": &schema.Schema{
										Description: `Total duration of deployment
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_time": &schema.Schema{
										Description: `EndTime of deployment
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"identifier": &schema.Schema{
										Description: `Identifier of device based on the target type
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_address": &schema.Schema{
										Description: `Device IPAddress
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_time": &schema.Schema{
										Description: `StartTime of deployment
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Current status of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"target_type": &schema.Schema{
										Description: `Target type of device
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"duration": &schema.Schema{
							Description: `Total deployment duration
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_time": &schema.Schema{
							Description: `Deployment end time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Description: `Deployment start time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Current status of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status_message": &schema.Schema{
							Description: `Status message of deployment
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"template_name": &schema.Schema{
							Description: `Name of template deployed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"template_version": &schema.Schema{
							Description: `Version of template deployed
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

func dataSourceConfigurationTemplateDeployStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeploymentID := d.Get("deployment_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: StatusOfTemplateDeployment")
		vvDeploymentID := vDeploymentID.(string)

		response1, restyResp1, err := client.ConfigurationTemplates.StatusOfTemplateDeployment(vvDeploymentID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing StatusOfTemplateDeployment", err,
				"Failure at StatusOfTemplateDeployment, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesStatusOfTemplateDeploymentItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting StatusOfTemplateDeployment response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesStatusOfTemplateDeploymentItem(item *dnacentersdkgo.ResponseConfigurationTemplatesStatusOfTemplateDeployment) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["deployment_id"] = item.DeploymentID
	respItem["deployment_name"] = item.DeploymentName
	respItem["devices"] = flattenConfigurationTemplatesStatusOfTemplateDeploymentItemDevices(item.Devices)
	respItem["duration"] = item.Duration
	respItem["end_time"] = item.EndTime
	respItem["project_name"] = item.ProjectName
	respItem["start_time"] = item.StartTime
	respItem["status"] = item.Status
	respItem["status_message"] = item.StatusMessage
	respItem["template_name"] = item.TemplateName
	respItem["template_version"] = item.TemplateVersion
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesStatusOfTemplateDeploymentItemDevices(items *[]dnacentersdkgo.ResponseConfigurationTemplatesStatusOfTemplateDeploymentDevices) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["detailed_status_message"] = item.DetailedStatusMessage
		respItem["device_id"] = item.DeviceID
		respItem["duration"] = item.Duration
		respItem["end_time"] = item.EndTime
		respItem["identifier"] = item.IDentifier
		respItem["ip_address"] = item.IPAddress
		respItem["name"] = item.Name
		respItem["start_time"] = item.StartTime
		respItem["status"] = item.Status
		respItem["target_type"] = item.TargetType
		respItems = append(respItems, respItem)
	}
	return respItems
}
