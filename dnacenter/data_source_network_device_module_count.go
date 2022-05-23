package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceModuleCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns Module Count
`,

		ReadContext: dataSourceNetworkDeviceModuleCountRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"name_list": &schema.Schema{
				Description: `nameList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"operational_state_code_list": &schema.Schema{
				Description: `operationalStateCodeList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"part_number_list": &schema.Schema{
				Description: `partNumberList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vendor_equipment_type_list": &schema.Schema{
				Description: `vendorEquipmentTypeList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
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

func dataSourceNetworkDeviceModuleCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID := d.Get("device_id")
	vNameList, okNameList := d.GetOk("name_list")
	vVendorEquipmentTypeList, okVendorEquipmentTypeList := d.GetOk("vendor_equipment_type_list")
	vPartNumberList, okPartNumberList := d.GetOk("part_number_list")
	vOperationalStateCodeList, okOperationalStateCodeList := d.GetOk("operational_state_code_list")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetModuleCount")
		queryParams1 := dnacentersdkgo.GetModuleCountQueryParams{}

		queryParams1.DeviceID = vDeviceID.(string)

		if okNameList {
			queryParams1.NameList = interfaceToSliceString(vNameList)
		}
		if okVendorEquipmentTypeList {
			queryParams1.VendorEquipmentTypeList = interfaceToSliceString(vVendorEquipmentTypeList)
		}
		if okPartNumberList {
			queryParams1.PartNumberList = interfaceToSliceString(vPartNumberList)
		}
		if okOperationalStateCodeList {
			queryParams1.OperationalStateCodeList = interfaceToSliceString(vOperationalStateCodeList)
		}

		response1, restyResp1, err := client.Devices.GetModuleCount(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetModuleCount", err,
				"Failure at GetModuleCount, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetModuleCountItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetModuleCount response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetModuleCountItem(item *dnacentersdkgo.ResponseDevicesGetModuleCount) []map[string]interface{} {
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
