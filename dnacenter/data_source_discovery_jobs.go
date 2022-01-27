package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryJobs() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns the list of discovery jobs for the given IP
`,

		ReadContext: dataSourceDiscoveryJobsRead,
		Schema: map[string]*schema.Schema{
			"ip_address": &schema.Schema{
				Description: `ipAddress query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeInt,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"cli_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"http_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory_collection_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory_reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"job_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"netconf_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ping_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"task_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryJobsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vIPAddress := d.Get("ip_address")
	vName, okName := d.GetOk("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDiscoveryJobsByIP")
		queryParams1 := dnacentersdkgo.GetDiscoveryJobsByIPQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		queryParams1.IPAddress = vIPAddress.(string)

		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.Discovery.GetDiscoveryJobsByIP(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDiscoveryJobsByIP", err,
				"Failure at GetDiscoveryJobsByIP, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDiscoveryGetDiscoveryJobsByIPItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDiscoveryJobsByIP response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetDiscoveryJobsByIPItems(items *[]dnacentersdkgo.ResponseDiscoveryGetDiscoveryJobsByIPResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDiscoveryGetDiscoveryJobsByIPItemsAttributeInfo(item.AttributeInfo)
		respItem["cli_status"] = item.CliStatus
		respItem["discovery_status"] = item.DiscoveryStatus
		respItem["end_time"] = item.EndTime
		respItem["http_status"] = item.HTTPStatus
		respItem["id"] = item.ID
		respItem["inventory_collection_status"] = item.InventoryCollectionStatus
		respItem["inventory_reachability_status"] = item.InventoryReachabilityStatus
		respItem["ip_address"] = item.IPAddress
		respItem["job_status"] = item.JobStatus
		respItem["name"] = item.Name
		respItem["netconf_status"] = item.NetconfStatus
		respItem["ping_status"] = item.PingStatus
		respItem["snmp_status"] = item.SNMPStatus
		respItem["start_time"] = item.StartTime
		respItem["task_id"] = item.TaskID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDiscoveryGetDiscoveryJobsByIPItemsAttributeInfo(item *dnacentersdkgo.ResponseDiscoveryGetDiscoveryJobsByIPResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
