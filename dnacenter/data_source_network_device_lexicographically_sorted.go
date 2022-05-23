package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceLexicographicallySorted() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the list of values of the first given required parameter. You can use the .* in any value to conduct a
wildcard search. For example, to get all the devices with the management IP address starting with 10.10. , issue the
following request: GET /dna/inten/api/v1/network-device/autocomplete?managementIpAddress=10.10..* It will return the
device management IP addresses that match fully or partially the provided attribute. {[10.10.1.1, 10.10.20.2, …]}.
`,

		ReadContext: dataSourceNetworkDeviceLexicographicallySortedRead,
		Schema: map[string]*schema.Schema{
			"associated_wlc_ip": &schema.Schema{
				Description: `associatedWlcIp query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"collection_interval": &schema.Schema{
				Description: `collectionInterval query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"collection_status": &schema.Schema{
				Description: `collectionStatus query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"error_code": &schema.Schema{
				Description: `errorCode query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"family": &schema.Schema{
				Description: `family query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname": &schema.Schema{
				Description: `hostname query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"management_ip_address": &schema.Schema{
				Description: `managementIpAddress query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"platform_id": &schema.Schema{
				Description: `platformId query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"reachability_failure_reason": &schema.Schema{
				Description: `reachabilityFailureReason query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"reachability_status": &schema.Schema{
				Description: `reachabilityStatus query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"role": &schema.Schema{
				Description: `role query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"role_source": &schema.Schema{
				Description: `roleSource query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"series": &schema.Schema{
				Description: `series query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_type": &schema.Schema{
				Description: `softwareType query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"software_version": &schema.Schema{
				Description: `softwareVersion query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": &schema.Schema{
				Description: `type query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"up_time": &schema.Schema{
				Description: `upTime query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"vrf_name": &schema.Schema{
				Description: `vrfName query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

func dataSourceNetworkDeviceLexicographicallySortedRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vVrfName, okVrfName := d.GetOk("vrf_name")
	vManagementIPAddress, okManagementIPAddress := d.GetOk("management_ip_address")
	vHostname, okHostname := d.GetOk("hostname")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vFamily, okFamily := d.GetOk("family")
	vCollectionStatus, okCollectionStatus := d.GetOk("collection_status")
	vCollectionInterval, okCollectionInterval := d.GetOk("collection_interval")
	vSoftwareVersion, okSoftwareVersion := d.GetOk("software_version")
	vSoftwareType, okSoftwareType := d.GetOk("software_type")
	vReachabilityStatus, okReachabilityStatus := d.GetOk("reachability_status")
	vReachabilityFailureReason, okReachabilityFailureReason := d.GetOk("reachability_failure_reason")
	vErrorCode, okErrorCode := d.GetOk("error_code")
	vPlatformID, okPlatformID := d.GetOk("platform_id")
	vSeries, okSeries := d.GetOk("series")
	vType, okType := d.GetOk("type")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vUpTime, okUpTime := d.GetOk("up_time")
	vRole, okRole := d.GetOk("role")
	vRoleSource, okRoleSource := d.GetOk("role_source")
	vAssociatedWlcIP, okAssociatedWlcIP := d.GetOk("associated_wlc_ip")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute")
		queryParams1 := dnacentersdkgo.GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams{}

		if okVrfName {
			queryParams1.VrfName = vVrfName.(string)
		}
		if okManagementIPAddress {
			queryParams1.ManagementIPAddress = vManagementIPAddress.(string)
		}
		if okHostname {
			queryParams1.Hostname = vHostname.(string)
		}
		if okMacAddress {
			queryParams1.MacAddress = vMacAddress.(string)
		}
		if okFamily {
			queryParams1.Family = vFamily.(string)
		}
		if okCollectionStatus {
			queryParams1.CollectionStatus = vCollectionStatus.(string)
		}
		if okCollectionInterval {
			queryParams1.CollectionInterval = vCollectionInterval.(string)
		}
		if okSoftwareVersion {
			queryParams1.SoftwareVersion = vSoftwareVersion.(string)
		}
		if okSoftwareType {
			queryParams1.SoftwareType = vSoftwareType.(string)
		}
		if okReachabilityStatus {
			queryParams1.ReachabilityStatus = vReachabilityStatus.(string)
		}
		if okReachabilityFailureReason {
			queryParams1.ReachabilityFailureReason = vReachabilityFailureReason.(string)
		}
		if okErrorCode {
			queryParams1.ErrorCode = vErrorCode.(string)
		}
		if okPlatformID {
			queryParams1.PlatformID = vPlatformID.(string)
		}
		if okSeries {
			queryParams1.Series = vSeries.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = vSerialNumber.(string)
		}
		if okUpTime {
			queryParams1.UpTime = vUpTime.(string)
		}
		if okRole {
			queryParams1.Role = vRole.(string)
		}
		if okRoleSource {
			queryParams1.RoleSource = vRoleSource.(string)
		}
		if okAssociatedWlcIP {
			queryParams1.AssociatedWlcIP = vAssociatedWlcIP.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Devices.GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute", err,
				"Failure at GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttributeItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttributeItems(items *dnacentersdkgo.ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute) []map[string]interface{} {
	if items == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = items.Response
	respItem["version"] = items.Version
	return []map[string]interface{}{
		respItem,
	}
}
