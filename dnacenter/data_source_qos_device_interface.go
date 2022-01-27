package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceQosDeviceInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get all or by network device id, existing qos device interface infos
`,

		ReadContext: dataSourceQosDeviceInterfaceRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. network device id
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cfs_change_info": &schema.Schema{
							Description: `Cfs change info
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"create_time": &schema.Schema{
							Description: `Create time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"custom_provisions": &schema.Schema{
							Description: `Custom provisions
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"deployed": &schema.Schema{
							Description: `Deployed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"excluded_interfaces": &schema.Schema{
							Description: `excluded interfaces ids
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"id": &schema.Schema{
							Description: `Id of Qos device info
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_created_on": &schema.Schema{
							Description: `Instance created on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_id": &schema.Schema{
							Description: `Instance id
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_updated_on": &schema.Schema{
							Description: `Instance updated on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"instance_version": &schema.Schema{
							Description: `Instance version
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"is_excluded": &schema.Schema{
							Description: `Is excluded
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_seeded": &schema.Schema{
							Description: `Is seeded
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_stale": &schema.Schema{
							Description: `Is stale
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Last update time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"namespace": &schema.Schema{
							Description: `Namespace
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network device id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"provisioning_state": &schema.Schema{
							Description: `Provisioning state
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"qos_device_interface_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dmvpn_remote_sites_bw": &schema.Schema{
										Description: `Dmvpn remote sites bandwidth
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},

									"download_bw": &schema.Schema{
										Description: `Download bandwidth
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Qos device interface info
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"interface_id": &schema.Schema{
										Description: `Interface id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"interface_name": &schema.Schema{
										Description: `Interface name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"label": &schema.Schema{
										Description: `SP Profile name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"role": &schema.Schema{
										Description: `Interface role
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"upload_bw": &schema.Schema{
										Description: `Upload bandwidth
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"qualifier": &schema.Schema{
							Description: `Qualifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"resource_version": &schema.Schema{
							Description: `Resource version
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"target_id_list": &schema.Schema{
							Description: `Target id list
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"type": &schema.Schema{
							Description: `Type
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

func dataSourceQosDeviceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetQosDeviceInterfaceInfo")
		queryParams1 := dnacentersdkgo.GetQosDeviceInterfaceInfoQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetQosDeviceInterfaceInfo(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetQosDeviceInterfaceInfo", err,
				"Failure at GetQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyGetQosDeviceInterfaceInfoItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetQosDeviceInterfaceInfo response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetQosDeviceInterfaceInfoItems(items *[]dnacentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["create_time"] = item.CreateTime
		respItem["deployed"] = boolPtrToString(item.Deployed)
		respItem["is_seeded"] = boolPtrToString(item.IsSeeded)
		respItem["is_stale"] = boolPtrToString(item.IsStale)
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["namespace"] = item.Namespace
		respItem["provisioning_state"] = item.ProvisioningState
		respItem["qualifier"] = item.Qualifier
		respItem["resource_version"] = item.ResourceVersion
		respItem["target_id_list"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsTargetIDList(item.TargetIDList)
		respItem["type"] = item.Type
		respItem["cfs_change_info"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsCfsChangeInfo(item.CfsChangeInfo)
		respItem["custom_provisions"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsCustomProvisions(item.CustomProvisions)
		respItem["excluded_interfaces"] = item.ExcludedInterfaces
		respItem["is_excluded"] = boolPtrToString(item.IsExcluded)
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["qos_device_interface_info"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo(item.QosDeviceInterfaceInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetQosDeviceInterfaceInfoItem(item *dnacentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["create_time"] = item.CreateTime
	respItem["deployed"] = boolPtrToString(item.Deployed)
	respItem["is_seeded"] = boolPtrToString(item.IsSeeded)
	respItem["is_stale"] = boolPtrToString(item.IsStale)
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["name"] = item.Name
	respItem["namespace"] = item.Namespace
	respItem["provisioning_state"] = item.ProvisioningState
	respItem["qualifier"] = item.Qualifier
	respItem["resource_version"] = item.ResourceVersion
	respItem["target_id_list"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsTargetIDList(item.TargetIDList)
	respItem["type"] = item.Type
	respItem["cfs_change_info"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsCfsChangeInfo(item.CfsChangeInfo)
	respItem["custom_provisions"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsCustomProvisions(item.CustomProvisions)
	respItem["excluded_interfaces"] = item.ExcludedInterfaces
	respItem["is_excluded"] = boolPtrToString(item.IsExcluded)
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["qos_device_interface_info"] = flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo(item.QosDeviceInterfaceInfo)

	return []map[string]interface{}{
		respItem,
	}
}

func flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsTargetIDList(items *[]dnacentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseTargetIDList) []interface{} {
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

func flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsCfsChangeInfo(items *[]dnacentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseCfsChangeInfo) []interface{} {
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

func flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsCustomProvisions(items *[]dnacentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseCustomProvisions) []interface{} {
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

func flattenApplicationPolicyGetQosDeviceInterfaceInfoItemsQosDeviceInterfaceInfo(items *[]dnacentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseQosDeviceInterfaceInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["dmvpn_remote_sites_bw"] = item.DmvpnRemoteSitesBw
		respItem["download_bw"] = item.DownloadBW
		respItem["interface_id"] = item.InterfaceID
		respItem["interface_name"] = item.InterfaceName
		respItem["label"] = item.Label
		respItem["role"] = item.Role
		respItem["upload_bw"] = item.UploadBW
		respItems = append(respItems, respItem)
	}
	return respItems
}
