package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseDeviceLicenseSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Show license summary of device(s).
`,

		ReadContext: dataSourceLicenseDeviceLicenseSummaryRead,
		Schema: map[string]*schema.Schema{
			"device_type": &schema.Schema{
				Description: `device_type query parameter. Type of device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuid": &schema.Schema{
				Description: `device_uuid query parameter. Id of device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"dna_level": &schema.Schema{
				Description: `dna_level query parameter. Device Cisco DNA license level
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Required:    true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Sorting order
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"page_number": &schema.Schema{
				Description: `page_number query parameter. Page number of response
`,
				Type:     schema.TypeFloat,
				Required: true,
			},
			"registration_status": &schema.Schema{
				Description: `registration_status query parameter. Smart license registration status of device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"smart_account_id": &schema.Schema{
				Description: `smart_account_id query parameter. Id of smart account
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sort_by query parameter. Sort result by field
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_account_name": &schema.Schema{
				Description: `virtual_account_name query parameter. Name of virtual account
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authorization_status": &schema.Schema{
							Description: `Smart license authorization status of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"customer_tag1": &schema.Schema{
							Description: `Customer Tag1 set on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"customer_tag2": &schema.Schema{
							Description: `Customer Tag2 set on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"customer_tag3": &schema.Schema{
							Description: `Customer Tag3 set on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"customer_tag4": &schema.Schema{
							Description: `Customer Tag4 set on device
`,
							Type:     schema.TypeString,
							Computed: true,
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

						"hsec_status": &schema.Schema{
							Description: `HSEC status of the device
`,
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

						"is_performance_allowed": &schema.Schema{
							Description: `Is performance license available
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_wireless": &schema.Schema{
							Description: `Is device wireless controller
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_wireless_capable": &schema.Schema{
							Description: `Is device wireless capable
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_successful_rum_usage_upload_time": &schema.Schema{
							Description: `Last successful rum usage upload time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_time": &schema.Schema{
							Description: `Time when license information was collected from device
`,
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
							Description: `Device Network license level
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"performance_license": &schema.Schema{
							Description: `Is performance license enabled
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"registration_status": &schema.Schema{
							Description: `Smart license registration status of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reservation_status": &schema.Schema{
							Description: `Smart license reservation status
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

						"sle_auth_code": &schema.Schema{
							Description: `SLE Authcode installed or not installed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sle_state": &schema.Schema{
							Description: `SLE state on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"smart_account_name": &schema.Schema{
							Description: `Name of smart account
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

						"throughput_level": &schema.Schema{
							Description: `Throughput level on device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"total_access_point_count": &schema.Schema{
							Description: `Total number of Access Points connected
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"virtual_account_name": &schema.Schema{
							Description: `Name of virtual account
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wireless_capable_dna_license": &schema.Schema{
							Description: `Wireless Cisco DNA license value
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wireless_capable_network_license": &schema.Schema{
							Description: `Wireless Cisco Network license value
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

func dataSourceLicenseDeviceLicenseSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPageNumber := d.Get("page_number")
	vOrder := d.Get("order")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vDnaLevel, okDnaLevel := d.GetOk("dna_level")
	vDeviceType, okDeviceType := d.GetOk("device_type")
	vLimit := d.Get("limit")
	vRegistrationStatus, okRegistrationStatus := d.GetOk("registration_status")
	vVirtualAccountName, okVirtualAccountName := d.GetOk("virtual_account_name")
	vSmartAccountID, okSmartAccountID := d.GetOk("smart_account_id")
	vDeviceUUID, okDeviceUUID := d.GetOk("device_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeviceLicenseSummary")
		queryParams1 := dnacentersdkgo.DeviceLicenseSummaryQueryParams{}

		queryParams1.PageNumber = vPageNumber.(float64)

		queryParams1.Order = vOrder.(string)

		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okDnaLevel {
			queryParams1.DnaLevel = vDnaLevel.(string)
		}
		if okDeviceType {
			queryParams1.DeviceType = vDeviceType.(string)
		}
		queryParams1.Limit = vLimit.(float64)

		if okRegistrationStatus {
			queryParams1.RegistrationStatus = vRegistrationStatus.(string)
		}
		if okVirtualAccountName {
			queryParams1.VirtualAccountName = vVirtualAccountName.(string)
		}
		if okSmartAccountID {
			queryParams1.SmartAccountID = vSmartAccountID.(float64)
		}
		if okDeviceUUID {
			queryParams1.DeviceUUID = vDeviceUUID.(string)
		}

		response1, restyResp1, err := client.Licenses.DeviceLicenseSummary(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeviceLicenseSummary", err,
				"Failure at DeviceLicenseSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLicensesDeviceLicenseSummaryItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeviceLicenseSummary response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesDeviceLicenseSummaryItems(items *[]dnacentersdkgo.ResponseLicensesDeviceLicenseSummaryResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["authorization_status"] = item.AuthorizationStatus
		respItem["last_updated_time"] = item.LastUpdatedTime
		respItem["is_performance_allowed"] = boolPtrToString(item.IsPerformanceAllowed)
		respItem["sle_auth_code"] = item.SleAuthCode
		respItem["throughput_level"] = item.ThroughputLevel
		respItem["hsec_status"] = item.HsecStatus
		respItem["device_uuid"] = item.DeviceUUID
		respItem["site"] = item.Site
		respItem["total_access_point_count"] = item.TotalAccessPointCount
		respItem["model"] = item.Model
		respItem["is_wireless_capable"] = boolPtrToString(item.IsWirelessCapable)
		respItem["registration_status"] = item.RegistrationStatus
		respItem["sle_state"] = item.SleState
		respItem["performance_license"] = item.PerformanceLicense
		respItem["license_mode"] = item.LicenseMode
		respItem["is_license_expired"] = boolPtrToString(item.IsLicenseExpired)
		respItem["software_version"] = item.SoftwareVersion
		respItem["reservation_status"] = item.ReservationStatus
		respItem["is_wireless"] = boolPtrToString(item.IsWireless)
		respItem["network_license"] = item.NetworkLicense
		respItem["evaluation_license_expiry"] = item.EvaluationLicenseExpiry
		respItem["wireless_capable_network_license"] = item.WirelessCapableNetworkLicense
		respItem["device_name"] = item.DeviceName
		respItem["device_type"] = item.DeviceType
		respItem["dna_level"] = item.DnaLevel
		respItem["virtual_account_name"] = item.VirtualAccountName
		respItem["last_successful_rum_usage_upload_time"] = item.LastSuccessfulRumUsageUploadTime
		respItem["ip_address"] = item.IPAddress
		respItem["wireless_capable_dna_license"] = item.WirelessCapableDnaLicense
		respItem["mac_address"] = item.MacAddress
		respItem["customer_tag1"] = item.CustomerTag1
		respItem["customer_tag2"] = item.CustomerTag2
		respItem["customer_tag3"] = item.CustomerTag3
		respItem["customer_tag4"] = item.CustomerTag4
		respItem["smart_account_name"] = item.SmartAccountName
		respItems = append(respItems, respItem)
	}
	return respItems
}
