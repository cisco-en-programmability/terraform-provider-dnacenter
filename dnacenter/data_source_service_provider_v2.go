package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServiceProviderV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- API to get Service Provider details (QoS).
`,

		ReadContext: dataSourceServiceProviderV2Read,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"group_uuid": &schema.Schema{
							Description: `Group Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"inherited_group_name": &schema.Schema{
							Description: `Inherited Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"inherited_group_uuid": &schema.Schema{
							Description: `Inherited Group Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_type": &schema.Schema{
							Description: `Instance Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"key": &schema.Schema{
							Description: `Key`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"namespace": &schema.Schema{
							Description: `Namespace`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"sla_profile_name": &schema.Schema{
										Description: `Sla Profile Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"sp_profile_name": &schema.Schema{
										Description: `Sp Profile Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"wan_provider": &schema.Schema{
										Description: `Wan Provider`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceServiceProviderV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetServiceProviderDetailsV2")

		response1, restyResp1, err := client.NetworkSettings.GetServiceProviderDetailsV2()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetServiceProviderDetailsV2", err,
				"Failure at GetServiceProviderDetailsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsGetServiceProviderDetailsV2Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetServiceProviderDetailsV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetServiceProviderDetailsV2Items(items *[]dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsV2Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_type"] = item.InstanceType
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["namespace"] = item.Namespace
		respItem["type"] = item.Type
		respItem["key"] = item.Key
		respItem["version"] = item.Version
		respItem["value"] = flattenNetworkSettingsGetServiceProviderDetailsV2ItemsValue(item.Value)
		respItem["group_uuid"] = item.GroupUUID
		respItem["inherited_group_uuid"] = item.InheritedGroupUUID
		respItem["inherited_group_name"] = item.InheritedGroupName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetServiceProviderDetailsV2ItemsValue(items *[]dnacentersdkgo.ResponseNetworkSettingsGetServiceProviderDetailsV2ResponseValue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["wan_provider"] = item.WanProvider
		respItem["sp_profile_name"] = item.SpProfileName
		respItem["sla_profile_name"] = item.SLAProfileName
		respItems = append(respItems, respItem)
	}
	return respItems
}
