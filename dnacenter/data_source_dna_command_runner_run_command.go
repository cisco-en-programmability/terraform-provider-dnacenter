package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCommandRunnerRunCommand() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCommandRunnerRunCommandRead,
		Schema: map[string]*schema.Schema{
			"commands": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"task_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCommandRunnerRunCommandRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	request := dnac.RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest{}
	if v, ok := d.GetOk("commands"); ok {
		request.Commands = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("device_uuids"); ok {
		request.DeviceUUIDs = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("description"); ok {
		request.Description = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		request.Name = v.(string)
	}
	if v, ok := d.GetOk("timeout"); ok {
		request.Timeout = v.(int)
	}

	// Prepare Request
	response, _, err := client.CommandRunner.RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(&request)
	if err != nil {
		return diag.FromErr(err)
	}

	// Call function to check task
	taskID := response.Response.TaskID

	// set response to Terraform data source
	if err := d.Set("task_id", taskID); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
