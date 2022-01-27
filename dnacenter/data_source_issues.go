package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIssues() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Intent API to get a list of global issues, issues for a specific device, or issue for a specific client device's MAC
address.
`,

		ReadContext: dataSourceIssuesRead,
		Schema: map[string]*schema.Schema{
			"ai_driven": &schema.Schema{
				Description: `aiDriven query parameter. The issue's AI driven value (Yes or No)(Use only when macAddress and deviceId are not provided)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_id": &schema.Schema{
				Description: `deviceId query parameter. Assurance UUID value of the device in the issue content
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. Ending epoch time in milliseconds of query time window
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"issue_status": &schema.Schema{
				Description: `issueStatus query parameter. The issue's status value (One of ACTIVE, IGNORED, RESOLVED)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter. Client's device MAC address of the issue (format xx:xx:xx:xx:xx:xx)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Description: `priority query parameter. The issue's priority value (One of P1, P2, P3, or P4)(Use only when macAddress and deviceId are not provided)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Assurance UUID value of the site in the issue content
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Starting epoch time in milliseconds of query time window
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ai_driven": &schema.Schema{
							Description: `Ai Driven`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"client_mac": &schema.Schema{
							Description: `Client Mac`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_id": &schema.Schema{
							Description: `Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_role": &schema.Schema{
							Description: `Device Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"issue_id": &schema.Schema{
							Description: `Issue Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"issue_occurence_count": &schema.Schema{
							Description: `Issue Occurence Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"last_occurence_time": &schema.Schema{
							Description: `Last Occurence Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIssuesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSiteID, okSiteID := d.GetOk("site_id")
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vPriority, okPriority := d.GetOk("priority")
	vAiDriven, okAiDriven := d.GetOk("ai_driven")
	vIssueStatus, okIssueStatus := d.GetOk("issue_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: Issues")
		queryParams1 := dnacentersdkgo.IssuesQueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okPriority {
			queryParams1.Priority = vPriority.(string)
		}
		if okAiDriven {
			queryParams1.AiDriven = vAiDriven.(string)
		}
		if okIssueStatus {
			queryParams1.IssueStatus = vIssueStatus.(string)
		}

		response1, restyResp1, err := client.Issues.Issues(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing Issues", err,
				"Failure at Issues, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenIssuesIssuesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting Issues response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesIssuesItems(items *[]dnacentersdkgo.ResponseIssuesIssuesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["issue_id"] = item.IssueID
		respItem["name"] = item.Name
		respItem["site_id"] = item.SiteID
		respItem["device_id"] = item.DeviceID
		respItem["device_role"] = item.DeviceRole
		respItem["ai_driven"] = boolPtrToString(item.AiDriven)
		respItem["client_mac"] = item.ClientMac
		respItem["issue_occurence_count"] = item.IssueOccurenceCount
		respItem["status"] = item.Status
		respItem["priority"] = item.Priority
		respItem["category"] = item.Category
		respItem["last_occurence_time"] = item.LastOccurenceTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
