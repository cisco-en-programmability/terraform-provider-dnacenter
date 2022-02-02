package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLanAutomationStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on LAN Automation.

- Invoke this API to get the LAN Automation session status

- Invoke this API to get the LAN Automation session status based on the given Lan Automation session Id
`,

		ReadContext: dataSourceLanAutomationStatusRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. LAN Automation Session Identifier
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of LAN Automations session status to be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Offset/starting row of the LAN Automation session from which status needs to be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"action": &schema.Schema{
							Description: `State (START/STOP) of the LAN Automation Session
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"creation_time": &schema.Schema{
							Description: `LAN Automation session creation time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovered_device_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_address_in_use_list": &schema.Schema{
										Description: `The list of IP address used by device
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"name": &schema.Schema{
										Description: `Name of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"serial_number": &schema.Schema{
										Description: `Serial number of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `state of the device like added to inventory/ deleted from inventory
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"discovered_device_site_name_hierarchy": &schema.Schema{
							Description: `Discovered sevice site name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `System generated identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_pool_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_pool_name": &schema.Schema{
										Description: `Name of the IP pool
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_pool_role": &schema.Schema{
										Description: `Role of the IP pool
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"multicast_enabled": &schema.Schema{
							Description: `Shows whether underlay multicast is enabled or not
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"peer_device_managment_ipaddress": &schema.Schema{
							Description: `Peer seed device management IP address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"primary_device_interface_names": &schema.Schema{
							Description: `The List of interfaces on primary seed device via which the discovered devices are connected
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"primary_device_managment_ipaddress": &schema.Schema{
							Description: `Primary seed device management IP address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of LAN Automation session and provides the number of discovered devices
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

						"action": &schema.Schema{
							Description: `State (START/STOP) of the LAN Automation Session
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"creation_time": &schema.Schema{
							Description: `LAN Automation session creation time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovered_device_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_address_in_use_list": &schema.Schema{
										Description: `The list of IP address used by device
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"name": &schema.Schema{
										Description: `Name of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"serial_number": &schema.Schema{
										Description: `Serial number of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `state of the device like added to inventory/ deleted from inventory
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"discovered_device_site_name_hierarchy": &schema.Schema{
							Description: `Discovered sevice site name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `System generated identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_pool_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_pool_name": &schema.Schema{
										Description: `Name of the IP pool
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_pool_role": &schema.Schema{
										Description: `Role of the IP pool
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"multicast_enabled": &schema.Schema{
							Description: `Shows whether underlay multicast is enabled or not
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"peer_device_managment_ipaddress": &schema.Schema{
							Description: `Peer seed device management IP address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"primary_device_interface_names": &schema.Schema{
							Description: `The List of interfaces on primary seed device via which the discovered devices are connected
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"primary_device_managment_ipaddress": &schema.Schema{
							Description: `Primary seed device management IP address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of LAN Automation session and provides the number of discovered devices
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

func dataSourceLanAutomationStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vID, okID := d.GetOk("id")

	method1 := []bool{okOffset, okLimit}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LanAutomationStatus")
		queryParams1 := dnacentersdkgo.LanAutomationStatusQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(string)
		}

		response1, restyResp1, err := client.LanAutomation.LanAutomationStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LanAutomationStatus", err,
				"Failure at LanAutomationStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLanAutomationLanAutomationStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: LanAutomationStatusByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.LanAutomation.LanAutomationStatusByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LanAutomationStatusByID", err,
				"Failure at LanAutomationStatusByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenLanAutomationLanAutomationStatusByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationStatusByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLanAutomationLanAutomationStatusItems(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["discovered_device_site_name_hierarchy"] = item.DiscoveredDeviceSiteNameHierarchy
		respItem["primary_device_managment_ipaddress"] = item.PrimaryDeviceManagmentIPAddress
		respItem["ip_pool_list"] = flattenLanAutomationLanAutomationStatusItemsIPPoolList(item.IPPoolList)
		respItem["primary_device_interface_names"] = item.PrimaryDeviceInterfaceNames
		respItem["status"] = item.Status
		respItem["action"] = item.Action
		respItem["creation_time"] = item.CreationTime
		respItem["multicast_enabled"] = boolPtrToString(item.MulticastEnabled)
		respItem["peer_device_managment_ipaddress"] = item.PeerDeviceManagmentIPAddress
		respItem["discovered_device_list"] = flattenLanAutomationLanAutomationStatusItemsDiscoveredDeviceList(item.DiscoveredDeviceList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusItemsIPPoolList(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationStatusResponseIPPoolList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["ip_pool_role"] = item.IPPoolRole
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusItemsDiscoveredDeviceList(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["serial_number"] = item.SerialNumber
		respItem["state"] = item.State
		respItem["ip_address_in_use_list"] = item.IPAddressInUseList
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusByIDItem(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationStatusByIDResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["discovered_device_site_name_hierarchy"] = item.DiscoveredDeviceSiteNameHierarchy
		respItem["primary_device_managment_ipaddress"] = item.PrimaryDeviceManagmentIPAddress
		respItem["ip_pool_list"] = flattenLanAutomationLanAutomationStatusByIDItemIPPoolList(item.IPPoolList)
		respItem["primary_device_interface_names"] = item.PrimaryDeviceInterfaceNames
		respItem["status"] = item.Status
		respItem["action"] = item.Action
		respItem["creation_time"] = item.CreationTime
		respItem["multicast_enabled"] = boolPtrToString(item.MulticastEnabled)
		respItem["peer_device_managment_ipaddress"] = item.PeerDeviceManagmentIPAddress
		respItem["discovered_device_list"] = flattenLanAutomationLanAutomationStatusByIDItemDiscoveredDeviceList(item.DiscoveredDeviceList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusByIDItemIPPoolList(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationStatusByIDResponseIPPoolList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["ip_pool_role"] = item.IPPoolRole
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusByIDItemDiscoveredDeviceList(items *[]dnacentersdkgo.ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["serial_number"] = item.SerialNumber
		respItem["state"] = item.State
		respItem["ip_address_in_use_list"] = item.IPAddressInUseList
		respItems = append(respItems, respItem)
	}
	return respItems
}
