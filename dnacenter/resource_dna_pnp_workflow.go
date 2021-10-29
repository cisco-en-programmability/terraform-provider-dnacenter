package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pnpWorkflowWorkItems() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"output_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time_taken": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func pnpWorkflowWorkflowTasks() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"curr_work_item_idx": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"end_time": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"start_time": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Computed: true,
				},
				"state": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"task_seq_no": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"time_taken": &schema.Schema{
					Type:     schema.TypeFloat,
					Optional: true,
					Computed: true,
				},
				"type": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"work_item_list": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Computed: true,
					Elem:     pnpWorkflowWorkItems(),
				},
			},
		},
	}
}

func pnpWorkflowWorkflow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_to_inventory": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"added_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"curr_task_idx": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"exec_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lastupdate_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tasks": pnpWorkflowWorkflowTasks(),
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePnPWorkflow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePnPWorkflowCreate,
		ReadContext:   resourcePnPWorkflowRead,
		UpdateContext: resourcePnPWorkflowUpdate,
		DeleteContext: resourcePnPWorkflowDelete,
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
				MaxItems: 1,
				Required: true,
				Elem:     pnpWorkflowWorkflow(),
			},
		},
	}
}

///// start construct for add

func constructAddPnPWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.AddAWorkflowRequestTasksWorkItemList {
	var result []dnac.AddAWorkflowRequestTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.AddAWorkflowRequestTasksWorkItemList
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructAddPnPWorkflowTasks(wTasks []interface{}) *[]dnac.AddAWorkflowRequestTasks {
	var result []dnac.AddAWorkflowRequestTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.AddAWorkflowRequestTasks
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
			if w := constructAddPnPWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructAddPnPWorkflow(ws map[string]interface{}) *dnac.AddAWorkflowRequest {
	var workflowItem dnac.AddAWorkflowRequest
	if v, ok := ws["id"]; ok && v != nil {
		workflowItem.TypeID = v.(string)
	}
	if v, ok := ws["add_to_inventory"]; ok && v != nil {
		workflowItem.AddToInventory = v.(bool)
	}
	if v, ok := ws["added_on"]; ok && v != nil {
		workflowItem.AddedOn = v.(float64)
	}
	if v, ok := ws["config_id"]; ok && v != nil {
		workflowItem.ConfigID = v.(string)
	}
	if v, ok := ws["curr_task_idx"]; ok && v != nil {
		workflowItem.CurrTaskIDx = v.(float64)
	}
	if v, ok := ws["description"]; ok && v != nil {
		workflowItem.Description = v.(string)
	}
	if v, ok := ws["end_time"]; ok && v != nil {
		workflowItem.EndTime = v.(int)
	}
	if v, ok := ws["exec_time"]; ok && v != nil {
		workflowItem.ExecTime = v.(float64)
	}
	if v, ok := ws["image_id"]; ok && v != nil {
		workflowItem.ImageID = v.(string)
	}
	if v, ok := ws["instance_type"]; ok && v != nil {
		workflowItem.InstanceType = v.(string)
	}
	if v, ok := ws["lastupdate_on"]; ok && v != nil {
		workflowItem.LastupdateOn = v.(float64)
	}
	if v, ok := ws["name"]; ok && v != nil {
		workflowItem.Name = v.(string)
	}
	if v, ok := ws["start_time"]; ok && v != nil {
		workflowItem.StartTime = v.(int)
	}
	if v, ok := ws["state"]; ok && v != nil {
		workflowItem.State = v.(string)
	}
	if v, ok := ws["tasks"]; ok && v != nil {
		if w := constructAddPnPWorkflowTasks(v.([]interface{})); w != nil {
			workflowItem.Tasks = *w
		}
	}
	if v, ok := ws["tenant_id"]; ok && v != nil {
		workflowItem.TenantID = v.(string)
	}
	if v, ok := ws["type"]; ok && v != nil {
		workflowItem.Type = v.(string)
	}
	if v, ok := ws["use_state"]; ok && v != nil {
		workflowItem.UseState = v.(string)
	}
	if v, ok := ws["version"]; ok && v != nil {
		workflowItem.Version = v.(float64)
	}
	return &workflowItem
}

///// end construct for add
///// start construct for update

func constructUpdatePnPWorkflowTasksWorkItemList(itemList []interface{}) *[]dnac.UpdateWorkflowRequestTasksWorkItemList {
	var result []dnac.UpdateWorkflowRequestTasksWorkItemList
	for _, item := range itemList {
		is := item.(map[string]interface{})
		var workItem dnac.UpdateWorkflowRequestTasksWorkItemList
		if v, ok := is["command"]; ok && v != nil {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok && v != nil {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok && v != nil {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok && v != nil {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok && v != nil {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok && v != nil {
			workItem.TimeTaken = v.(float64)
		}
		result = append(result, workItem)
	}
	return &result
}

func constructUpdatePnPWorkflowTasks(wTasks []interface{}) *[]dnac.UpdateWorkflowRequestTasks {
	var result []dnac.UpdateWorkflowRequestTasks
	for _, wTask := range wTasks {
		wts := wTask.(map[string]interface{})
		var workflowTask dnac.UpdateWorkflowRequestTasks
		if v, ok := wts["curr_work_item_idx"]; ok && v != nil {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok && v != nil {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok && v != nil {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok && v != nil {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok && v != nil {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok && v != nil {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok && v != nil {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok && v != nil {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok && v != nil {
			if w := constructUpdatePnPWorkflowTasksWorkItemList(v.([]interface{})); w != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructUpdatePnPWorkflow(ws map[string]interface{}) *dnac.UpdateWorkflowRequest {
	var workflowItem dnac.UpdateWorkflowRequest
	if v, ok := ws["id"]; ok && v != nil {
		workflowItem.TypeID = v.(string)
	}
	if v, ok := ws["add_to_inventory"]; ok && v != nil {
		workflowItem.AddToInventory = v.(bool)
	}
	if v, ok := ws["added_on"]; ok && v != nil {
		workflowItem.AddedOn = v.(float64)
	}
	if v, ok := ws["config_id"]; ok && v != nil {
		workflowItem.ConfigID = v.(string)
	}
	if v, ok := ws["curr_task_idx"]; ok && v != nil {
		workflowItem.CurrTaskIDx = v.(float64)
	}
	if v, ok := ws["description"]; ok && v != nil {
		workflowItem.Description = v.(string)
	}
	if v, ok := ws["end_time"]; ok && v != nil {
		workflowItem.EndTime = v.(int)
	}
	if v, ok := ws["exec_time"]; ok && v != nil {
		workflowItem.ExecTime = v.(float64)
	}
	if v, ok := ws["image_id"]; ok && v != nil {
		workflowItem.ImageID = v.(string)
	}
	if v, ok := ws["instance_type"]; ok && v != nil {
		workflowItem.InstanceType = v.(string)
	}
	if v, ok := ws["lastupdate_on"]; ok && v != nil {
		workflowItem.LastupdateOn = v.(float64)
	}
	if v, ok := ws["name"]; ok && v != nil {
		workflowItem.Name = v.(string)
	}
	if v, ok := ws["start_time"]; ok && v != nil {
		workflowItem.StartTime = v.(int)
	}
	if v, ok := ws["state"]; ok && v != nil {
		workflowItem.State = v.(string)
	}
	if v, ok := ws["tasks"]; ok && v != nil {
		if w := constructUpdatePnPWorkflowTasks(v.([]interface{})); w != nil {
			workflowItem.Tasks = *w
		}
	}
	if v, ok := ws["tenant_id"]; ok && v != nil {
		workflowItem.TenantID = v.(string)
	}
	if v, ok := ws["type"]; ok && v != nil {
		workflowItem.Type = v.(string)
	}
	if v, ok := ws["use_state"]; ok && v != nil {
		workflowItem.UseState = v.(string)
	}
	if v, ok := ws["version"]; ok && v != nil {
		workflowItem.Version = v.(float64)
	}
	return &workflowItem
}

///// end construct for update

func resourcePnPWorkflowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	item := d.Get("item").([]interface{})[0]
	pnpRequest := item.(map[string]interface{})

	name := pnpRequest["name"].(string)
	queryParam := dnac.GetWorkflowsQueryParams{Name: []string{name}}

	searchResponse, _, err := client.DeviceOnboardingPnP.GetWorkflows(&queryParam)
	if err == nil && searchResponse != nil {
		workflows := *searchResponse
		if len(workflows) > 0 {
			workflowID := workflows[0].TypeID
			updateRequest := constructUpdatePnPWorkflow(pnpRequest)
			_, _, err = client.DeviceOnboardingPnP.UpdateWorkflow(workflowID, updateRequest)
			if err != nil {
				return diag.FromErr(err)
			}

			// Update resource id
			d.SetId(workflowID)
			// Update resource on Terraform
			resourcePnPWorkflowRead(ctx, d, m)
			return diags
		}
	}

	request := constructAddPnPWorkflow(pnpRequest)

	response, _, err := client.DeviceOnboardingPnP.AddAWorkflow(request)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update resource id
	d.SetId(response.TypeID)
	// Update resource on Terraform
	resourcePnPWorkflowRead(ctx, d, m)
	return diags
}

func resourcePnPWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics
	workflowID := d.Id()
	response, _, err := client.DeviceOnboardingPnP.GetWorkflowByID(workflowID)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	if response == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	deviceItem := flattenPnPWorkflowReadItem(response)
	if err := d.Set("item", deviceItem); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourcePnPWorkflowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	workflowID := d.Id()
	searchResponse, _, err := client.DeviceOnboardingPnP.GetWorkflowByID(workflowID)
	if err != nil || searchResponse == nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		item := d.Get("item").([]interface{})[0]
		pnpRequest := item.(map[string]interface{})
		request := constructUpdatePnPWorkflow(pnpRequest)
		_, _, err := client.DeviceOnboardingPnP.UpdateWorkflow(workflowID, request)
		if err != nil {
			return diag.FromErr(err)
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourcePnPWorkflowRead(ctx, d, m)
}

func resourcePnPWorkflowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	workflowID := d.Id()
	response, _, err := client.DeviceOnboardingPnP.GetWorkflowByID(workflowID)
	if err != nil || response == nil {
		return diags
	}

	// Call function to delete application resource
	_, _, err = client.DeviceOnboardingPnP.DeleteWorkflowByID(workflowID)
	if err != nil {
		return diag.FromErr(err)
	}

	response, _, err = client.DeviceOnboardingPnP.GetWorkflowByID(workflowID)
	if err != nil || response == nil {
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to delete PnP workflow",
		Detail:   "",
	})

	return diags
}
