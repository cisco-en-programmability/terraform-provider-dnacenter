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

// dataSourceAction
func dataSourceConfigurationTemplateExportProject() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- Exports the projects for given projectNames.
`,

		ReadContext: dataSourceConfigurationTemplateExportProjectRead,
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
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
			"payload": &schema.Schema{
				Description: `Array of RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteria`,
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func dataSourceConfigurationTemplateExportProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ExportsTheProjectsForAGivenCriteria")
		request1 := expandRequestConfigurationTemplateExportProjectExportsTheProjectsForAGivenCriteria(ctx, "", d)

		response1, restyResp1, err := client.ConfigurationTemplates.ExportsTheProjectsForAGivenCriteria(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ExportsTheProjectsForAGivenCriteria", err,
				"Failure at ExportsTheProjectsForAGivenCriteria, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesExportsTheProjectsForAGivenCriteriaItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ExportsTheProjectsForAGivenCriteria response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestConfigurationTemplateExportProjectExportsTheProjectsForAGivenCriteria(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteria {
	request := dnacentersdkgo.RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteria{}
	if v := expandRequestConfigurationTemplateExportProjectExportsTheProjectsForAGivenCriteriaItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateExportProjectExportsTheProjectsForAGivenCriteriaItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteria {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteria{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestConfigurationTemplateExportProjectExportsTheProjectsForAGivenCriteriaItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestConfigurationTemplateExportProjectExportsTheProjectsForAGivenCriteriaItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteria {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteria
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenConfigurationTemplatesExportsTheProjectsForAGivenCriteriaItem(item *dnacentersdkgo.ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
