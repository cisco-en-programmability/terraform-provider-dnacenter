package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaSiteMemberMember() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- API to get devices that are assigned to a site.
`,

		ReadContext: dataSourceSdaSiteMemberMemberRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Site Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"level": &schema.Schema{
				Description: `level query parameter. Depth of site hierarchy to be considered to list the devices. If the provided value is -1, devices for all child sites will be listed.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of devices to be listed. Default and max supported value is 500
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"member_type": &schema.Schema{
				Description: `memberType query parameter. Member type (This API only supports the 'networkdevice' type)
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Offset/starting index for pagination
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_manager_interface_ip": &schema.Schema{
							Description: `Access Point manager interface IP
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"associated_wlc_ip": &schema.Schema{
							Description: `Associated Wireless Controller IP
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_entity_class": &schema.Schema{
							Description: `Authentication entity class (Internal record)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"auth_entity_id": &schema.Schema{
							Description: `Authentication Entity Id (Internal record)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"boot_date_time": &schema.Schema{
							Description: `Device boot date and time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"collection_interval": &schema.Schema{
							Description: `Device resync interval type (E.g. Global Default)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"collection_interval_value": &schema.Schema{
							Description: `Device resync interval value
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"collection_status": &schema.Schema{
							Description: `Device inventory collection status (E.g. Managed)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"deploy_pending": &schema.Schema{
							Description: `Deploy pending (Internal record)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Device description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_support_level": &schema.Schema{
							Description: `Device support level (E.g. Supported)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
							Description: `Device display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dns_resolved_management_address": &schema.Schema{
							Description: `DNS resolved management address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"family": &schema.Schema{
							Description: `Device family (E.g. Routers)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"hostname": &schema.Schema{
							Description: `Device hostname
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_id": &schema.Schema{
							Description: `Device Id (E.g. 230230)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Device tenant Id (E.g. 64472cc32d3bc1658597669c)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Device UUID (E.g. 48eebb3e-b3fc-4928-a7df-1c80e216f930)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_version": &schema.Schema{
							Description: `Instance version (Internal record)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"interface_count": &schema.Schema{
							Description: `Device interface count
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory_status_detail": &schema.Schema{
							Description: `Device inventory collection status detail
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_device_resync_start_time": &schema.Schema{
							Description: `Last device inventory resync start date and time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last update time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"last_updated": &schema.Schema{
							Description: `Last updated date and time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"line_card_count": &schema.Schema{
							Description: `Line card count
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"line_card_id": &schema.Schema{
							Description: `Line card Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `MAC address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed_atleast_once": &schema.Schema{
							Description: `If device managed atleast once, value will be true otherwise false
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management IP address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_state": &schema.Schema{
							Description: `Device management state (E.g. Managed)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"memory_size": &schema.Schema{
							Description: `Memory size
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"padded_mgmt_ip_address": &schema.Schema{
							Description: `Padded management IP address. Internal record
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"pending_sync_requests_count": &schema.Schema{
							Description: `Pending sync requests count. Internal record
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"platform_id": &schema.Schema{
							Description: `Device platform Id (E.g. CSR1000V)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reachability_failure_reason": &schema.Schema{
							Description: `Device reachability failure reason
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reachability_status": &schema.Schema{
							Description: `Device reachability status (E.g. Reachable)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reasons_for_device_resync": &schema.Schema{
							Description: `Reasons for device resync (E.g. Periodic)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reasons_for_pending_sync_requests": &schema.Schema{
							Description: `Reasons for pending device sync requests
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"role": &schema.Schema{
							Description: `Device role (E.g. BORDER ROUTER)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"role_source": &schema.Schema{
							Description: `Device role source. Internal record
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Device serial Number
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"series": &schema.Schema{
							Description: `Device series
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_contact": &schema.Schema{
							Description: `Device snmp contact. Internal record
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp_location": &schema.Schema{
							Description: `Device snmp location
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_type": &schema.Schema{
							Description: `Device software type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Description: `Device software version
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tag_count": &schema.Schema{
							Description: `Device tag Count
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Device type (E.g. Cisco Cloud Services Router 1000V)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"up_time": &schema.Schema{
							Description: `Device up time (E.g. 112 days, 6:09:13.86)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"uptime_seconds": &schema.Schema{
							Description: `Device uptime in seconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"vendor": &schema.Schema{
							Description: `Vendor (E.g. Cisco)
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

func dataSourceSdaSiteMemberMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vMemberType := d.Get("member_type")
	vLevel, okLevel := d.GetOk("level")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDevicesThatAreAssignedToASite")
		vvID := vID.(string)
		queryParams1 := dnacentersdkgo.GetDevicesThatAreAssignedToASiteQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}
		queryParams1.MemberType = vMemberType.(string)

		if okLevel {
			queryParams1.Level = vLevel.(string)
		}

		response1, restyResp1, err := client.Sites.GetDevicesThatAreAssignedToASite(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDevicesThatAreAssignedToASite", err,
				"Failure at GetDevicesThatAreAssignedToASite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesGetDevicesThatAreAssignedToASiteItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDevicesThatAreAssignedToASite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetDevicesThatAreAssignedToASiteItems(items *[]dnacentersdkgo.ResponseSitesGetDevicesThatAreAssignedToASiteResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["instance_id"] = item.InstanceID
		respItem["auth_entity_id"] = item.AuthEntityID
		respItem["auth_entity_class"] = item.AuthEntityClass
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["deploy_pending"] = item.DeployPending
		respItem["instance_version"] = item.InstanceVersion
		respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
		respItem["associated_wlc_ip"] = item.AssociatedWlcIP
		respItem["boot_date_time"] = item.BootDateTime
		respItem["collection_interval"] = item.CollectionInterval
		respItem["collection_interval_value"] = item.CollectionIntervalValue
		respItem["collection_status"] = item.CollectionStatus
		respItem["description"] = item.Description
		respItem["device_support_level"] = item.DeviceSupportLevel
		respItem["dns_resolved_management_address"] = item.DNSResolvedManagementAddress
		respItem["family"] = item.Family
		respItem["hostname"] = item.Hostname
		respItem["interface_count"] = item.InterfaceCount
		respItem["inventory_status_detail"] = item.InventoryStatusDetail
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["last_updated"] = item.LastUpdated
		respItem["line_card_count"] = item.LineCardCount
		respItem["line_card_id"] = item.LineCardID
		respItem["last_device_resync_start_time"] = item.LastDeviceResyncStartTime
		respItem["mac_address"] = item.MacAddress
		respItem["managed_atleast_once"] = boolPtrToString(item.ManagedAtleastOnce)
		respItem["management_ip_address"] = item.ManagementIPAddress
		respItem["management_state"] = item.ManagementState
		respItem["memory_size"] = item.MemorySize
		respItem["padded_mgmt_ip_address"] = item.PaddedMgmtIPAddress
		respItem["pending_sync_requests_count"] = item.PendingSyncRequestsCount
		respItem["platform_id"] = item.PlatformID
		respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
		respItem["reachability_status"] = item.ReachabilityStatus
		respItem["reasons_for_device_resync"] = item.ReasonsForDeviceResync
		respItem["reasons_for_pending_sync_requests"] = item.ReasonsForPendingSyncRequests
		respItem["role"] = item.Role
		respItem["role_source"] = item.RoleSource
		respItem["serial_number"] = item.SerialNumber
		respItem["series"] = item.Series
		respItem["snmp_contact"] = item.SNMPContact
		respItem["snmp_location"] = item.SNMPLocation
		respItem["software_type"] = item.SoftwareType
		respItem["software_version"] = item.SoftwareVersion
		respItem["tag_count"] = item.TagCount
		respItem["type"] = item.Type
		respItem["up_time"] = item.UpTime
		respItem["uptime_seconds"] = item.UptimeSeconds
		respItem["vendor"] = item.Vendor
		respItem["display_name"] = item.DisplayName
		respItems = append(respItems, respItem)
	}
	return respItems
}
