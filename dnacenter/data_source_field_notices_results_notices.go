package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsNotices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get field notices
`,

		ReadContext: dataSourceFieldNoticesResultsNoticesRead,
		Schema: map[string]*schema.Schema{
			"device_count": &schema.Schema{
				Description: `deviceCount query parameter. Return field notices with deviceCount greater than this deviceCount
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Id of the field notice
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. Default value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Return field notices with this type. Available values : SOFTWARE, HARDWARE.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_count": &schema.Schema{
							Description: `Number of devices which are vulnerable to this field notice
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"first_publish_date": &schema.Schema{
							Description: `Time at which the field notice was published
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of the field notice
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_date": &schema.Schema{
							Description: `Time at which the field notice was last updated
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the field notice
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"potential_device_count": &schema.Schema{
							Description: `Number of devices which are potentially vulnerable to this field notice
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"publication_url": &schema.Schema{
							Description: `Url for getting field notice details on cisco website
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
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

func dataSourceFieldNoticesResultsNoticesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vDeviceCount, okDeviceCount := d.GetOk("device_count")
	vType, okType := d.GetOk("type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFieldNotices")
		queryParams1 := dnacentersdkgo.GetFieldNoticesQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okDeviceCount {
			queryParams1.DeviceCount = vDeviceCount.(float64)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		response1, restyResp1, err := client.Compliance.GetFieldNotices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFieldNotices", err,
				"Failure at GetFieldNotices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceGetFieldNoticesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFieldNotices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetFieldNoticesItems(items *[]dnacentersdkgo.ResponseComplianceGetFieldNoticesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["publication_url"] = item.PublicationURL
		respItem["device_count"] = item.DeviceCount
		respItem["potential_device_count"] = item.PotentialDeviceCount
		respItem["type"] = item.Type
		respItem["first_publish_date"] = item.FirstPublishDate
		respItem["last_updated_date"] = item.LastUpdatedDate
		respItems = append(respItems, respItem)
	}
	return respItems
}
