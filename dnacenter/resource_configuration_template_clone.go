package dnacenter

import (
	"context"
	"errors"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigurationTemplateClone() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.
- API to clone template
`,

		CreateContext: resourceConfigurationTemplateCloneCreate,
		ReadContext:   resourceConfigurationTemplateCloneRead,
		DeleteContext: resourceConfigurationTemplateCloneDelete,

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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Description: `name path parameter. Template name to clone template(Name should be different than existing template name within same project)
			`,
							Type:     schema.TypeString,
							ForceNew: true,
							Required: true,
						},
						"project_id": &schema.Schema{
							Description: `projectId path parameter. UUID of the project in which the template needs to be created
			`,
							Type:     schema.TypeString,
							ForceNew: true,
							Required: true,
						},
						"template_id": &schema.Schema{
							Description: `templateId path parameter. UUID of the template to clone it
			`,
							Type:     schema.TypeString,
							ForceNew: true,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceConfigurationTemplateCloneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vName := resourceItem["name"]
	vTemplateID := resourceItem["name"]
	vProjectID := resourceItem["name"]
	vvName := vName.(string)
	vvTemplateID := vTemplateID.(string)
	vvProjectID := vProjectID.(string)

	response1, restyResp1, err := client.ConfigurationTemplates.CreatesACloneOfTheGivenTemplate(vvName, vvTemplateID, vvProjectID)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreatesACloneOfTheGivenTemplate", err,
			"Failure at CreatesACloneOfTheGivenTemplate, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreatesACloneOfTheGivenTemplate", err))
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
			errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreatesACloneOfTheGivenTemplate", err1))
			return diags
		}
	}

	vItem1 := flattenConfigurationTemplatesCreatesACloneOfTheGivenTemplateItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreatesACloneOfTheGivenTemplate response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}

func resourceConfigurationTemplateCloneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceConfigurationTemplateCloneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceConfigurationTemplateCloneRead(ctx, d, m)
}

func resourceConfigurationTemplateCloneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenConfigurationTemplatesCreatesACloneOfTheGivenTemplateItem(item *dnacentersdkgo.ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateResponse) []map[string]interface{} {
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
