package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceChassisDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns chassis details for given device ID
`,

		ReadContext: dataSourceNetworkDeviceChassisDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"assembly_number": &schema.Schema{
							Description: `Assembly Number of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"assembly_revision": &schema.Schema{
							Description: `Assembly Revision of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"containment_entity": &schema.Schema{
							Description: `Containment Entity of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"entity_physical_index": &schema.Schema{
							Description: `Entity Physical Index of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"hardware_version": &schema.Schema{
							Description: `Hardware Version of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `ID of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_field_replaceable": &schema.Schema{
							Description: `To mention if field is replaceable
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_reporting_alarms_allowed": &schema.Schema{
							Description: `To mention if reporting alarms are allowed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"manufacturer": &schema.Schema{
							Description: `Manufacturer of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"part_number": &schema.Schema{
							Description: `Part Number of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial Number of the chassis
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vendor_equipment_type": &schema.Schema{
							Description: `Vendor Equipment Type of the chassis
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

func dataSourceNetworkDeviceChassisDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID := d.Get("device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetChassisDetailsForDevice")
		vvDeviceID := vDeviceID.(string)

		response1, restyResp1, err := client.Devices.GetChassisDetailsForDevice(vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetChassisDetailsForDevice", err,
				"Failure at GetChassisDetailsForDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetChassisDetailsForDeviceItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetChassisDetailsForDevice response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetChassisDetailsForDeviceItems(items *[]dnacentersdkgo.ResponseDevicesGetChassisDetailsForDeviceResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["assembly_number"] = item.AssemblyNumber
		respItem["assembly_revision"] = item.AssemblyRevision
		respItem["containment_entity"] = item.ContainmentEntity
		respItem["description"] = item.Description
		respItem["entity_physical_index"] = item.EntityPhysicalIndex
		respItem["hardware_version"] = item.HardwareVersion
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["is_field_replaceable"] = item.IsFieldReplaceable
		respItem["is_reporting_alarms_allowed"] = item.IsReportingAlarmsAllowed
		respItem["manufacturer"] = item.Manufacturer
		respItem["name"] = item.Name
		respItem["part_number"] = item.PartNumber
		respItem["serial_number"] = item.SerialNumber
		respItem["vendor_equipment_type"] = item.VendorEquipmentType
		respItems = append(respItems, respItem)
	}
	return respItems
}
