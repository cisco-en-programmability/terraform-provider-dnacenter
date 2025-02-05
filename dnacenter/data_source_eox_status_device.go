package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEoXStatusDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on EoX.

- Retrieves EoX status for all devices in the network

- Retrieves EoX details for a device
`,

		ReadContext: dataSourceEoXStatusDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device instance UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page, the first record is numbered 1
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"alert_count": &schema.Schema{
							Description: `Number of EoX alerts on the network device
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"comments": &schema.Schema{
							Description: `More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"device_id": &schema.Schema{
							Description: `Device instance UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"eox_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bulletin_headline": &schema.Schema{
										Description: `Title of the EoX bulletin
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"bulletin_name": &schema.Schema{
										Description: `Name of the EoX bulletin
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"bulletin_number": &schema.Schema{
										Description: `Identifier of the EoX bulletin. Usually the same as name.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"bulletin_pid": &schema.Schema{
										Description: `The part number for the EoX alert. eg:- PWR-C1-1100WAC
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"bulletin_url": &schema.Schema{
										Description: `URL where the EoX bulletin is posted
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_hardware_new_service_attachment_date": &schema.Schema{
										Description: `For equipment and software that is not covered by a service-and-support contract, this is the last date to order a new service-and-support contract or add the equipment and/or software to an existing service-and-support contract
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_hardware_service_contract_renewal_date": &schema.Schema{
										Description: `The last date to extend or renew a service contract for the product
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_last_hardware_ship_date": &schema.Schema{
										Description: `The last-possible ship date that can be requested of Cisco and/or its contract manufacturers
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_life_date": &schema.Schema{
										Description: `The last date to receive applicable service and support for the product as entitled by active service contracts or by warranty terms and conditions. This will be populated for software alerts only.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_life_external_announcement_date": &schema.Schema{
										Description: `The date the document that announces the end-of-sale and end-of-life of a product is distributed to the general public
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_sale_date": &schema.Schema{
										Description: `The last date to order the product through Cisco point-of-sale mechanisms
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_signature_releases_date": &schema.Schema{
										Description: `The date after which there will be no more signature update release for the product
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_software_maintenance_releases_date": &schema.Schema{
										Description: `The last date that Cisco Engineering may release any final software maintenance releases or bug fixes for the product
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_software_vulnerability_or_security_support_date": &schema.Schema{
										Description: `The last date that Cisco Engineering may release bug fixes for Vulnerability or Security issues for the product. This will be populated for software alerts only.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"end_of_software_vulnerability_or_security_support_date_hw": &schema.Schema{
										Description: `The last date that Cisco Engineering may release bug fixes for Vulnerability or Security issues for the product. This will be populated for hardware or module alerts only.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"eox_physical_type": &schema.Schema{
										Description: `The type of part for EoX alert. eg:- Power Supply, Chassis, Fan etc.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"eox_alert_type": &schema.Schema{
										Description: `Type of EoX alert
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_date_of_support": &schema.Schema{
										Description: `The last date to receive applicable service and support for the product as entitled by active service contracts or by warranty terms and conditions. This will be populated for hardware and module alerts only.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the EoX alert. Every EoX announcement has a unique name. ie:- EOL13873
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"last_scan_time": &schema.Schema{
							Description: `Time at which the network device was scanned. The representation is unix time.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"scan_status": &schema.Schema{
							Description: `Status of the scan performed on the network device
`,
							Type:     schema.TypeString,
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

						"alert_count": &schema.Schema{
							Description: `Number of EoX alerts on the network device
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"comments": &schema.Schema{
							Description: `More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"device_id": &schema.Schema{
							Description: `Device instance UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_scan_time": &schema.Schema{
							Description: `Time at which the network device was scanned. The representation is unix time.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"scan_status": &schema.Schema{
							Description: `Status of the scan performed on the network device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"summary": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"eox_type": &schema.Schema{
										Description: `Type of EoX Alert
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
		},
	}
}

func dataSourceEoXStatusDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vDeviceID, okDeviceID := d.GetOk("device_id")

	method1 := []bool{okLimit, okOffset}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okDeviceID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEoxStatusForAllDevices")
		queryParams1 := dnacentersdkgo.GetEoXStatusForAllDevicesQueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		response1, restyResp1, err := client.EoX.GetEoXStatusForAllDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetEoXStatusForAllDevices", err,
				"Failure at GetEoXStatusForAllDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenEoXGetEoXStatusForAllDevicesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEoXStatusForAllDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetEoXDetailsPerDevice")
		vvDeviceID := vDeviceID.(string)

		response2, restyResp2, err := client.EoX.GetEoXDetailsPerDevice(vvDeviceID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetEoXDetailsPerDevice", err,
				"Failure at GetEoXDetailsPerDevice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenEoXGetEoXDetailsPerDeviceItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEoXDetailsPerDevice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEoXGetEoXStatusForAllDevicesItems(items *[]dnacentersdkgo.ResponseEoXGetEoXStatusForAllDevicesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_id"] = item.DeviceID
		respItem["alert_count"] = item.AlertCount
		respItem["summary"] = flattenEoXGetEoXStatusForAllDevicesItemsSummary(item.Summary)
		respItem["scan_status"] = item.ScanStatus
		respItem["comments"] = item.Comments
		respItem["last_scan_time"] = item.LastScanTime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEoXGetEoXStatusForAllDevicesItemsSummary(items *[]dnacentersdkgo.ResponseEoXGetEoXStatusForAllDevicesResponseSummary) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["eox_type"] = item.EoXType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEoXGetEoXDetailsPerDeviceItem(item *dnacentersdkgo.ResponseEoXGetEoXDetailsPerDeviceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_id"] = item.DeviceID
	respItem["alert_count"] = item.AlertCount
	respItem["eox_details"] = flattenEoXGetEoXDetailsPerDeviceItemEoXDetails(item.EoXDetails)
	respItem["scan_status"] = item.ScanStatus
	respItem["comments"] = item.Comments
	respItem["last_scan_time"] = item.LastScanTime
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEoXGetEoXDetailsPerDeviceItemEoXDetails(items *[]dnacentersdkgo.ResponseEoXGetEoXDetailsPerDeviceResponseEoXDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["bulletin_headline"] = item.BulletinHeadline
		respItem["bulletin_name"] = item.BulletinName
		respItem["bulletin_number"] = item.BulletinNumber
		respItem["bulletin_url"] = item.BulletinURL
		respItem["end_of_hardware_new_service_attachment_date"] = item.EndOfHardwareNewServiceAttachmentDate
		respItem["end_of_hardware_service_contract_renewal_date"] = item.EndOfHardwareServiceContractRenewalDate
		respItem["end_of_last_hardware_ship_date"] = item.EndOfLastHardwareShipDate
		respItem["end_of_life_external_announcement_date"] = item.EndOfLifeExternalAnnouncementDate
		respItem["end_of_signature_releases_date"] = item.EndOfSignatureReleasesDate
		respItem["end_of_software_vulnerability_or_security_support_date"] = item.EndOfSoftwareVulnerabilityOrSecuritySupportDate
		respItem["end_of_software_vulnerability_or_security_support_date_hw"] = item.EndOfSoftwareVulnerabilityOrSecuritySupportDateHw
		respItem["end_of_sale_date"] = item.EndOfSaleDate
		respItem["end_of_life_date"] = item.EndOfLifeDate
		respItem["last_date_of_support"] = item.LastDateOfSupport
		respItem["end_of_software_maintenance_releases_date"] = item.EndOfSoftwareMaintenanceReleasesDate
		respItem["eox_alert_type"] = item.EoXAlertType
		respItem["eox_physical_type"] = item.EoXPhysicalType
		respItem["bulletin_pid"] = item.BulletinPID
		respItems = append(respItems, respItem)
	}
	return respItems
}
