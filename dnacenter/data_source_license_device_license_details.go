package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseDeviceLicenseDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Get detailed license information of a device.
`,

		ReadContext: dataSourceLicenseDeviceLicenseDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `device_uuid path parameter. Id of device
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"access_points": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_type": &schema.Schema{
										Description: `Type of access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"count": &schema.Schema{
										Description: `Number of access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"chassis_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"board_serial_number": &schema.Schema{
										Description: `Board serial number
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"modules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id of module
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"module_name": &schema.Schema{
													Description: `Name of module
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"module_type": &schema.Schema{
													Description: `Type of module
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"serial_number": &schema.Schema{
													Description: `Serial number of module
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"port": &schema.Schema{
										Description: `Number of port
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"supervisor_cards": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"serial_number": &schema.Schema{
													Description: `Serial number of supervisor card
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"status": &schema.Schema{
													Description: `Status of supervisor card
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"supervisor_card_type": &schema.Schema{
													Description: `Type of supervisor card
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"device_name": &schema.Schema{
							Description: `Name of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_type": &schema.Schema{
							Description: `Type of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_uuid": &schema.Schema{
							Description: `Id of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dna_level": &schema.Schema{
							Description: `Device Cisco DNA license level
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"evaluation_license_expiry": &schema.Schema{
							Description: `Evaluation period expiry date
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"feature_license": &schema.Schema{
							Description: `Name of feature licenses
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"has_sup_cards": &schema.Schema{
							Description: `Whether device has supervisor cards
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Description: `IP address of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_license_expired": &schema.Schema{
							Description: `Is device license expired
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_stacked_device": &schema.Schema{
							Description: `Is Stacked Device
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"license_mode": &schema.Schema{
							Description: `Mode of license
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `MAC address of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"model": &schema.Schema{
							Description: `Model of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_license": &schema.Schema{
							Description: `Device network license level
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site": &schema.Schema{
							Description: `Site of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sntc_status": &schema.Schema{
							Description: `Valid if device is covered under license contract and invalid if not covered, otherwise unknown.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Description: `Software image version of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"stacked_devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mac_address": &schema.Schema{
										Description: `Stack mac address
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"role": &schema.Schema{
										Description: `Chassis role
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"serial_number": &schema.Schema{
										Description: `Chassis serial number
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"udi": &schema.Schema{
							Description: `Unique Device Identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_account_name": &schema.Schema{
							Description: `Name of virtual account
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

func dataSourceLicenseDeviceLicenseDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeviceLicenseDetails")
		vvDeviceUUID := vDeviceUUID.(string)

		response1, restyResp1, err := client.Licenses.DeviceLicenseDetails(vvDeviceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeviceLicenseDetails", err,
				"Failure at DeviceLicenseDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLicensesDeviceLicenseDetailsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeviceLicenseDetails response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesDeviceLicenseDetailsItems(items *[]dnacentersdkgo.ResponseLicensesDeviceLicenseDetailsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_uuid"] = item.DeviceUUID
		respItem["site"] = item.Site
		respItem["model"] = item.Model
		respItem["license_mode"] = item.LicenseMode
		respItem["is_license_expired"] = boolPtrToString(item.IsLicenseExpired)
		respItem["software_version"] = item.SoftwareVersion
		respItem["network_license"] = item.NetworkLicense
		respItem["evaluation_license_expiry"] = item.EvaluationLicenseExpiry
		respItem["device_name"] = item.DeviceName
		respItem["device_type"] = item.DeviceType
		respItem["dna_level"] = item.DnaLevel
		respItem["virtual_account_name"] = item.VirtualAccountName
		respItem["ip_address"] = item.IPAddress
		respItem["mac_address"] = item.MacAddress
		respItem["sntc_status"] = item.SntcStatus
		respItem["feature_license"] = item.FeatureLicense
		respItem["has_sup_cards"] = boolPtrToString(item.HasSupCards)
		respItem["udi"] = item.Udi
		respItem["stacked_devices"] = flattenLicensesDeviceLicenseDetailsItemsStackedDevices(item.StackedDevices)
		respItem["is_stacked_device"] = boolPtrToString(item.IsStackedDevice)
		respItem["access_points"] = flattenLicensesDeviceLicenseDetailsItemsAccessPoints(item.AccessPoints)
		respItem["chassis_details"] = flattenLicensesDeviceLicenseDetailsItemsChassisDetails(item.ChassisDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesDeviceLicenseDetailsItemsStackedDevices(items *[]dnacentersdkgo.ResponseLicensesDeviceLicenseDetailsResponseStackedDevices) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mac_address"] = item.MacAddress
		respItem["id"] = item.ID
		respItem["role"] = item.Role
		respItem["serial_number"] = item.SerialNumber
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesDeviceLicenseDetailsItemsAccessPoints(items *[]dnacentersdkgo.ResponseLicensesDeviceLicenseDetailsResponseAccessPoints) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ap_type"] = item.ApType
		respItem["count"] = item.Count
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesDeviceLicenseDetailsItemsChassisDetails(item *dnacentersdkgo.ResponseLicensesDeviceLicenseDetailsResponseChassisDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["board_serial_number"] = item.BoardSerialNumber
	respItem["modules"] = flattenLicensesDeviceLicenseDetailsItemsChassisDetailsModules(item.Modules)
	respItem["supervisor_cards"] = flattenLicensesDeviceLicenseDetailsItemsChassisDetailsSupervisorCards(item.SupervisorCards)
	respItem["port"] = item.Port

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesDeviceLicenseDetailsItemsChassisDetailsModules(items *[]dnacentersdkgo.ResponseLicensesDeviceLicenseDetailsResponseChassisDetailsModules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["module_type"] = item.ModuleType
		respItem["module_name"] = item.ModuleName
		respItem["serial_number"] = item.SerialNumber
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesDeviceLicenseDetailsItemsChassisDetailsSupervisorCards(items *[]dnacentersdkgo.ResponseLicensesDeviceLicenseDetailsResponseChassisDetailsSupervisorCards) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["serial_number"] = item.SerialNumber
		respItem["supervisor_card_type"] = item.SupervisorCardType
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
