package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEoxStatusDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on EoX.

- Retrieves EoX status for all devices in the network

- Retrieves EoX details for a device
`,

		ReadContext: dataSourceEoxStatusDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device instance UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"alert_count": &schema.Schema{
							Description: `Alert Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"comments": &schema.Schema{
							Description: `Comments`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"device_id": &schema.Schema{
							Description: `Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"eox_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bulletin_headline": &schema.Schema{
										Description: `Bulletin Headline`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bulletin_number": &schema.Schema{
										Description: `Bulletin Number`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"bulletin_url": &schema.Schema{
										Description: `Bulletin U R L`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"end_of_hardware_new_service_attachment_date": &schema.Schema{
										Description: `End Of Hardware New Service Attachment Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_hardware_service_contract_renewal_date": &schema.Schema{
										Description: `End Of Hardware Service Contract Renewal Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_last_hardware_ship_date": &schema.Schema{
										Description: `End Of Last Hardware Ship Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_life_date": &schema.Schema{
										Description: `End Of Life Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_life_external_announcement_date": &schema.Schema{
										Description: `End Of Life External Announcement Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_sale_date": &schema.Schema{
										Description: `End Of Sale Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_signature_releases_date": &schema.Schema{
										Description: `End Of Signature Releases Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_software_maintenance_releases_date": &schema.Schema{
										Description: `End Of Software Maintenance Releases Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_software_vulnerability_or_security_support_date": &schema.Schema{
										Description: `End Of Software Vulnerability Or Security Support Date`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"end_of_software_vulnerability_or_security_support_date_hw": &schema.Schema{
										Description: `End Of Software Vulnerability Or Security Support Date Hw`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"eox_alert_type": &schema.Schema{
										Description: `Eox Alert Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"last_date_of_support": &schema.Schema{
										Description: `Last Date Of Support`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"last_scan_time": &schema.Schema{
							Description: `Last Scan Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"scan_status": &schema.Schema{
							Description: `Scan Status`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `Alert Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"device_id": &schema.Schema{
							Description: `Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"last_scan_time": &schema.Schema{
							Description: `Last Scan Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"scan_status": &schema.Schema{
							Description: `Scan Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"summary": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"eox_type": &schema.Schema{
										Description: `Eox Type`,
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

func dataSourceEoxStatusDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID, okDeviceID := d.GetOk("device_id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okDeviceID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEoXStatusForAllDevices")

		response1, restyResp1, err := client.EoX.GetEoXStatusForAllDevices()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetEoXStatusForAllDevices", err,
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
				"Failure when executing GetEoXDetailsPerDevice", err,
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
		respItem["eox_type"] = item.EoxType
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
	respItem["eox_details"] = flattenEoXGetEoXDetailsPerDeviceItemEoxDetails(item.EoxDetails)
	respItem["scan_status"] = item.ScanStatus
	respItem["comments"] = flattenEoXGetEoXDetailsPerDeviceItemComments(item.Comments)
	respItem["last_scan_time"] = item.LastScanTime
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEoXGetEoXDetailsPerDeviceItemEoxDetails(items *[]dnacentersdkgo.ResponseEoXGetEoXDetailsPerDeviceResponseEoxDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bulletin_headline"] = item.BulletinHeadline
		respItem["bulletin_number"] = item.BulletinNumber
		respItem["bulletin_url"] = item.BulletinURL
		respItem["end_of_hardware_new_service_attachment_date"] = item.EndOfHardwareNewServiceAttachmentDate
		respItem["end_of_hardware_service_contract_renewal_date"] = item.EndOfHardwareServiceContractRenewalDate
		respItem["end_of_last_hardware_ship_date"] = item.EndOfLastHardwareShipDate
		respItem["end_of_life_date"] = item.EndOfLifeDate
		respItem["end_of_life_external_announcement_date"] = item.EndOfLifeExternalAnnouncementDate
		respItem["end_of_sale_date"] = item.EndOfSaleDate
		respItem["end_of_signature_releases_date"] = item.EndOfSignatureReleasesDate
		respItem["end_of_software_vulnerability_or_security_support_date"] = item.EndOfSoftwareVulnerabilityOrSecuritySupportDate
		respItem["end_of_software_vulnerability_or_security_support_date_hw"] = item.EndOfSoftwareVulnerabilityOrSecuritySupportDateHw
		respItem["end_of_software_maintenance_releases_date"] = item.EndOfSoftwareMaintenanceReleasesDate
		respItem["eox_alert_type"] = item.EoxAlertType
		respItem["last_date_of_support"] = item.LastDateOfSupport
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenEoXGetEoXDetailsPerDeviceItemComments(items *[]dnacentersdkgo.ResponseEoXGetEoXDetailsPerDeviceResponseComments) []interface{} {
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
