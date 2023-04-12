package dnacenter

import (
	"context"

	"errors"

	"time"

	"fmt"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceConfigurationTemplateExportTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- Exports the templates for given templateIds.
`,

		CreateContext: resourceConfigurationTemplateExportTemplateCreate,
		ReadContext:   resourceConfigurationTemplateExportTemplateRead,
		DeleteContext: resourceConfigurationTemplateExportTemplateDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceConfigurationTemplateExportTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteria(ctx, "parameters.0", d)

	response1, restyResp1, err := client.ConfigurationTemplates.ExportsTheTemplatesForAGivenCriteria(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ExportTemplates", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ExportTemplates", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ExportTemplates", err1))
			return diags
		}
	}

	vItem1 := flattenConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ExportsTheTemplatesForAGivenCriteria response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceConfigurationTemplateExportTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceConfigurationTemplateExportTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteria(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria {
	request := dnacentersdkgo.RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria{}
	if v := expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteriaItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteriaItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria {
	request := []dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria{}
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
		i := expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteriaItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestConfigurationTemplateExportTemplateExportsTheTemplatesForAGivenCriteriaItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria {
	var request dnacentersdkgo.RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaItem(item *dnacentersdkgo.ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaResponse) []map[string]interface{} {
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
