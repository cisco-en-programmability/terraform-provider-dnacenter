package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplicationSet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationSetCreate,
		ReadContext:   resourceApplicationSetRead,
		UpdateContext: resourceApplicationSetUpdate,
		DeleteContext: resourceApplicationSetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"item": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_source": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceApplicationSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	var requests []dnac.CreateApplicationSetRequest
	var request dnac.CreateApplicationSetRequest

	item := d.Get("item").([]interface{})[0]
	applicationSetRequest := item.(map[string]interface{})
	name := applicationSetRequest["name"].(string)

	queryParams := &dnac.GetApplicationSetsQueryParams{
		Name: name,
	}
	searchResponse, _, err := client.ApplicationPolicy.GetApplicationSets(queryParams)
	if err == nil && searchResponse != nil && len(searchResponse.Response) > 0 {
		// Update resource id
		d.SetId(name)
		// Update resource on Terraform
		resourceApplicationSetRead(ctx, d, m)
		return diags
	}

	request.Name = name

	requests = append(requests, request)

	response, _, err := client.ApplicationPolicy.CreateApplicationSet(&requests)
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
			Summary:  "Unable to create application set",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// Update resource id
	d.SetId(name)
	// Update resource on Terraform
	resourceApplicationSetRead(ctx, d, m)
	return diags
}

func resourceApplicationSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics
	appSetName := d.Id()
	queryParams := &dnac.GetApplicationSetsQueryParams{
		Name: appSetName,
	}

	response, _, err := client.ApplicationPolicy.GetApplicationSets(queryParams)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	if len(response.Response) == 0 || response.Response[0].ID == "" {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	applicationSetItem := flattenApplicationSetReadItem(&response.Response[0])
	if err := d.Set("item", applicationSetItem); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceApplicationSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceApplicationSetRead(ctx, d, m)
}

func resourceApplicationSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	name := d.Id()
	queryParams := &dnac.GetApplicationSetsQueryParams{
		Name: name,
	}
	response, _, err := client.ApplicationPolicy.GetApplicationSets(queryParams)
	if err != nil {
		return diags
	}
	if len(response.Response) == 0 {
		return diags
	}
	if response.Response[0].ID == "" {
		return diags
	}

	appSetID := response.Response[0].ID

	deleteApplicationSetQueryParams := &dnac.DeleteApplicationSetQueryParams{
		ID: appSetID,
	}

	// Call function to delete application resource
	_, _, err = client.ApplicationPolicy.DeleteApplicationSet(deleteApplicationSetQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	response, _, err = client.ApplicationPolicy.GetApplicationSets(queryParams)
	if err != nil {
		return diags
	}
	if len(response.Response) == 0 {
		return diags
	}
	if response.Response[0].ID == "" {
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to delete application set",
		Detail:   "",
	})

	return diags
}
