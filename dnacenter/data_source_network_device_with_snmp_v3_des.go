package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceWithSNMPV3Des() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns devices added to Cisco DNA center with snmp v3 DES, where siteId is mandatory & accepts offset, limit, sortby,
order which are optional.
`,

		ReadContext: dataSourceNetworkDeviceWithSNMPV3DesRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. Default value is 500
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Row Number.  Default value is 1
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort By
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"family": &schema.Schema{
							Description: `Family`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"hostname": &schema.Schema{
							Description: `Hostname`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_updated": &schema.Schema{
							Description: `Last Updated`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management Ip Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"reachability_status": &schema.Schema{
							Description: `Reachability Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"up_time": &schema.Schema{
							Description: `Up Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceWithSNMPV3DesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES")
		vvSiteID := vSiteID.(string)
		queryParams1 := dnacentersdkgo.ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Devices.ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES(vvSiteID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES", err,
				"Failure at ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESItems(items *[]dnacentersdkgo.ResponseDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["management_ip_address"] = item.ManagementIPAddress
		respItem["hostname"] = item.Hostname
		respItem["type"] = item.Type
		respItem["family"] = item.Family
		respItem["last_updated"] = item.LastUpdated
		respItem["up_time"] = item.UpTime
		respItem["reachability_status"] = item.ReachabilityStatus
		respItems = append(respItems, respItem)
	}
	return respItems
}
