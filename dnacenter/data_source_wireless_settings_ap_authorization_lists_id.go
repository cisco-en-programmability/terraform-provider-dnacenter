package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsApAuthorizationListsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get an AP Authorization List by AP Authorization List ID that captured in wireless
settings design.
`,

		ReadContext: dataSourceWirelessSettingsApAuthorizationListsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. AP Authorization List ID
`,
				Type:     schema.TypeString,
				Required: true,
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

func dataSourceWirelessSettingsApAuthorizationListsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApAuthorizationListByID")
		vvID := vID.(string)

		response1, restyResp1, err := client.Wireless.GetApAuthorizationListByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApAuthorizationListByID", err,
				"Failure at GetApAuthorizationListByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetApAuthorizationListByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApAuthorizationListByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetApAuthorizationListByIDItem(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["ap_authorization_list_name"] = item.ApAuthorizationListName
	respItem["local_authorization"] = flattenWirelessGetApAuthorizationListByIDItemLocalAuthorization(item.LocalAuthorization)
	respItem["remote_authorization"] = flattenWirelessGetApAuthorizationListByIDItemRemoteAuthorization(item.RemoteAuthorization)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApAuthorizationListByIDItemLocalAuthorization(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListByIDResponseLocalAuthorization) []map[string]interface{} {
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

func flattenWirelessGetApAuthorizationListByIDItemRemoteAuthorization(item *dnacentersdkgo.ResponseWirelessGetApAuthorizationListByIDResponseRemoteAuthorization) []map[string]interface{} {
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
