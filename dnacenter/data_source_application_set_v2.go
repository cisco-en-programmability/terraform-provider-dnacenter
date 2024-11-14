package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplicationSetV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get application set/s by offset/limit or by name
`,

		ReadContext: dataSourceApplicationSetV2Read,
		Schema: map[string]*schema.Schema{
			"attributes": &schema.Schema{
				Description: `attributes query parameter. Attributes to retrieve, valid value applicationSet
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
				Description: `name query parameter. Application set name
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

						"default_business_relevance": &schema.Schema{
							Description: `Default business relevance
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

						"id": &schema.Schema{
							Description: `Id of Application Set
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
										Description: `Type of identify source. NBAR: build in Application Set, APIC_EM: custom Application Set
`,
										Type:     schema.TypeString,
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
							Description: `Application Set name
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

						"scalable_group_external_handle": &schema.Schema{
							Description: `Scalable group external handle, should be equal to Application Set name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"scalable_group_type": &schema.Schema{
							Description: `Scalable group type, valid value APPLICATION_GROUP
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

func dataSourceApplicationSetV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vAttributes := d.Get("attributes")
	vName, okName := d.GetOk("name")
	vOffset := d.Get("offset")
	vLimit := d.Get("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationSetsV2")
		queryParams1 := dnacentersdkgo.GetApplicationSetsV2QueryParams{}

		queryParams1.Attributes = vAttributes.(string)

		if okName {
			queryParams1.Name = vName.(string)
		}
		queryParams1.Offset = vOffset.(float64)

		queryParams1.Limit = vLimit.(float64)

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationSetsV2(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApplicationSetsV2", err,
				"Failure at GetApplicationSetsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyGetApplicationSetsV2Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationSetsV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationSetsV2Items(items *[]dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsV2Response) []map[string]interface{} {
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
		respItem["default_business_relevance"] = item.DefaultBusinessRelevance
		respItem["identity_source"] = flattenApplicationPolicyGetApplicationSetsV2ItemsIDentitySource(item.IDentitySource)
		respItem["name"] = item.Name
		respItem["namespace"] = item.Namespace
		respItem["scalable_group_external_handle"] = item.ScalableGroupExternalHandle
		respItem["scalable_group_type"] = item.ScalableGroupType
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationSetsV2ItemsIDentitySource(item *dnacentersdkgo.ResponseApplicationPolicyGetApplicationSetsV2ResponseIDentitySource) []map[string]interface{} {
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
