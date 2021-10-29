package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func templateDeployParam(isMain bool) *schema.Resource {
	if isMain {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{
				"force_push_template": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"is_composite": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
				},
				"main_template_id": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"template_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"target_info": &schema.Schema{
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"hostname": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
							},
							"id": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
							},
							"params": &schema.Schema{
								Type:     schema.TypeMap,
								Optional: true,
							},
							"type": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
			},
		}
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"force_push_template": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_composite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"main_template_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"params": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTemplateDeploy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateDeployRead,
		Schema: map[string]*schema.Schema{
			"template_deployment_info": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem:     templateDeployParam(true),
			},
			"member_templates_deployment_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     templateDeployParam(false),
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deployment_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"deployment_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"duration": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"duration": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
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

func getDeployTemplateRequest(deployTemplate map[string]interface{}) *dnac.DeployTemplateRequest {
	deployTemplateRequest := dnac.DeployTemplateRequest{}
	if v, ok := deployTemplate["force_push_template"]; ok && v != nil {
		deployTemplateRequest.ForcePushTemplate = v.(bool)
	}
	if v, ok := deployTemplate["is_composite"]; ok && v != nil {
		deployTemplateRequest.IsComposite = v.(bool)
	}
	if v, ok := deployTemplate["main_template_id"]; ok && v != nil {
		deployTemplateRequest.MainTemplateID = v.(string)
	}
	if v, ok := deployTemplate["template_id"]; ok && v != nil {
		deployTemplateRequest.TemplateID = v.(string)
	}
	if v, ok := deployTemplate["target_info"]; ok && v != nil {
		targetInfoList := v.([]interface{})
		for _, ti := range targetInfoList {
			targetInfo := ti.(map[string]interface{})
			targetInfoRequest := dnac.DeployTemplateRequestTargetInfo{}
			if v, ok := targetInfo["hostname"]; ok && v != nil {
				targetInfoRequest.HostName = v.(string)
			}
			if v, ok := targetInfo["id"]; ok && v != nil {
				targetInfoRequest.ID = v.(string)
			}
			if v, ok := targetInfo["params"]; ok && v != nil {
				targetInfoRequest.Params = v.(map[string]interface{})
			}
			if v, ok := targetInfo["type"]; ok && v != nil {
				targetInfoRequest.Type = v.(string)
			}
			deployTemplateRequest.TargetInfo = append(deployTemplateRequest.TargetInfo, targetInfoRequest)
		}
	}
	return &deployTemplateRequest
}

func dataSourceTemplateDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	deployTemplate := d.Get("template_deployment_info").([]interface{})[0]
	deployTemplateMap := deployTemplate.(map[string]interface{})

	deployTemplateRequest := *getDeployTemplateRequest(deployTemplateMap)
	if v, ok := d.GetOk("member_templates_deployment_info"); ok {
		var memberDeployTemplates []dnac.DeployTemplateRequest
		memberTemplatesDeploymentInfo := v.([]interface{})
		for _, mt := range memberTemplatesDeploymentInfo {
			memberTemplate := mt.(map[string]interface{})
			memberDeployTemplates = append(memberDeployTemplates, *getDeployTemplateRequest(memberTemplate))
		}
		deployTemplateRequest.MemberTemplateDeploymentInfo = &memberDeployTemplates
	}

	response, _, err := client.ConfigurationTemplates.DeployTemplate(&deployTemplateRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	templateDeploy := flattenDeployTemplateReadItem(response)
	if err := d.Set("item", templateDeploy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
