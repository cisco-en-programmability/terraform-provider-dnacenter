package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceCount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceCountRead,
		Schema: map[string]*schema.Schema{

			"serial_number": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"onb_state": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cm_state": &schema.Schema{
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
			"pid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"workflow_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"smart_account_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_account_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"last_contact": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func dataSourcePnPDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	queryParam := dnac.GetPnpDeviceCountQueryParams{}

	if v, ok := d.GetOk("serial_number"); ok {
		queryParam.SerialNumber = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("state"); ok {
		queryParam.State = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("onb_state"); ok {
		queryParam.OnbState = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("cm_state"); ok {
		queryParam.CmState = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("name"); ok {
		queryParam.Name = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("pid"); ok {
		queryParam.Pid = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("source"); ok {
		queryParam.Source = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("project_id"); ok {
		queryParam.ProjectID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("workflow_id"); ok {
		queryParam.WorkflowID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("project_name"); ok {
		queryParam.ProjectName = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("workflow_name"); ok {
		queryParam.WorkflowName = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("smart_account_id"); ok {
		queryParam.SmartAccountID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("virtual_account_id"); ok {
		queryParam.VirtualAccountID = convertSliceInterfaceToSliceString(v.([]interface{}))
	}
	if v, ok := d.GetOk("last_contact"); ok {
		queryParam.LastContact = v.(bool)
	}

	// Prepare Request
	response, _, err := client.DeviceOnboardingPnP.GetPnpDeviceCount(&queryParam)
	if err != nil {
		return diag.FromErr(err)
	}

	// set response to Terraform data source
	if err := d.Set("response", response.Response); err != nil {
		return diag.FromErr(err)
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
