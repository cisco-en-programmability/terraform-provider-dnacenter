package dnacenter

import (
	"context"
	"time"

	dnac "dnacenter-go-sdk/sdk"

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
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
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
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructAddPnPWorkflowTasksWorkItemList(v.([]interface{})); v != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructAddPnPWorkflow(ws map[string]interface{}) *dnac.AddAWorkflowRequest {
	var workflowItem dnac.AddAWorkflowRequest
	if v, ok := ws["id"]; ok {
		workflowItem.TypeID = v.(string)
	}
	if v, ok := ws["add_to_inventory"]; ok {
		workflowItem.AddToInventory = v.(bool)
	}
	if v, ok := ws["added_on"]; ok {
		workflowItem.AddedOn = v.(float64)
	}
	if v, ok := ws["config_id"]; ok {
		workflowItem.ConfigID = v.(string)
	}
	if v, ok := ws["curr_task_idx"]; ok {
		workflowItem.CurrTaskIDx = v.(float64)
	}
	if v, ok := ws["description"]; ok {
		workflowItem.Description = v.(string)
	}
	if v, ok := ws["end_time"]; ok {
		workflowItem.EndTime = v.(int)
	}
	if v, ok := ws["exec_time"]; ok {
		workflowItem.ExecTime = v.(float64)
	}
	if v, ok := ws["image_id"]; ok {
		workflowItem.ImageID = v.(string)
	}
	if v, ok := ws["instance_type"]; ok {
		workflowItem.InstanceType = v.(string)
	}
	if v, ok := ws["lastupdate_on"]; ok {
		workflowItem.LastupdateOn = v.(float64)
	}
	if v, ok := ws["name"]; ok {
		workflowItem.Name = v.(string)
	}
	if v, ok := ws["start_time"]; ok {
		workflowItem.StartTime = v.(int)
	}
	if v, ok := ws["state"]; ok {
		workflowItem.State = v.(string)
	}
	if v, ok := ws["tasks"]; ok {
		if w := constructAddPnPWorkflowTasks(v.([]interface{})); v != nil {
			workflowItem.Tasks = *w
		}
	}
	if v, ok := ws["tenant_id"]; ok {
		workflowItem.TenantID = v.(string)
	}
	if v, ok := ws["type"]; ok {
		workflowItem.Type = v.(string)
	}
	if v, ok := ws["use_state"]; ok {
		workflowItem.UseState = v.(string)
	}
	if v, ok := ws["version"]; ok {
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
		if v, ok := is["command"]; ok {
			workItem.Command = v.(string)
		}
		if v, ok := is["end_time"]; ok {
			workItem.EndTime = v.(int)
		}
		if v, ok := is["output_str"]; ok {
			workItem.OutputStr = v.(string)
		}
		if v, ok := is["start_time"]; ok {
			workItem.StartTime = v.(int)
		}
		if v, ok := is["state"]; ok {
			workItem.State = v.(string)
		}
		if v, ok := is["time_taken"]; ok {
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
		if v, ok := wts["curr_work_item_idx"]; ok {
			workflowTask.CurrWorkItemIDx = v.(int)
		}
		if v, ok := wts["end_time"]; ok {
			workflowTask.EndTime = v.(int)
		}
		if v, ok := wts["name"]; ok {
			workflowTask.Name = v.(string)
		}
		if v, ok := wts["start_time"]; ok {
			workflowTask.StartTime = v.(int)
		}
		if v, ok := wts["state"]; ok {
			workflowTask.State = v.(string)
		}
		if v, ok := wts["task_seq_no"]; ok {
			workflowTask.TaskSeqNo = v.(int)
		}
		if v, ok := wts["time_taken"]; ok {
			workflowTask.TimeTaken = v.(float64)
		}
		if v, ok := wts["type"]; ok {
			workflowTask.Type = v.(string)
		}
		if v, ok := wts["work_item_list"]; ok {
			if w := constructUpdatePnPWorkflowTasksWorkItemList(v.([]interface{})); v != nil {
				workflowTask.WorkItemList = *w
			}
		}
		result = append(result, workflowTask)
	}
	return &result
}

func constructUpdatePnPWorkflow(ws map[string]interface{}) *dnac.UpdateWorkflowRequest {
	var workflowItem dnac.UpdateWorkflowRequest
	if v, ok := ws["id"]; ok {
		workflowItem.TypeID = v.(string)
	}
	if v, ok := ws["add_to_inventory"]; ok {
		workflowItem.AddToInventory = v.(bool)
	}
	if v, ok := ws["added_on"]; ok {
		workflowItem.AddedOn = v.(float64)
	}
	if v, ok := ws["config_id"]; ok {
		workflowItem.ConfigID = v.(string)
	}
	if v, ok := ws["curr_task_idx"]; ok {
		workflowItem.CurrTaskIDx = v.(float64)
	}
	if v, ok := ws["description"]; ok {
		workflowItem.Description = v.(string)
	}
	if v, ok := ws["end_time"]; ok {
		workflowItem.EndTime = v.(int)
	}
	if v, ok := ws["exec_time"]; ok {
		workflowItem.ExecTime = v.(float64)
	}
	if v, ok := ws["image_id"]; ok {
		workflowItem.ImageID = v.(string)
	}
	if v, ok := ws["instance_type"]; ok {
		workflowItem.InstanceType = v.(string)
	}
	if v, ok := ws["lastupdate_on"]; ok {
		workflowItem.LastupdateOn = v.(float64)
	}
	if v, ok := ws["name"]; ok {
		workflowItem.Name = v.(string)
	}
	if v, ok := ws["start_time"]; ok {
		workflowItem.StartTime = v.(int)
	}
	if v, ok := ws["state"]; ok {
		workflowItem.State = v.(string)
	}
	if v, ok := ws["tasks"]; ok {
		if w := constructUpdatePnPWorkflowTasks(v.([]interface{})); v != nil {
			workflowItem.Tasks = *w
		}
	}
	if v, ok := ws["tenant_id"]; ok {
		workflowItem.TenantID = v.(string)
	}
	if v, ok := ws["type"]; ok {
		workflowItem.Type = v.(string)
	}
	if v, ok := ws["use_state"]; ok {
		workflowItem.UseState = v.(string)
	}
	if v, ok := ws["version"]; ok {
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

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		workflowID := d.Id()
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
	// Call function to delete application resource
	_, _, err := client.DeviceOnboardingPnP.DeleteWorkflowByID(workflowID)
	if err != nil {
		return diag.FromErr(err)
	}

	response, _, err := client.DeviceOnboardingPnP.GetWorkflowByID(workflowID)
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
