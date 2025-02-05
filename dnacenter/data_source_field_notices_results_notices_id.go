package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsNoticesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get field notice by Id
`,

		ReadContext: dataSourceFieldNoticesResultsNoticesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Id of the field notice
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
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

func dataSourceFieldNoticesResultsNoticesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFieldNoticeByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Compliance.GetFieldNoticeByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFieldNoticeByID", err,
				"Failure at GetFieldNoticeByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetFieldNoticeByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFieldNoticeByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetFieldNoticeByIDItem(item *dnacentersdkgo.ResponseComplianceGetFieldNoticeByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["publication_url"] = item.PublicationURL
	respItem["device_count"] = item.DeviceCount
	respItem["potential_device_count"] = item.PotentialDeviceCount
	respItem["type"] = item.Type
	respItem["first_publish_date"] = item.FirstPublishDate
	respItem["last_updated_date"] = item.LastUpdatedDate
	return []map[string]interface{}{
		respItem,
	}
}
