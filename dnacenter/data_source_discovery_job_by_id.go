package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryJobByID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns the list of discovery jobs for the given Discovery ID. The results can be optionally filtered based on IP.
Discovery ID can be obtained using the "Get Discoveries by range" API.
`,

		ReadContext: dataSourceDiscoveryJobByIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Discovery ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": &schema.Schema{
				Description: `ipAddress query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeInt,
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

func dataSourceDiscoveryJobByIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vIPAddress, okIPAddress := d.GetOk("ip_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetListOfDiscoveriesByDiscoveryID")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetListOfDiscoveriesByDiscoveryIDQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okIPAddress {
			queryParams1.IPAddress = vIPAddress.(string)
		}

		response1, restyResp1, err := client.Discovery.GetListOfDiscoveriesByDiscoveryID(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetListOfDiscoveriesByDiscoveryID", err,
				"Failure at GetListOfDiscoveriesByDiscoveryID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDiscoveryGetListOfDiscoveriesByDiscoveryIDItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetListOfDiscoveriesByDiscoveryID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetListOfDiscoveriesByDiscoveryIDItems(items *[]dnacentersdkgo.ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDiscoveryGetListOfDiscoveriesByDiscoveryIDItemsAttributeInfo(item.AttributeInfo)
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

func flattenDiscoveryGetListOfDiscoveriesByDiscoveryIDItemsAttributeInfo(item *dnacentersdkgo.ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
