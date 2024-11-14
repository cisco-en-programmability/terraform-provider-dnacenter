package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoverySummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns the devices discovered in the given discovery based on given filters. Discovery ID can be obtained using the
"Get Discoveries by range" API.
`,

		ReadContext: dataSourceDiscoverySummaryRead,
		Schema: map[string]*schema.Schema{
			"clistatus": &schema.Schema{
				Description: `cliStatus query parameter. CLI status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"http_status": &schema.Schema{
				Description: `httpStatus query parameter. HTTP staus for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"id": &schema.Schema{
				Description: `id path parameter. Discovery ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": &schema.Schema{
				Description: `ipAddress query parameter. IP Address of the device
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"netconf_status": &schema.Schema{
				Description: `netconfStatus query parameter. NETCONF status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ping_status": &schema.Schema{
				Description: `pingStatus query parameter. Ping status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"snmp_status": &schema.Schema{
				Description: `snmpStatus query parameter. SNMP status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort by field. Available values are pingStatus, cliStatus,snmpStatus, httpStatus and netconfStatus
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Order of sorting based on sortBy. Available values are 'asc' and 'des'
`,
				Type:     schema.TypeString,
				Optional: true,
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
							Description: `The number of network devices from the discovery job based on the given filters
`,
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

func dataSourceDiscoverySummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vTaskID, okTaskID := d.GetOk("task_id")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vIPAddress, okIPAddress := d.GetOk("ip_address")
	vPingStatus, okPingStatus := d.GetOk("ping_status")
	vSNMPStatus, okSNMPStatus := d.GetOk("snmp_status")
	vClistatus, okClistatus := d.GetOk("clistatus")
	vNetconfStatus, okNetconfStatus := d.GetOk("netconf_status")
	vHTTPStatus, okHTTPStatus := d.GetOk("http_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkDevicesFromDiscovery")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetNetworkDevicesFromDiscoveryQueryParams{}

		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okIPAddress {
			queryParams1.IPAddress = interfaceToSliceString(vIPAddress)
		}
		if okPingStatus {
			queryParams1.PingStatus = interfaceToSliceString(vPingStatus)
		}
		if okSNMPStatus {
			queryParams1.SNMPStatus = interfaceToSliceString(vSNMPStatus)
		}
		if okClistatus {
			queryParams1.Clistatus = interfaceToSliceString(vClistatus)
		}
		if okNetconfStatus {
			queryParams1.NetconfStatus = interfaceToSliceString(vNetconfStatus)
		}
		if okHTTPStatus {
			queryParams1.HTTPStatus = interfaceToSliceString(vHTTPStatus)
		}

		response1, restyResp1, err := client.Discovery.GetNetworkDevicesFromDiscovery(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetNetworkDevicesFromDiscovery", err,
				"Failure at GetNetworkDevicesFromDiscovery, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryGetNetworkDevicesFromDiscoveryItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkDevicesFromDiscovery response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetNetworkDevicesFromDiscoveryItem(item *dnacentersdkgo.ResponseDiscoveryGetNetworkDevicesFromDiscovery) []map[string]interface{} {
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
