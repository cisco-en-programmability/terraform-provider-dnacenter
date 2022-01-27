package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNfvProvisionDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Returns provisioning device information for the specified IP address.
`,

		ReadContext: dataSourceNfvProvisionDetailRead,
		Schema: map[string]*schema.Schema{
			"device_ip": &schema.Schema{
				Description: `deviceIp query parameter. Device to which the provisioning detail has to be retrieved
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"begin_step": &schema.Schema{
							Description: `Begin Step`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"duration": &schema.Schema{
							Description: `Duration`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status_message": &schema.Schema{
							Description: `Status Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"task_nodes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_template_user_message_dto": &schema.Schema{
										Description: `Cli Template User Message D T O`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"duration": &schema.Schema{
										Description: `Duration`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"end_time": &schema.Schema{
										Description: `End Time`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"error_payload": &schema.Schema{
										Description: `Error Payload`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"next_task": &schema.Schema{
										Description: `Next Task`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"parent_task": &schema.Schema{
										Description: `Parent Task`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"payload": &schema.Schema{
										Description: `Payload`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"provisioned_names": &schema.Schema{
										Description: `Provisioned Names`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"start_time": &schema.Schema{
										Description: `Start Time`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"status": &schema.Schema{
										Description: `Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"status_message": &schema.Schema{
										Description: `Status Message`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"step_ran": &schema.Schema{
										Description: `Step Ran`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"target": &schema.Schema{
										Description: `Target`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"topology": &schema.Schema{
							Description: `Topology`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNfvProvisionDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceIP := d.Get("device_ip")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceDetailsByIP")
		queryParams1 := dnacentersdkgo.GetDeviceDetailsByIPQueryParams{}

		queryParams1.DeviceIP = vDeviceIP.(string)

		response1, restyResp1, err := client.SiteDesign.GetDeviceDetailsByIP(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceDetailsByIP", err,
				"Failure at GetDeviceDetailsByIP, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetDeviceDetailsByIPItem(response1.ProvisionDetails)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceDetailsByIP response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetDeviceDetailsByIPItem(item *dnacentersdkgo.ResponseSiteDesignGetDeviceDetailsByIPProvisionDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["duration"] = item.Duration
	respItem["status_message"] = item.StatusMessage
	respItem["status"] = item.Status
	respItem["task_nodes"] = flattenSiteDesignGetDeviceDetailsByIPItemTaskNodes(item.TaskNodes)
	respItem["topology"] = item.Topology
	respItem["begin_step"] = item.BeginStep
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSiteDesignGetDeviceDetailsByIPItemTaskNodes(items *[]dnacentersdkgo.ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["duration"] = item.Duration
		respItem["status"] = item.Status
		respItem["next_task"] = item.NextTask
		respItem["name"] = item.Name
		respItem["target"] = item.Target
		respItem["status_message"] = item.StatusMessage
		respItem["payload"] = item.Payload
		respItem["provisioned_names"] = flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesProvisionedNames(item.ProvisionedNames)
		respItem["error_payload"] = flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesErrorPayload(item.ErrorPayload)
		respItem["parent_task"] = flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesParentTask(item.ParentTask)
		respItem["cli_template_user_message_dto"] = flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesCliTemplateUserMessageDTO(item.CliTemplateUserMessageDTO)
		respItem["step_ran"] = item.StepRan
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesProvisionedNames(item *dnacentersdkgo.ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesProvisionedNames) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesErrorPayload(item *dnacentersdkgo.ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesErrorPayload) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesParentTask(item *dnacentersdkgo.ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesParentTask) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSiteDesignGetDeviceDetailsByIPItemTaskNodesCliTemplateUserMessageDTO(item *dnacentersdkgo.ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesCliTemplateUserMessageDTO) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
