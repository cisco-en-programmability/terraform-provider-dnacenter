package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsInterfaces() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get all Interfaces

- This data source allows the user to get an interface by ID
`,

		ReadContext: dataSourceWirelessSettingsInterfacesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Interface ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_name": &schema.Schema{
				Description: `interfaceName query parameter. Interface Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page. The first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Description: `vlanId query parameter. Vlan Id
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Interface ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_name": &schema.Schema{
							Description: `Interface Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Description: `VLAN ID
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Interface ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_name": &schema.Schema{
							Description: `Interface Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Description: `VLAN ID
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessSettingsInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vInterfaceName, okInterfaceName := d.GetOk("interface_name")
	vVLANID, okVLANID := d.GetOk("vlan_id")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset, okInterfaceName, okVLANID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetInterfaces")
		queryParams1 := dnacentersdkgo.GetInterfacesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okInterfaceName {
			queryParams1.InterfaceName = vInterfaceName.(string)
		}
		if okVLANID {
			queryParams1.VLANID = vVLANID.(float64)
		}

		response1, restyResp1, err := client.Wireless.GetInterfaces(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetInterfaces", err,
				"Failure at GetInterfaces, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetInterfacesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInterfaces response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetInterfaceByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.Wireless.GetInterfaceByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetInterfaceByID", err,
				"Failure at GetInterfaceByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenWirelessGetInterfaceByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInterfaceByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetInterfacesItems(items *[]dnacentersdkgo.ResponseWirelessGetInterfacesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface_name"] = item.InterfaceName
		respItem["vlan_id"] = item.VLANID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetInterfaceByIDItem(item *dnacentersdkgo.ResponseWirelessGetInterfaceByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interface_name"] = item.InterfaceName
	respItem["vlan_id"] = item.VLANID
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
