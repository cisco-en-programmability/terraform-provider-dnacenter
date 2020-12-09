package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryRangeDelete() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDiscoveryRangeDeleteRead,
		Schema: map[string]*schema.Schema{
			"confirm": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Confirm deletion",
				Required:    true,
			},
			"start_index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"records_to_delete": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceDiscoveryRangeDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	confirmed := d.Get("confirm").(bool)
	startIndex := d.Get("start_index").(int)
	recordsToDelete := d.Get("records_to_delete").(int)

	if !confirmed {
		return diags
	}

	// Prepare Request
	response, _, err := client.Discovery.DeleteDiscoveryBySpecifiedRange(startIndex, recordsToDelete)
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
			Summary:  "Unable to delete by given range",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// always run, Set resource id
	// Unix time  forces this resource to refresh during every Terraform apply
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
