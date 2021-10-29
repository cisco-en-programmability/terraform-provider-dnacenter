package dnacenter

import (
	"context"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnPDeviceReset() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePnPDeviceResetRead,
		Schema: map[string]*schema.Schema{
			"device_reset_list": &schema.Schema{
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

func constructDataPnPDeviceResetDeviceResetListConfigListConfigParameters(response []interface{}) *[]dnac.ResetDeviceRequestDeviceResetListConfigListConfigParameters {
	var result []dnac.ResetDeviceRequestDeviceResetListConfigListConfigParameters
	for _, item := range response {
		ci := item.(map[string]interface{})
		configItem := dnac.ResetDeviceRequestDeviceResetListConfigListConfigParameters{}
		if v, ok := ci["key"]; ok && v != nil {
			configItem.Key = v.(string)
		}
		if v, ok := ci["value"]; ok && v != nil {
			configItem.Value = v.(string)
		}
		result = append(result, configItem)
	}
	return &result
}

func constructDataPnPDeviceResetDeviceResetListConfigList(response []interface{}) *[]dnac.ResetDeviceRequestDeviceResetListConfigList {
	var result []dnac.ResetDeviceRequestDeviceResetListConfigList
	for _, item := range response {
		ci := item.(map[string]interface{})
		configItem := dnac.ResetDeviceRequestDeviceResetListConfigList{}
		if v, ok := ci["config_parameters"]; ok && v != nil {
			if w := constructDataPnPDeviceResetDeviceResetListConfigListConfigParameters(v.([]interface{})); w != nil {
				configItem.ConfigParameters = *w
			}
		}
		if v, ok := ci["config_id"]; ok && v != nil {
			configItem.ConfigID = v.(string)
		}
		result = append(result, configItem)
	}
	return &result
}

func constructDataPnPDeviceResetDeviceResetList(response []interface{}) *[]dnac.ResetDeviceRequestDeviceResetList {
	var result []dnac.ResetDeviceRequestDeviceResetList
	for _, item := range response {
		ci := item.(map[string]interface{})
		resetItem := dnac.ResetDeviceRequestDeviceResetList{}
		if v, ok := ci["config_list"]; ok && v != nil {
			if w := constructDataPnPDeviceResetDeviceResetListConfigList(v.([]interface{})); w != nil {
				resetItem.ConfigList = *w
			}
		}
		if v, ok := ci["device_id"]; ok && v != nil {
			resetItem.DeviceID = v.(string)
		}
		if v, ok := ci["license_level"]; ok && v != nil {
			resetItem.LicenseLevel = v.(string)
		}
		if v, ok := ci["license_type"]; ok && v != nil {
			resetItem.LicenseType = v.(string)
		}
		if v, ok := ci["top_of_stack_serial_number"]; ok && v != nil {
			resetItem.TopOfStackSerialNumber = v.(string)
		}
		result = append(result, resetItem)
	}
	return &result
}

func dataSourcePnPDeviceResetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resetDeviceRequest := dnac.ResetDeviceRequest{}
	if v, ok := d.GetOk("device_reset_list"); ok {
		if w := constructDataPnPDeviceResetDeviceResetList(v.([]interface{})); w != nil {
			resetDeviceRequest.DeviceResetList = *w
		}
	}
	if v, ok := d.GetOk("project_id"); ok {
		resetDeviceRequest.ProjectID = v.(string)
	}
	if v, ok := d.GetOk("workflow_id"); ok {
		resetDeviceRequest.WorkflowID = v.(string)
	}

	response, _, err := client.DeviceOnboardingPnP.ResetDevice(&resetDeviceRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	templateDeploy := flattenPnPDeviceResetReadItem(response)
	if err := d.Set("item", templateDeploy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
