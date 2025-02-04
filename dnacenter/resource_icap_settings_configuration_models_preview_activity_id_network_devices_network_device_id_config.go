package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Sensors.

- Generates the device's CLIs of the ICAP intent for preview and approve prior to deploying the ICAP configuration
intent to the device.  After deploying the configuration intent, generating intent CLIs will not be available for
preview.
`,

		CreateContext: resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigCreate,
		ReadContext:   resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead,
		UpdateContext: resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigUpdate,
		DeleteContext: resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigDelete,
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
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"preview_items": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_preview": &schema.Schema{
										Description: `Config Preview`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"config_type": &schema.Schema{
										Description: `Config Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"error_messages": &schema.Schema{
										Description: `Error Messages`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"network_device_id": &schema.Schema{
							Description: `networkDeviceId path parameter. device id from intent/api/v1/network-device
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"object": &schema.Schema{
							Description: `object`,
							Type:        schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional: true,
						},
						"preview_activity_id": &schema.Schema{
							Description: `previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigGeneratesTheDevicesClisOfTheICapConfigurationIntent(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vPreviewActivityID := resourceItem["preview_activity_id"]
	vvPreviewActivityID := interfaceToString(vPreviewActivityID)
	vNetworkDeviceID := resourceItem["network_device_id"]
	vvNetworkDeviceID := interfaceToString(vNetworkDeviceID)
	item2, _, err := client.Sensors.RetrievesTheDevicesClisOfTheICapintent(vvPreviewActivityID, vvNetworkDeviceID)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["preview_activity_id"] = vvPreviewActivityID
		resourceMap["network_device_id"] = vvNetworkDeviceID
		d.SetId(joinResourceID(resourceMap))
		return resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sensors.GeneratesTheDevicesClisOfTheICapConfigurationIntent(vvPreviewActivityID, vvNetworkDeviceID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing GeneratesTheDevicesClisOfTheICapConfigurationIntent", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing GeneratesTheDevicesClisOfTheICapConfigurationIntent", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing GeneratesTheDevicesClisOfTheICapConfigurationIntent", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing GeneratesTheDevicesClisOfTheICapConfigurationIntent", err1))
			return diags
		}
	}
	item3, _, err := client.Sensors.RetrievesTheDevicesClisOfTheICapintent(vvPreviewActivityID, vvNetworkDeviceID)
	if err != nil || item3 != nil {
		resourceMap := make(map[string]string)
		resourceMap["preview_activity_id"] = vvPreviewActivityID
		resourceMap["network_device_id"] = vvNetworkDeviceID
		d.SetId(joinResourceID(resourceMap))
		return resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead(ctx, d, m)
	}
	item4, _, err := client.Sensors.RetrievesTheDevicesClisOfTheICapintent(vvPreviewActivityID, vvNetworkDeviceID)
	if err != nil || item4 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GeneratesTheDevicesClisOfTheICapConfigurationIntent", err,
			"Failure at GeneratesTheDevicesClisOfTheICapConfigurationIntent, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["preview_activity_id"] = vvPreviewActivityID
	resourceMap["network_device_id"] = vvNetworkDeviceID
	d.SetId(joinResourceID(resourceMap))
	return resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead(ctx, d, m)
}

func resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vPreviewActivityID := resourceMap["preview_activity_id"]

	vNetworkDeviceID := resourceMap["network_device_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheDevicesClisOfTheICapintent")
		vvPreviewActivityID := vPreviewActivityID
		vvNetworkDeviceID := vNetworkDeviceID

		response1, restyResp1, err := client.Sensors.RetrievesTheDevicesClisOfTheICapintent(vvPreviewActivityID, vvNetworkDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSensorsRetrievesTheDevicesClisOfTheICapintentItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheDevicesClisOfTheICapintent response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigRead(ctx, d, m)
}

func resourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete IcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfig on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestIcapSettingsConfigurationModelsPreviewActivityIDNetworkDevicesNetworkDeviceIDConfigGeneratesTheDevicesClisOfTheICapConfigurationIntent(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent {
	request := dnacentersdkgo.RequestSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".object")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".object")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".object")))) {
		request = v.(map[string]interface{})
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
