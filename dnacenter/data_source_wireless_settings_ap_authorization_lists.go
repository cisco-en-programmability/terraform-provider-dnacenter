package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsApAuthorizationLists() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieves the AP Authorization Lists that are created in the Catalyst Centre network Design for wireless. If an AP
Authorization List name is given as query parameter, then returns respective AP Authorization List details including
Local and/or Remote authorization.
`,

		ReadContext: dataSourceWirelessSettingsApAuthorizationListsRead,
		Schema: map[string]*schema.Schema{
			"ap_authorization_list_name": &schema.Schema{
				Description: `apAuthorizationListName query parameter. Employ this query parameter to obtain the details of the AP Authorization List corresponding to the provided apAuthorizationListName.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page. The first record is numbered 1.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_authorization_list_name": &schema.Schema{
							Description: `Ap Authorization List Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"local_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_mac_entries": &schema.Schema{
										Description: `AP Mac Addresses`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ap_serial_number_entries": &schema.Schema{
										Description: `AP Serial Number Entries`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"remote_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_servers": &schema.Schema{
										Description: `AAA Servers`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"authorize_ap_with_mac": &schema.Schema{
										Description: `Authorize AP With Mac`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"authorize_ap_with_serial_number": &schema.Schema{
										Description: `Authorize AP With Serial Number`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceWirelessSettingsApAuthorizationListsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vApAuthorizationListName, okApAuthorizationListName := d.GetOk("ap_authorization_list_name")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApAuthorizationLists")
		queryParams1 := dnacentersdkgo.GetApAuthorizationListsQueryParams{}

		if okApAuthorizationListName {
			queryParams1.ApAuthorizationListName = vApAuthorizationListName.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}

		response1, restyResp1, err := client.Wireless.GetApAuthorizationLists(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApAuthorizationLists", err,
				"Failure at GetApAuthorizationLists, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetApAuthorizationListsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApAuthorizationLists response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetApAuthorizationListsItem(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["ap_authorization_list_name"] = item.ApAuthorizationListName
	respItem["local_authorization"] = flattenWirelessGetApAuthorizationListsItemLocalAuthorization(item.LocalAuthorization)
	respItem["remote_authorization"] = flattenWirelessGetApAuthorizationListsItemRemoteAuthorization(item.RemoteAuthorization)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApAuthorizationListsItemLocalAuthorization(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListsResponseLocalAuthorization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ap_mac_entries"] = item.ApMacEntries
	respItem["ap_serial_number_entries"] = item.ApSerialNumberEntries

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApAuthorizationListsItemRemoteAuthorization(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListsResponseRemoteAuthorization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_servers"] = item.AAAServers
	respItem["authorize_ap_with_mac"] = boolPtrToString(item.AuthorizeApWithMac)
	respItem["authorize_ap_with_serial_number"] = boolPtrToString(item.AuthorizeApWithSerialNumber)

	return []map[string]interface{}{
		respItem,
	}

}
