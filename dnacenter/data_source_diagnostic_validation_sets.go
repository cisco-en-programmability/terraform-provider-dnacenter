package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiagnosticValidationSets() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Health and Performance.

- Retrieves all the validation sets and optionally the contained validations

- Retrieves validation details for the given validation set id
`,

		ReadContext: dataSourceDiagnosticValidationSetsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Validation set id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. When the query parameter *view=DETAIL* is passed, all validation sets and associated validations will be returned. When the query parameter *view=DEFAULT* is passed, only validation sets metadata will be returned.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description of the validation set
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Validation set id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the validation set
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"validation_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description of the validation group
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Validation group id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the validation group
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"validations": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Validation id
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Name of the validation
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Version of validation set
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

						"description": &schema.Schema{
							Description: `Description of the validation set
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Validation set id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the validation set
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"validation_groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description of the validation group
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Validation group id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the validation group
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"validations": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Validation id
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Name of the validation
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Version of the validation set
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

func dataSourceDiagnosticValidationSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vView, okView := d.GetOk("view")
	vID, okID := d.GetOk("id")

	method1 := []bool{okView}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesAllTheValidationSets")
		queryParams1 := dnacentersdkgo.RetrievesAllTheValidationSetsQueryParams{}

		if okView {
			queryParams1.View = vView.(string)
		}

		response1, restyResp1, err := client.HealthAndPerformance.RetrievesAllTheValidationSets(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesAllTheValidationSets", err,
				"Failure at RetrievesAllTheValidationSets, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenHealthAndPerformanceRetrievesAllTheValidationSetsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesAllTheValidationSets response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: RetrievesValidationDetailsForAValidationSet")
		vvID := vID.(string)

		response2, restyResp2, err := client.HealthAndPerformance.RetrievesValidationDetailsForAValidationSet(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesValidationDetailsForAValidationSet", err,
				"Failure at RetrievesValidationDetailsForAValidationSet, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenHealthAndPerformanceRetrievesValidationDetailsForAValidationSetItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesValidationDetailsForAValidationSet response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenHealthAndPerformanceRetrievesAllTheValidationSetsItems(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["version"] = item.Version
		respItem["validation_groups"] = flattenHealthAndPerformanceRetrievesAllTheValidationSetsItemsValidationGroups(item.ValidationGroups)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenHealthAndPerformanceRetrievesAllTheValidationSetsItemsValidationGroups(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponseValidationGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["description"] = item.Description
		respItem["validations"] = flattenHealthAndPerformanceRetrievesAllTheValidationSetsItemsValidationGroupsValidations(item.Validations)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenHealthAndPerformanceRetrievesAllTheValidationSetsItemsValidationGroupsValidations(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponseValidationGroupsValidations) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenHealthAndPerformanceRetrievesValidationDetailsForAValidationSetItem(item *dnacentersdkgo.ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["version"] = item.Version
	respItem["validation_groups"] = flattenHealthAndPerformanceRetrievesValidationDetailsForAValidationSetItemValidationGroups(item.ValidationGroups)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenHealthAndPerformanceRetrievesValidationDetailsForAValidationSetItemValidationGroups(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponseValidationGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["description"] = item.Description
		respItem["validations"] = flattenHealthAndPerformanceRetrievesValidationDetailsForAValidationSetItemValidationGroupsValidations(item.Validations)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenHealthAndPerformanceRetrievesValidationDetailsForAValidationSetItemValidationGroupsValidations(items *[]dnacentersdkgo.ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponseValidationGroupsValidations) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
