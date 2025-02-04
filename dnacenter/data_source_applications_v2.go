package dnacenter

import (
	"context"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationsV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get application/s by offset/limit or by name
`,

		ReadContext: dataSourceApplicationsV2Read,
		Schema: map[string]*schema.Schema{
			"attributes": &schema.Schema{
				Description: `attributes query parameter. Attributes to retrieve, valid value application
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The limit which is the maximum number of items to include in a single page of results, max value 500
`,
				Type:     schema.TypeFloat,
				Required: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. The application name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The starting point or index from where the paginated results should begin.
`,
				Type:     schema.TypeFloat,
				Required: true,
			},

			"items": &schema.Schema{
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

						"id": &schema.Schema{
							Description: `Id of Application
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"identity_source": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `Type of identify source. NBAR: build in Application, APIC_EM: custom Application
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"indicative_network_identity": &schema.Schema{
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

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"lower_port": &schema.Schema{
										Description: `Lower port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"ports": &schema.Schema{
										Description: `Ports
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"upper_port": &schema.Schema{
										Description: `Upper port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"instance_id": &schema.Schema{
							Description: `Instance id
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

						"name": &schema.Schema{
							Description: `Application name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"namespace": &schema.Schema{
							Description: `Namespace, valid value scalablegroup:application
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_applications": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_protocol": &schema.Schema{
										Description: `App protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"application_sub_type": &schema.Schema{
										Description: `Application sub type, LEARNED: discovered application, NONE: nbar and custom application
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"application_type": &schema.Schema{
										Description: `Application type, DEFAULT: nbar application, DEFAULT_MODIFIED: nbar modified application, CUSTOM: custom application
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"category_id": &schema.Schema{
										Description: `Category id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dscp": &schema.Schema{
										Description: `Dscp
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"engine_id": &schema.Schema{
										Description: `Engine id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"help_string": &schema.Schema{
										Description: `Help string
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"long_description": &schema.Schema{
										Description: `Long description
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Application name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"popularity": &schema.Schema{
										Description: `Popularity
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"rank": &schema.Schema{
										Description: `Rank, any value between 1 to 65535
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"selector_id": &schema.Schema{
										Description: `Selector id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"server_name": &schema.Schema{
										Description: `Server name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"traffic_class": &schema.Schema{
										Description: `Traffic class
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"url": &schema.Schema{
										Description: `Url
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"network_identity": &schema.Schema{
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

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ipv4_subnet": &schema.Schema{
										Description: `Ipv4 subnet
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ipv6_subnet": &schema.Schema{
										Description: `Ipv6 subnet
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"lower_port": &schema.Schema{
										Description: `Lower port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"ports": &schema.Schema{
										Description: `Ports
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"upper_port": &schema.Schema{
										Description: `Upper port
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"parent_scalable_group": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id_ref": &schema.Schema{
										Description: `Id reference to parent application set
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"qualifier": &schema.Schema{
							Description: `Qualifier, valid value application
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"scalable_group_external_handle": &schema.Schema{
							Description: `Scalable group external handle, should be equal to Application name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"scalable_group_type": &schema.Schema{
							Description: `Scalable group type, valid value APPLICATION
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type, valid value scalablegroup
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

func dataSourceApplicationsV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vAttributes := d.Get("attributes")
	vName, okName := d.GetOk("name")
	vOffset := d.Get("offset")
	vLimit := d.Get("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationsV2")
		queryParams1 := dnacentersdkgo.GetApplicationsV2QueryParams{}

		queryParams1.Attributes = vAttributes.(string)

		if okName {
			queryParams1.Name = vName.(string)
		}
		queryParams1.Offset = vOffset.(float64)

		queryParams1.Limit = vLimit.(float64)

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationsV2(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApplicationsV2", err,
				"Failure at GetApplicationsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyGetApplicationsV2Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationsV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationsV2Items(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_version"] = item.InstanceVersion
		respItem["identity_source"] = flattenApplicationPolicyGetApplicationsV2ItemsIDentitySource(item.IDentitySource)
		respItem["indicative_network_identity"] = flattenApplicationPolicyGetApplicationsV2ItemsIndicativeNetworkIDentity(item.IndicativeNetworkIDentity)
		respItem["name"] = item.Name
		respItem["namespace"] = item.Namespace
		respItem["network_applications"] = flattenApplicationPolicyGetApplicationsV2ItemsNetworkApplications(item.NetworkApplications)
		respItem["network_identity"] = flattenApplicationPolicyGetApplicationsV2ItemsNetworkIDentity(item.NetworkIDentity)
		respItem["parent_scalable_group"] = flattenApplicationPolicyGetApplicationsV2ItemsParentScalableGroup(item.ParentScalableGroup)
		respItem["qualifier"] = item.Qualifier
		respItem["scalable_group_external_handle"] = item.ScalableGroupExternalHandle
		respItem["scalable_group_type"] = item.ScalableGroupType
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsV2ItemsIDentitySource(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2ResponseIDentitySource) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationsV2ItemsIndicativeNetworkIDentity(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2ResponseIndicativeNetworkIDentity) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["display_name"] = item.DisplayName
		respItem["lower_port"] = item.LowerPort
		respItem["ports"] = item.Ports
		respItem["protocol"] = item.Protocol
		respItem["upper_port"] = item.UpperPort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsV2ItemsNetworkApplications(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2ResponseNetworkApplications) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["app_protocol"] = item.AppProtocol
		respItem["application_sub_type"] = item.ApplicationSubType
		respItem["application_type"] = item.ApplicationType
		respItem["category_id"] = item.CategoryID
		respItem["display_name"] = item.DisplayName
		respItem["dscp"] = item.Dscp
		respItem["engine_id"] = item.EngineID
		respItem["help_string"] = item.HelpString
		respItem["long_description"] = item.LongDescription
		respItem["name"] = item.Name
		respItem["popularity"] = item.Popularity
		respItem["rank"] = item.Rank
		respItem["selector_id"] = item.SelectorID
		respItem["server_name"] = item.ServerName
		respItem["url"] = item.URL
		respItem["traffic_class"] = item.TrafficClass
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsV2ItemsNetworkIDentity(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentity) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["display_name"] = item.DisplayName
		respItem["ipv4_subnet"] = item.IPv4Subnet
		respItem["ipv6_subnet"] = flattenApplicationPolicyGetApplicationsV2ItemsNetworkIDentityIPv6Subnet(item.IPv6Subnet)
		respItem["lower_port"] = item.LowerPort
		respItem["ports"] = item.Ports
		respItem["protocol"] = item.Protocol
		respItem["upper_port"] = item.UpperPort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationsV2ItemsNetworkIDentityIPv6Subnet(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentityIPv6Subnet) []interface{} {
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

func flattenApplicationPolicyGetApplicationsV2ItemsParentScalableGroup(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationsV2ResponseParentScalableGroup) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["id_ref"] = item.IDRef

	return []map[string]interface{}{
		respItem,
	}

}
