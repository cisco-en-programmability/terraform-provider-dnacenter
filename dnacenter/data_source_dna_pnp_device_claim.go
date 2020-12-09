package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceClaim() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceClaimRead,
		Schema: map[string]*schema.Schema{
			"config_file_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_claim_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"config_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"config_parameters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"value": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"device_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"license_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"license_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"top_of_stack_serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"file_service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"populate_inventory": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"workflow_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"json_array_response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"json_response": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_code": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func constructDataPnPClaimDeviceClaimListConfigListConfigParameters(response []interface{}) *[]dnac.ClaimDeviceRequestDeviceClaimListConfigListConfigParameters {
	var result []dnac.ClaimDeviceRequestDeviceClaimListConfigListConfigParameters
	for _, item := range response {
		ci := item.(map[string]interface{})
		configItem := dnac.ClaimDeviceRequestDeviceClaimListConfigListConfigParameters{}
		if v, ok := ci["key"]; ok {
			configItem.Key = v.(string)
		}
		if v, ok := ci["value"]; ok {
			configItem.Value = v.(string)
		}
		result = append(result, configItem)
	}
	return &result
}

func constructDataPnPClaimDeviceClaimListConfigList(response []interface{}) *[]dnac.ClaimDeviceRequestDeviceClaimListConfigList {
	var result []dnac.ClaimDeviceRequestDeviceClaimListConfigList
	for _, item := range response {
		ci := item.(map[string]interface{})
		configItem := dnac.ClaimDeviceRequestDeviceClaimListConfigList{}
		if v, ok := ci["config_parameters"]; ok {
			if w := constructDataPnPClaimDeviceClaimListConfigListConfigParameters(v.([]interface{})); w != nil {
				configItem.ConfigParameters = *w
			}
		}
		if v, ok := ci["config_id"]; ok {
			configItem.ConfigID = v.(string)
		}
		result = append(result, configItem)
	}
	return &result
}

func constructDataPnPClaimDeviceClaimList(response []interface{}) *[]dnac.ClaimDeviceRequestDeviceClaimList {
	var result []dnac.ClaimDeviceRequestDeviceClaimList
	for _, item := range response {
		ci := item.(map[string]interface{})
		claimItem := dnac.ClaimDeviceRequestDeviceClaimList{}
		if v, ok := ci["config_list"]; ok {
			if w := constructDataPnPClaimDeviceClaimListConfigList(v.([]interface{})); w != nil {
				claimItem.ConfigList = *w
			}
		}
		if v, ok := ci["device_id"]; ok {
			claimItem.DeviceID = v.(string)
		}
		if v, ok := ci["license_level"]; ok {
			claimItem.LicenseLevel = v.(string)
		}
		if v, ok := ci["license_type"]; ok {
			claimItem.LicenseType = v.(string)
		}
		if v, ok := ci["top_of_stack_serial_number"]; ok {
			claimItem.TopOfStackSerialNumber = v.(string)
		}
		result = append(result, claimItem)
	}
	return &result
}

func dataSourcePnPDeviceClaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	claimDeviceRequest := dnac.ClaimDeviceRequest{}

	if v, ok := d.GetOk("config_file_url"); ok {
		claimDeviceRequest.ConfigFileURL = v.(string)
	}
	if v, ok := d.GetOk("config_id"); ok {
		claimDeviceRequest.ConfigID = v.(string)
	}
	if v, ok := d.GetOk("device_claim_list"); ok {
		if w := constructDataPnPClaimDeviceClaimList(v.([]interface{})); w != nil {
			claimDeviceRequest.DeviceClaimList = *w
		}
	}
	if v, ok := d.GetOk("file_service_id"); ok {
		claimDeviceRequest.FileServiceID = v.(string)
	}
	if v, ok := d.GetOk("image_id"); ok {
		claimDeviceRequest.ImageID = v.(string)
	}
	if v, ok := d.GetOk("image_url"); ok {
		claimDeviceRequest.ImageURL = v.(string)
	}
	if v, ok := d.GetOk("populate_inventory"); ok {
		claimDeviceRequest.PopulateInventory = v.(bool)
	}
	if v, ok := d.GetOk("project_id"); ok {
		claimDeviceRequest.ProjectID = v.(string)
	}
	if v, ok := d.GetOk("workflow_id"); ok {
		claimDeviceRequest.WorkflowID = v.(string)
	}

	response, _, err := client.DeviceOnboardingPnP.ClaimDevice(&claimDeviceRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	templateDeploy := flattenPnPDeviceClaimReadItem(response)
	if err := d.Set("item", templateDeploy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
