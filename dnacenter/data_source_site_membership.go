package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteMembership() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Getting the site children details and device details.
`,

		ReadContext: dataSourceSiteMembershipRead,
		Schema: map[string]*schema.Schema{
			"device_family": &schema.Schema{
				Description: `deviceFamily query parameter. Device family name 
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of sites to be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset/starting row
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. Device serial number
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site id to retrieve device associated with the site.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"response": &schema.Schema{
										Description: `Response`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"version": &schema.Schema{
										Description: `Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"site": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"response": &schema.Schema{
										Description: `Response`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"version": &schema.Schema{
										Description: `Version`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vDeviceFamily, okDeviceFamily := d.GetOk("device_family")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetMembership")
		vvSiteID := vSiteID.(string)
		queryParams1 := dnacentersdkgo.GetMembershipQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		if okDeviceFamily {
			queryParams1.DeviceFamily = vDeviceFamily.(string)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber.(string)
		}

		response1, restyResp1, err := client.Sites.GetMembership(vvSiteID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetMembership", err,
				"Failure at GetMembership, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesGetMembershipItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMembership response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetMembershipItem(item *dnacentersdkgo.ResponseSitesGetMembership) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site"] = flattenSitesGetMembershipItemSite(item.Site)
	respItem["device"] = flattenSitesGetMembershipItemDevice(item.Device)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSitesGetMembershipItemSite(item *dnacentersdkgo.ResponseSitesGetMembershipSite) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = flattenSitesGetMembershipItemSiteResponse(item.Response)
	respItem["version"] = item.Version

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesGetMembershipItemSiteResponse(items *[]dnacentersdkgo.ResponseSitesGetMembershipSiteResponse) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenSitesGetMembershipItemDevice(items *[]dnacentersdkgo.ResponseSitesGetMembershipDevice) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["response"] = flattenSitesGetMembershipItemDeviceResponse(item.Response)
		respItem["version"] = item.Version
		respItem["site_id"] = item.SiteID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetMembershipItemDeviceResponse(items *[]dnacentersdkgo.ResponseSitesGetMembershipDeviceResponse) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
