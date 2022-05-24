package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkDeviceList() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Devices.

- Adds the device with given credential

- Sync the devices provided as input
`,

		CreateContext: resourceNetworkDeviceListCreate,
		ReadContext:   resourceNetworkDeviceListRead,
		UpdateContext: resourceNetworkDeviceListUpdate,
		DeleteContext: resourceNetworkDeviceListDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli_transport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"compute_device": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"enable_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extended_discovery_info": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"http_secure": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"http_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"meraki_org_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"netconf_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_auth_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_auth_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_passphrase": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_priv_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_ro_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_rw_community": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_retry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"snmp_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"snmp_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"snmp_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_mgmt_ipaddress_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"exist_mgmt_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"new_mgmt_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"user_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDeviceListCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkDeviceListAddDevice2(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	resp1, restyResp1, err := client.Devices.AddDevice2(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddDevice2", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddDevice2", err))
		return diags
	}
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkDeviceListRead(ctx, d, m)
}

func resourceNetworkDeviceListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname := resourceMap["hostname"]
	vManagementIPAddress := resourceMap["management_ip_address"]
	vMacAddress := resourceMap["mac_address"]
	vLocationName := resourceMap["location_name"]
	vSerialNumber := resourceMap["serial_number"]
	vLocation := resourceMap["location"]
	vFamily := resourceMap["family"]
	vType := resourceMap["type"]
	vSeries := resourceMap["series"]
	vCollectionStatus := resourceMap["collection_status"]
	vCollectionInterval := resourceMap["collection_interval"]
	vNotSyncedForMinutes := resourceMap["not_synced_for_minutes"]
	vErrorCode := resourceMap["error_code"]
	vErrorDescription := resourceMap["error_description"]
	vSoftwareVersion := resourceMap["software_version"]
	vSoftwareType := resourceMap["software_type"]
	vPlatformID := resourceMap["platform_id"]
	vRole := resourceMap["role"]
	vReachabilityStatus := resourceMap["reachability_status"]
	vUpTime := resourceMap["up_time"]
	vAssociatedWlcIP := resourceMap["associated_wlc_ip"]
	vLicensename := resourceMap["license_name"]
	vLicensetype := resourceMap["license_type"]
	vLicensestatus := resourceMap["license_status"]
	vModulename := resourceMap["module_name"]
	vModuleequpimenttype := resourceMap["module_equpimenttype"]
	vModuleservicestate := resourceMap["module_servicestate"]
	vModulevendorequipmenttype := resourceMap["module_vendorequipmenttype"]
	vModulepartnumber := resourceMap["module_partnumber"]
	vModuleoperationstatecode := resourceMap["module_operationstatecode"]
	vID := resourceMap["id"]
	vDeviceSupportLevel := resourceMap["device_support_level"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceList")
		queryParams1 := dnacentersdkgo.GetDeviceListQueryParams{}

		if okHostname {
			queryParams1.Hostname = interfaceToSliceString(vHostname)
		}
		if okManagementIPAddress {
			queryParams1.ManagementIPAddress = interfaceToSliceString(vManagementIPAddress)
		}
		if okMacAddress {
			queryParams1.MacAddress = interfaceToSliceString(vMacAddress)
		}
		if okLocationName {
			queryParams1.LocationName = interfaceToSliceString(vLocationName)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = interfaceToSliceString(vSerialNumber)
		}
		if okLocation {
			queryParams1.Location = interfaceToSliceString(vLocation)
		}
		if okFamily {
			queryParams1.Family = interfaceToSliceString(vFamily)
		}
		if okType {
			queryParams1.Type = interfaceToSliceString(vType)
		}
		if okSeries {
			queryParams1.Series = interfaceToSliceString(vSeries)
		}
		if okCollectionStatus {
			queryParams1.CollectionStatus = interfaceToSliceString(vCollectionStatus)
		}
		if okCollectionInterval {
			queryParams1.CollectionInterval = interfaceToSliceString(vCollectionInterval)
		}
		if okNotSyncedForMinutes {
			queryParams1.NotSyncedForMinutes = interfaceToSliceString(vNotSyncedForMinutes)
		}
		if okErrorCode {
			queryParams1.ErrorCode = interfaceToSliceString(vErrorCode)
		}
		if okErrorDescription {
			queryParams1.ErrorDescription = interfaceToSliceString(vErrorDescription)
		}
		if okSoftwareVersion {
			queryParams1.SoftwareVersion = interfaceToSliceString(vSoftwareVersion)
		}
		if okSoftwareType {
			queryParams1.SoftwareType = interfaceToSliceString(vSoftwareType)
		}
		if okPlatformID {
			queryParams1.PlatformID = interfaceToSliceString(vPlatformID)
		}
		if okRole {
			queryParams1.Role = interfaceToSliceString(vRole)
		}
		if okReachabilityStatus {
			queryParams1.ReachabilityStatus = interfaceToSliceString(vReachabilityStatus)
		}
		if okUpTime {
			queryParams1.UpTime = interfaceToSliceString(vUpTime)
		}
		if okAssociatedWlcIP {
			queryParams1.AssociatedWlcIP = interfaceToSliceString(vAssociatedWlcIP)
		}
		if okLicensename {
			queryParams1.Licensename = interfaceToSliceString(vLicensename)
		}
		if okLicensetype {
			queryParams1.Licensetype = interfaceToSliceString(vLicensetype)
		}
		if okLicensestatus {
			queryParams1.Licensestatus = interfaceToSliceString(vLicensestatus)
		}
		if okModulename {
			queryParams1.Modulename = interfaceToSliceString(vModulename)
		}
		if okModuleequpimenttype {
			queryParams1.Moduleequpimenttype = interfaceToSliceString(vModuleequpimenttype)
		}
		if okModuleservicestate {
			queryParams1.Moduleservicestate = interfaceToSliceString(vModuleservicestate)
		}
		if okModulevendorequipmenttype {
			queryParams1.Modulevendorequipmenttype = interfaceToSliceString(vModulevendorequipmenttype)
		}
		if okModulepartnumber {
			queryParams1.Modulepartnumber = interfaceToSliceString(vModulepartnumber)
		}
		if okModuleoperationstatecode {
			queryParams1.Moduleoperationstatecode = interfaceToSliceString(vModuleoperationstatecode)
		}
		if okID {
			queryParams1.ID = vID
		}
		if okDeviceSupportLevel {
			queryParams1.DeviceSupportLevel = vDeviceSupportLevel
		}
		if okOffset {
			queryParams1.Offset = *stringToFloat64Ptr(vOffset)
		}
		if okLimit {
			queryParams1.Limit = *stringToFloat64Ptr(vLimit)
		}

		response1, restyResp1, err := client.Devices.GetDeviceList(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceList", err,
				"Failure at GetDeviceList, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsDevicesGetDeviceList(m, response1, nil)
		item1, err := searchDevicesGetDeviceList(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetDeviceList response", err,
				"Failure when searching item from GetDeviceList, unexpected response", ""))
			return diags
		}
		// Review flatten function used
		vItem1 := flattenDevicesGetDeviceListByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceList search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceNetworkDeviceListUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vHostname := resourceMap["hostname"]
	vManagementIPAddress := resourceMap["management_ip_address"]
	vMacAddress := resourceMap["mac_address"]
	vLocationName := resourceMap["location_name"]
	vSerialNumber := resourceMap["serial_number"]
	vLocation := resourceMap["location"]
	vFamily := resourceMap["family"]
	vType := resourceMap["type"]
	vSeries := resourceMap["series"]
	vCollectionStatus := resourceMap["collection_status"]
	vCollectionInterval := resourceMap["collection_interval"]
	vNotSyncedForMinutes := resourceMap["not_synced_for_minutes"]
	vErrorCode := resourceMap["error_code"]
	vErrorDescription := resourceMap["error_description"]
	vSoftwareVersion := resourceMap["software_version"]
	vSoftwareType := resourceMap["software_type"]
	vPlatformID := resourceMap["platform_id"]
	vRole := resourceMap["role"]
	vReachabilityStatus := resourceMap["reachability_status"]
	vUpTime := resourceMap["up_time"]
	vAssociatedWlcIP := resourceMap["associated_wlc_ip"]
	vLicensename := resourceMap["license_name"]
	vLicensetype := resourceMap["license_type"]
	vLicensestatus := resourceMap["license_status"]
	vModulename := resourceMap["module_name"]
	vModuleequpimenttype := resourceMap["module_equpimenttype"]
	vModuleservicestate := resourceMap["module_servicestate"]
	vModulevendorequipmenttype := resourceMap["module_vendorequipmenttype"]
	vModulepartnumber := resourceMap["module_partnumber"]
	vModuleoperationstatecode := resourceMap["module_operationstatecode"]
	vID := resourceMap["id"]
	vDeviceSupportLevel := resourceMap["device_support_level"]
	vOffset := resourceMap["offset"]
	vLimit := resourceMap["limit"]

	queryParams1 := dnacentersdkgo.GetDeviceListQueryParams
	queryParams1.Hostname = interfaceToSliceString(vHostname)
	queryParams1.ManagementIPAddress = interfaceToSliceString(vManagementIPAddress)
	queryParams1.MacAddress = interfaceToSliceString(vMacAddress)
	queryParams1.LocationName = interfaceToSliceString(vLocationName)
	queryParams1.SerialNumber = interfaceToSliceString(vSerialNumber)
	queryParams1.Location = interfaceToSliceString(vLocation)
	queryParams1.Family = interfaceToSliceString(vFamily)
	queryParams1.Type = interfaceToSliceString(vType)
	queryParams1.Series = interfaceToSliceString(vSeries)
	queryParams1.CollectionStatus = interfaceToSliceString(vCollectionStatus)
	queryParams1.CollectionInterval = interfaceToSliceString(vCollectionInterval)
	queryParams1.NotSyncedForMinutes = interfaceToSliceString(vNotSyncedForMinutes)
	queryParams1.ErrorCode = interfaceToSliceString(vErrorCode)
	queryParams1.ErrorDescription = interfaceToSliceString(vErrorDescription)
	queryParams1.SoftwareVersion = interfaceToSliceString(vSoftwareVersion)
	queryParams1.SoftwareType = interfaceToSliceString(vSoftwareType)
	queryParams1.PlatformID = interfaceToSliceString(vPlatformID)
	queryParams1.Role = interfaceToSliceString(vRole)
	queryParams1.ReachabilityStatus = interfaceToSliceString(vReachabilityStatus)
	queryParams1.UpTime = interfaceToSliceString(vUpTime)
	queryParams1.AssociatedWlcIP = interfaceToSliceString(vAssociatedWlcIP)
	queryParams1.Licensename = interfaceToSliceString(vLicensename)
	queryParams1.Licensetype = interfaceToSliceString(vLicensetype)
	queryParams1.Licensestatus = interfaceToSliceString(vLicensestatus)
	queryParams1.Modulename = interfaceToSliceString(vModulename)
	queryParams1.Moduleequpimenttype = interfaceToSliceString(vModuleequpimenttype)
	queryParams1.Moduleservicestate = interfaceToSliceString(vModuleservicestate)
	queryParams1.Modulevendorequipmenttype = interfaceToSliceString(vModulevendorequipmenttype)
	queryParams1.Modulepartnumber = interfaceToSliceString(vModulepartnumber)
	queryParams1.Moduleoperationstatecode = interfaceToSliceString(vModuleoperationstatecode)
	queryParams1.ID = vID
	queryParams1.DeviceSupportLevel = vDeviceSupportLevel
	queryParams1.Offset = *stringToFloat64Ptr(vOffset)
	queryParams1.Limit = *stringToFloat64Ptr(vLimit)
	item, err := searchDevicesGetDeviceList(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDeviceList", err,
			"Failure at GetDeviceList, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestNetworkDeviceListSyncDevices2(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Devices.SyncDevices2(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing SyncDevices2", err, restyResp1.String(),
					"Failure at SyncDevices2, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing SyncDevices2", err,
				"Failure at SyncDevices2, unexpected response", ""))
			return diags
		}
	}

	return resourceNetworkDeviceListRead(ctx, d, m)
}

func resourceNetworkDeviceListDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete NetworkDeviceList on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestNetworkDeviceListAddDevice2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesAddDevice2 {
	request := dnacentersdkgo.RequestDevicesAddDevice2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_transport")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_transport")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_transport")))) {
		request.CliTransport = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".compute_device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".compute_device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".compute_device")))) {
		request.ComputeDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password")))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".extended_discovery_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".extended_discovery_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".extended_discovery_info")))) {
		request.ExtendedDiscoveryInfo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_password")))) {
		request.HTTPPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_port")))) {
		request.HTTPPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_secure")))) {
		request.HTTPSecure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_user_name")))) {
		request.HTTPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".meraki_org_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".meraki_org_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".meraki_org_id")))) {
		request.MerakiOrgID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPROCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRWCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_retry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_retry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_retry")))) {
		request.SNMPRetry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_timeout")))) {
		request.SNMPTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_version")))) {
		request.SNMPVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_mgmt_ipaddress_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_mgmt_ipaddress_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_mgmt_ipaddress_list")))) {
		request.UpdateMgmtIPaddressList = expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressListArray(ctx, key+".update_mgmt_ipaddress_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList {
	request := []dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListAddDevice2UpdateMgmtIPaddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList {
	request := dnacentersdkgo.RequestDevicesAddDevice2UpdateMgmtIPaddressList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exist_mgmt_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exist_mgmt_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exist_mgmt_ip_address")))) {
		request.ExistMgmtIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_mgmt_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_mgmt_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_mgmt_ip_address")))) {
		request.NewMgmtIPAddress = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListSyncDevices2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesSyncDevices2 {
	request := dnacentersdkgo.RequestDevicesSyncDevices2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_transport")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_transport")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_transport")))) {
		request.CliTransport = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".compute_device")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".compute_device")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".compute_device")))) {
		request.ComputeDevice = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password")))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".extended_discovery_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".extended_discovery_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".extended_discovery_info")))) {
		request.ExtendedDiscoveryInfo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_password")))) {
		request.HTTPPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_port")))) {
		request.HTTPPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_secure")))) {
		request.HTTPSecure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_user_name")))) {
		request.HTTPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".meraki_org_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".meraki_org_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".meraki_org_id")))) {
		request.MerakiOrgID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".netconf_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".netconf_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".netconf_port")))) {
		request.NetconfPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_passphrase")))) {
		request.SNMPAuthPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_auth_protocol")))) {
		request.SNMPAuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_passphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_passphrase")))) {
		request.SNMPPrivPassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_priv_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_priv_protocol")))) {
		request.SNMPPrivProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_ro_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_ro_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_ro_community")))) {
		request.SNMPROCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_rw_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_rw_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_rw_community")))) {
		request.SNMPRWCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_retry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_retry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_retry")))) {
		request.SNMPRetry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_timeout")))) {
		request.SNMPTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_user_name")))) {
		request.SNMPUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_version")))) {
		request.SNMPVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_mgmt_ipaddress_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_mgmt_ipaddress_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_mgmt_ipaddress_list")))) {
		request.UpdateMgmtIPaddressList = expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressListArray(ctx, key+".update_mgmt_ipaddress_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressListArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList {
	request := []dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkDeviceListSyncDevices2UpdateMgmtIPaddressList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList {
	request := dnacentersdkgo.RequestDevicesSyncDevices2UpdateMgmtIPaddressList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exist_mgmt_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exist_mgmt_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exist_mgmt_ip_address")))) {
		request.ExistMgmtIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_mgmt_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_mgmt_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_mgmt_ip_address")))) {
		request.NewMgmtIPAddress = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchDevicesGetDeviceList(m interface{}, queryParams dnacentersdkgo.GetDeviceListQueryParams) (*dnacentersdkgo.ResponseItemDevicesGetDeviceList, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemDevicesGetDeviceList
	var ite *dnacentersdkgo.ResponseDevicesGetDeviceList
	ite, _, err = client.Devices.GetDeviceList(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemDevicesGetDeviceList
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
