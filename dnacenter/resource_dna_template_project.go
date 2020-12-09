package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateProject() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceTemplateProjectCreate,
		ReadContext:   resourceTemplateProjectRead,
		UpdateContext: resourceTemplateProjectUpdate,
		DeleteContext: resourceTemplateProjectDelete,
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

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_deletable": &schema.Schema{
							Type:             schema.TypeBool,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: diffSuppressAlways(),
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"templates": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"composite": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"language": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_params_order": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"last_update_time": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"latest_version_time": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"project_associated": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"document_database": &schema.Schema{
										Type:     schema.TypeBool,
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

func resourceTemplateProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	item := d.Get("item").([]interface{})[0]
	project := item.(map[string]interface{})

	// Check if element already exists
	projectName := project["name"].(string)
	projectQueryParams := &dnac.GetProjectsQueryParams{
		Name: projectName,
	}

	searchResponse, _, err := client.ConfigurationTemplates.GetProjects(projectQueryParams)
	if err == nil && searchResponse != nil {
		found := false
		for _, project := range *searchResponse {
			if project.Name == projectName {
				found = true
				break
			}
		}
		if found {
			// Update resource id
			d.SetId(projectName)
			// Update resource data
			resourceTemplateProjectRead(ctx, d, m)
			return diags
		}
	}

	// Construct payload from resource schema (item)
	projectRequest := dnac.CreateProjectRequest{
		IsDeletable: project["is_deletable"].(bool),
		Name:        project["name"].(string),
	}

	// Call function to create template project resource
	response, _, err := client.ConfigurationTemplates.CreateProject(&projectRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(10 * time.Second)

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
			Summary:  "Unable to create template project",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// Update resource id
	d.SetId(projectName)
	// Update resource on Terraform
	resourceTemplateProjectRead(ctx, d, m)
	return diags
}

func resourceTemplateProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Get resource id (that's also the value of template project.id)
	projectName := d.Id()
	projectQueryParams := &dnac.GetProjectsQueryParams{
		Name: projectName,
	}

	searchResponse, _, err := client.ConfigurationTemplates.GetProjects(projectQueryParams)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	var singleResponse *dnac.GetProjectsResponse
	found := false
	for _, project := range *searchResponse {
		if project.Name == projectName {
			found = true
			singleResponse = &project
			break
		}
	}
	if !found {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	projectRead := flattenTemplateProjectReadItem(singleResponse)
	if err := d.Set("item", projectRead); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceTemplateProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	projectName := d.Id()
	projectQueryParams := &dnac.GetProjectsQueryParams{
		Name: projectName,
	}

	searchResponse, _, err := client.ConfigurationTemplates.GetProjects(projectQueryParams)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	var singleResponse *dnac.GetProjectsResponse
	found := false
	for _, project := range *searchResponse {
		if project.Name == projectName {
			found = true
			singleResponse = &project
			break
		}
	}
	if !found {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	projectID := singleResponse.ID

	projectRequest := dnac.UpdateProjectRequest{
		ID:   projectID,
		Name: projectName,
	}

	// Check if properties inside resource has changes
	if d.HasChange("item") {
		// Call function to update template project resource
		response, _, err := client.ConfigurationTemplates.UpdateProject(&projectRequest)
		if err != nil {
			return diag.FromErr(err)
		}

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
				Summary:  "Unable to update template project",
				Detail:   taskResponse.Response.FailureReason,
			})
			return diags
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceTemplateProjectRead(ctx, d, m)
}

func resourceTemplateProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	projectName := d.Id()
	projectQueryParams := &dnac.GetProjectsQueryParams{
		Name: projectName,
	}

	searchResponse, _, err := client.ConfigurationTemplates.GetProjects(projectQueryParams)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	var singleResponse *dnac.GetProjectsResponse
	found := false
	for _, project := range *searchResponse {
		if project.Name == projectName {
			found = true
			singleResponse = &project
			break
		}
	}
	if !found {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	projectID := singleResponse.ID

	// Call function to delete template project resource
	deleteResponse, _, err := client.ConfigurationTemplates.DeleteProject(projectID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Call function to check task
	taskID := deleteResponse.Response.TaskID
	taskResponse, _, err := client.Task.GetTaskByID(taskID)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResponse != nil && taskResponse.Response.IsError {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete template project",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	return diags
}
