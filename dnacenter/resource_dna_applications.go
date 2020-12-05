package dnacenter

import (
	"context"
	"time"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_set_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"application_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"application_network_applications": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_protocol": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"application_subtype": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"application_type": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"category_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"dscp": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"engine_id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"help_string": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ignore_conflict": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"long_description": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"popularity": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
									"rank": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
									"server_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"traffic_class": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"url": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"application_network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"lower_port": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
									"ports": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"protocol": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"upper_port": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
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

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	var requests []dnac.CreateApplicationRequest
	var request dnac.CreateApplicationRequest
	name := d.Get("name").(string)
	items := d.Get("items").([]interface{})
	request.Name = name

	for _, item := range items {
		i := item.(map[string]interface{})

		applicationSetID := i["application_set_id"].(string)
		request.ApplicationSet.IDRef = applicationSetID

		applications := i["application_network_applications"].([]interface{})
		identities := i["application_network_identity"].([]interface{})

		for _, application := range applications {
			j := application.(map[string]interface{})
			na := dnac.CreateApplicationRequestNetworkApplications{
				AppProtocol:        j["app_protocol"].(string),
				ApplicationSubType: j["application_subtype"].(string),
				ApplicationType:    j["application_type"].(string),
				CategoryID:         j["category_id"].(string),
				DisplayName:        j["display_name"].(string),
				Dscp:               j["dscp"].(string),
				EngineID:           j["engine_id"].(string),
				HelpString:         j["help_string"].(string),
				IgnoreConflict:     j["ignore_conflict"].(string),
				LongDescription:    j["long_description"].(string),
				Name:               j["name"].(string),
				Popularity:         j["popularity"].(int),
				Rank:               j["rank"].(int),
				ServerName:         j["server_name"].(string),
				TrafficClass:       j["traffic_class"].(string),
				URL:                j["url"].(string),
			}
			request.NetworkApplications = append(request.NetworkApplications, na)
		}
		for _, identity := range identities {
			k := identity.(map[string]interface{})
			ni := dnac.CreateApplicationRequestNetworkIDentity{
				DisplayName: k["display_name"].(string),
				LowerPort:   k["lower_port"].(int),
				Ports:       k["ports"].(string),
				Protocol:    k["protocol"].(string),
				UpperPort:   k["upper_port"].(int),
			}
			request.NetworkIDentity = append(request.NetworkIDentity, ni)
		}
	}

	requests = append(requests, request)

	response, _, err := client.ApplicationPolicy.CreateApplication(&requests)
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
			Summary:  "Unable to create application",
			Detail:   taskResponse.Response.FailureReason,
		})
		return diags
	}

	// Update resource id
	d.SetId(name)
	// Update resource on Terraform
	resourceApplicationRead(ctx, d, m)
	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics
	appName := d.Id()
	applicationsQueryParams := &dnac.GetApplicationsQueryParams{
		Name: appName,
	}

	response, _, err := client.ApplicationPolicy.GetApplications(applicationsQueryParams)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}
	if len(response.Response) == 0 {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	applicationItems := flattenApplicationsReadItems(&response.Response)
	if err := d.Set("items", applicationItems); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Check if properties inside resource has changes
	if d.HasChange("items") {
		var requests []dnac.EditApplicationRequest
		var request dnac.EditApplicationRequest
		name := d.Id()
		items := d.Get("items").([]interface{})
		request.Name = name

		for _, item := range items {
			i := item.(map[string]interface{})

			applicationID := i["application_id"].(string)
			applicationSetID := i["application_set_id"].(string)
			request.ApplicationSet.IDRef = applicationSetID
			request.ID = applicationID

			applications := i["application_network_applications"].([]interface{})
			identities := i["application_network_identity"].([]interface{})

			for _, application := range applications {
				j := application.(map[string]interface{})
				na := dnac.EditApplicationRequestNetworkApplications{
					AppProtocol:        j["app_protocol"].(string),
					ApplicationSubType: j["application_subtype"].(string),
					ApplicationType:    j["application_type"].(string),
					CategoryID:         j["category_id"].(string),
					DisplayName:        j["display_name"].(string),
					Dscp:               j["dscp"].(string),
					EngineID:           j["engine_id"].(string),
					HelpString:         j["help_string"].(string),
					IgnoreConflict:     j["ignore_conflict"].(string),
					LongDescription:    j["long_description"].(string),
					Name:               j["name"].(string),
					Popularity:         j["popularity"].(int),
					Rank:               j["rank"].(int),
					ServerName:         j["server_name"].(string),
					TrafficClass:       j["traffic_class"].(string),
					URL:                j["url"].(string),
				}
				request.NetworkApplications = append(request.NetworkApplications, na)
			}
			for _, identity := range identities {
				k := identity.(map[string]interface{})
				ni := dnac.EditApplicationRequestNetworkIDentity{
					DisplayName: k["display_name"].(string),
					LowerPort:   k["lower_port"].(int),
					Ports:       k["ports"].(string),
					Protocol:    k["protocol"].(string),
					UpperPort:   k["upper_port"].(int),
				}
				request.NetworkIDentity = append(request.NetworkIDentity, ni)
			}
		}

		requests = append(requests, request)

		response, _, err := client.ApplicationPolicy.EditApplication(&requests)
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
				Summary:  "Unable to update application",
				Detail:   taskResponse.Response.FailureReason,
			})
			return diags
		}
		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceApplicationRead(ctx, d, m)
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	name := d.Id()
	applicationsQueryParams := &dnac.GetApplicationsQueryParams{
		Name: name,
	}

	response, _, err := client.ApplicationPolicy.GetApplications(applicationsQueryParams)
	if err != nil {
		return diags
	}
	if len(response.Response) == 0 {
		return diags
	}

	appID := response.Response[0].ID

	deleteApplicationQueryParams := &dnac.DeleteApplicationQueryParams{
		ID: appID,
	}

	// Call function to delete application resource
	_, _, err = client.ApplicationPolicy.DeleteApplication(deleteApplicationQueryParams)
	if err != nil {
		return diag.FromErr(err)
	}

	response, _, err = client.ApplicationPolicy.GetApplications(applicationsQueryParams)
	if err != nil {
		return diags
	}
	if len(response.Response) == 0 {
		return diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to delete application",
		Detail:   "",
	})

	return diags
}
