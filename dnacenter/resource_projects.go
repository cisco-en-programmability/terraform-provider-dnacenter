package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceProjects() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Configuration Templates.

- Create a template project.
`,

		CreateContext: resourceProjectsCreate,
		ReadContext:   resourceProjectsRead,
		UpdateContext: resourceProjectsUpdate,
		DeleteContext: resourceProjectsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"description": &schema.Schema{
							Description: `Description of the project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Timestamp of when the project was updated or modified
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of the project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_id": &schema.Schema{
							Description: `UUID of the project
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description of the project
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of the project
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceProjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestProjectsCreateTemplateProject(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	queryParamImport := dnacentersdkgo.GetTemplateProjectsQueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchConfigurationTemplatesGetTemplateProjects(m, queryParamImport, vvName)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = item2.Name
		d.SetId(joinResourceID(resourceMap))
		return resourceProjectsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.ConfigurationTemplates.CreateTemplateProject(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTemplateProject", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTemplateProject", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateTemplateProject", err))
		return diags
	}
	taskId := resp1.Response.TaskID
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
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateTemplateProject", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetTemplateProjectsQueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchConfigurationTemplatesGetTemplateProjects(m, queryParamValidate, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateTemplateProject", err,
			"Failure at CreateTemplateProject, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = item3.Name
	d.SetId(joinResourceID(resourceMap))
	return resourceProjectsRead(ctx, d, m)
}

func resourceProjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTemplateProjects")
		queryParams1 := dnacentersdkgo.GetTemplateProjectsQueryParams{}

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplateProjects(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		item1, err := searchConfigurationTemplatesGetTemplateProjects(m, queryParams1, vName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		vItem1 := flattenConfigurationTemplatesGetTemplateProjectsByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplateProjects search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplateProjectsByIDItem(item *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateProjectsResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["projectId"] = item.ProjectID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["lastUpdateTime"] = item.LastUpdateTime
	return []map[string]interface{}{
		respItem,
	}
}

func resourceProjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceProjectsRead(ctx, d, m)
}

func resourceProjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete Projects on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestProjectsCreateTemplateProject(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateProject {
	request := dnacentersdkgo.RequestConfigurationTemplatesCreateTemplateProject{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchConfigurationTemplatesGetTemplateProjects(m interface{}, queryParams dnacentersdkgo.GetTemplateProjectsQueryParams, name string) (*dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateProjectsResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateProjectsResponse
	var ite *dnacentersdkgo.ResponseConfigurationTemplatesGetTemplateProjects
	if name != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.ConfigurationTemplates.GetTemplateProjects(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if name == item.Name {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.ConfigurationTemplates.GetTemplateProjects(&queryParams)
		}
		return nil, err
	} else if queryParams.Name != "" {
		ite, _, err = client.ConfigurationTemplates.GetTemplateProjects(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.Name == queryParams.Name {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
