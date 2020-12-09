package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func pnpWorkflowsWorkItems() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"output_str": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_taken": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func pnpWorkflowsWorkflowTasks() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"curr_work_item_idx": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"end_time": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"start_time": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"state": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"task_seq_no": &schema.Schema{
					Type:     schema.TypeInt,
					Computed: true,
				},
				"time_taken": &schema.Schema{
					Type:     schema.TypeFloat,
					Computed: true,
				},
				"type": &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
				"work_item_list": &schema.Schema{
					Type:     schema.TypeList,
					Computed: true,
					Elem:     pnpWorkflowsWorkItems(),
				},
			},
		},
	}
}

func pnpWorkflowsWorkflow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"add_to_inventory": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"added_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"curr_task_idx": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"exec_time": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lastupdate_on": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tasks": pnpWorkflowsWorkflowTasks(),
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourcePnPWorkflow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPWorkflowRead,
		Schema: map[string]*schema.Schema{
			"offset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sort": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_order": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"asc", "des"}),
			},
			"type": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     pnpWorkflowsWorkflow(),
			},
		},
	}
}

func dataSourcePnPWorkflowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParam := dnac.GetWorkflowsQueryParams{}
	if v, ok := d.GetOk("limit"); ok {
		queryParam.Limit = v.(int)
	}
	if v, ok := d.GetOk("offset"); ok {
		queryParam.Offset = v.(int)
	}
	if v, ok := d.GetOk("sort"); ok {
		queryParam.Sort = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("sort_order"); ok {
		queryParam.SortOrder = v.(string)
	}
	if v, ok := d.GetOk("type"); ok {
		queryParam.Type = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("name"); ok {
		queryParam.Name = convertSliceInterfaceToSliceString(v.([]interface{}))
	}

	response, _, err := client.DeviceOnboardingPnP.GetWorkflows(&queryParam)
	if err != nil {
		return diag.FromErr(err)
	}

	sItem := flattenPnPWorkflowsReadItems(response)
	if err := d.Set("items", sItem); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
