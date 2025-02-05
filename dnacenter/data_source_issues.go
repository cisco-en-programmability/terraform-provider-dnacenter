package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

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
				Description: `aiDriven query parameter. The issue's AI driven value: YES or NO (case insensitive) (Use only when macAddress and deviceId are not provided)
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
				Description: `issueStatus query parameter. The issue's status value: ACTIVE, IGNORED, RESOLVED (case insensitive)
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
				Description: `priority query parameter. The issue's priority value: P1, P2, P3, or P4 (case insensitive) (Use only when macAddress and deviceId are not provided)
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
							Description: `Whether the issue is AI driven ('Yes' or 'No')
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"category": &schema.Schema{
							Description: `Category of the issue
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"client_mac": &schema.Schema{
							Description: `The client MAC address related to this issue
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Description: `The device UUID where the issue occurred
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_role": &schema.Schema{
							Description: `The device role
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"issue_id": &schema.Schema{
							Description: `The issue's unique identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"issue_occurence_count": &schema.Schema{
							Description: `Total number of instances of this issue in the query time window
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"last_occurence_time": &schema.Schema{
							Description: `The UTC timestamp of last occurence of this issue
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `The issue's display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"priority": &schema.Schema{
							Description: `Priority setting of the issue
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_id": &schema.Schema{
							Description: `The site UUID where the issue occurred
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `The status of the issue
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

func dataSourceIssuesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vSiteID, okSiteID := d.GetOk("site_id")
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vPriority, okPriority := d.GetOk("priority")
	vIssueStatus, okIssueStatus := d.GetOk("issue_status")
	vAiDriven, okAiDriven := d.GetOk("ai_driven")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: Issues")
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
		if okIssueStatus {
			queryParams1.IssueStatus = vIssueStatus.(string)
		}
		if okAiDriven {
			queryParams1.AiDriven = vAiDriven.(string)
		}

		response1, restyResp1, err := client.Issues.Issues(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 Issues", err,
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
		respItem["ai_driven"] = item.AiDriven
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
