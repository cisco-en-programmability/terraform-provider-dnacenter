package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceEquipment() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Return all types of equipment details like PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other and SFP for
the Given device.
`,

		ReadContext: dataSourceNetworkDeviceEquipmentRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Type value can be PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other, SFP. If no type is mentioned, All equipments are fetched for the device.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"manufacturer": &schema.Schema{
							Description: `Manufacturer`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"operational_state_code": &schema.Schema{
							Description: `Operational State Code`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"product_id": &schema.Schema{
							Description: `Product Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial Number`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"vendor_equipment_type": &schema.Schema{
							Description: `Vendor Equipment Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceEquipmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")
	vType, okType := d.GetOk("type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTheDetailsOfPhysicalComponentsOfTheGivenDevice")
		vvDeviceUUID := vDeviceUUID.(string)
		queryParams1 := dnacentersdkgo.GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams{}

		if okType {
			queryParams1.Type = vType.(string)
		}

		response1, restyResp1, err := client.Devices.GetTheDetailsOfPhysicalComponentsOfTheGivenDevice(vvDeviceUUID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheDetailsOfPhysicalComponentsOfTheGivenDevice", err,
				"Failure at GetTheDetailsOfPhysicalComponentsOfTheGivenDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheDetailsOfPhysicalComponentsOfTheGivenDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceItems(items *[]dnacentersdkgo.ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["operational_state_code"] = item.OperationalStateCode
		respItem["product_id"] = item.ProductID
		respItem["serial_number"] = item.SerialNumber
		respItem["vendor_equipment_type"] = item.VendorEquipmentType
		respItem["description"] = item.Description
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["name"] = item.Name
		respItem["manufacturer"] = item.Manufacturer
		respItems = append(respItems, respItem)
	}
	return respItems
}
