package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoverySNMPPropertyAdd() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoverySNMPPropertyAddRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"int_value": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"system_property_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func constructDataSNMPPropertiesRequest(response []interface{}) *[]dnac.CreateUpdateSNMPPropertiesRequest {
	var result []dnac.CreateUpdateSNMPPropertiesRequest
	for _, item := range response {
		ci := item.(map[string]interface{})
		requestItem := dnac.CreateUpdateSNMPPropertiesRequest{}

		if v, ok := ci["id"]; ok {
			requestItem.ID = v.(string)
		}
		if v, ok := ci["instance_tenant_id"]; ok {
			requestItem.InstanceTenantID = v.(string)
		}
		if v, ok := ci["instance_uuid"]; ok {
			requestItem.InstanceUUID = v.(string)
		}
		if v, ok := ci["int_value"]; ok {
			requestItem.IntValue = v.(int)
		}
		if v, ok := ci["system_property_name"]; ok {
			requestItem.SystemPropertyName = v.(string)
		}
		result = append(result, requestItem)
	}
	return &result
}

func dataSourceDiscoverySNMPPropertyAddRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	requests := constructDataSNMPPropertiesRequest(d.Get("items").([]interface{}))

	response, _, err := client.Discovery.CreateUpdateSNMPProperties(requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Call function to check task
	taskID := response.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Check if task was completed successfully
	if taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create/update SNMP properties",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
