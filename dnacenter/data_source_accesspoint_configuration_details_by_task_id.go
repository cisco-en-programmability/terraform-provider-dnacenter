package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccesspointConfigurationDetailsByTaskID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Users can query the access point configuration result using this intent API
`,

		ReadContext: dataSourceAccesspointConfigurationDetailsByTaskIDRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `task_id path parameter. task id information of ap config
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"creation_order_index": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"is_being_changed": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ordered_list_oeassoc_name": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"ordered_list_oeindex": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"ap_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"auth_entity_class": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"auth_entity_id": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"change_log_list": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"deploy_pending": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_created_on": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_id": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"instance_origin": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_updated_on": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_version": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"internal_key": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"long_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"url": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"lazy_loaded_entities": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"location_heirarchy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"status_details": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAccesspointConfigurationDetailsByTaskIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID := d.Get("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointConfigurationTaskResult")
		vvTaskID := vTaskID.(string)

		response1, restyResp1, err := client.Wireless.GetAccessPointConfigurationTaskResult(vvTaskID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointConfigurationTaskResult", err,
				"Failure at GetAccessPointConfigurationTaskResult, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetAccessPointConfigurationTaskResultItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointConfigurationTaskResult response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAccessPointConfigurationTaskResultItems(items *dnacentersdkgo.ResponseWirelessGetAccessPointConfigurationTaskResult) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_uuid"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceUUID(item.InstanceUUID)
		respItem["instance_id"] = item.InstanceID
		respItem["auth_entity_id"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsAuthEntityID(item.AuthEntityID)
		respItem["display_name"] = item.DisplayName
		respItem["auth_entity_class"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsAuthEntityClass(item.AuthEntityClass)
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["ordered_list_oeindex"] = item.OrderedListOEIndex
		respItem["ordered_list_oeassoc_name"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsOrderedListOEAssocName(item.OrderedListOEAssocName)
		respItem["creation_order_index"] = item.CreationOrderIndex
		respItem["is_being_changed"] = boolPtrToString(item.IsBeingChanged)
		respItem["deploy_pending"] = item.DeployPending
		respItem["instance_created_on"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceCreatedOn(item.InstanceCreatedOn)
		respItem["instance_updated_on"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceUpdatedOn(item.InstanceUpdatedOn)
		respItem["change_log_list"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsChangeLogList(item.ChangeLogList)
		respItem["instance_origin"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceOrigin(item.InstanceOrigin)
		respItem["lazy_loaded_entities"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsLazyLoadedEntities(item.LazyLoadedEntities)
		respItem["instance_version"] = item.InstanceVersion
		respItem["ap_name"] = item.ApName
		respItem["controller_name"] = item.ControllerName
		respItem["location_heirarchy"] = item.LocationHeirarchy
		respItem["mac_address"] = item.MacAddress
		respItem["status"] = item.Status
		respItem["status_details"] = item.StatusDetails
		respItem["internal_key"] = flattenWirelessGetAccessPointConfigurationTaskResultItemsInternalKey(item.InternalKey)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceUUID(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUUID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsAuthEntityID(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsAuthEntityClass(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityClass) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsOrderedListOEAssocName(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultOrderedListOEAssocName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceCreatedOn(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceCreatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceUpdatedOn(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUpdatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsChangeLogList(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultChangeLogList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsInstanceOrigin(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceOrigin) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsLazyLoadedEntities(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultLazyLoadedEntities) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationTaskResultItemsInternalKey(item *dnacentersdkgo.ResponseItemWirelessGetAccessPointConfigurationTaskResultInternalKey) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["id"] = item.ID
	respItem["long_type"] = item.LongType
	respItem["url"] = item.URL

	return []map[string]interface{}{
		respItem,
	}

}
