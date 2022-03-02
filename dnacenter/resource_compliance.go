package dnacenter

import (
	"context"
	"log"
	"reflect"
	"time"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCompliance() *schema.Resource {
	return &schema.Resource{
		Description: `
`,

		CreateContext: resourceComplianceCreate,
		ReadContext:   resourceComplianceRead,
		UpdateContext: resourceComplianceUpdate,
		DeleteContext: resourceComplianceDelete,
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

						"ap_manager_interface_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"associated_wlc_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"boot_date_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_interval": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"collection_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"family": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"inventory_status_detail": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"line_card_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"platform_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_failure_reason": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"reachability_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"series": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_contact": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"snmp_location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_count": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tunnel_udp_port": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"up_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"waas_device_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"categories": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"device_uuids": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"trigger_full": &schema.Schema{

							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func resourceComplianceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	vDeviceUuids := resourceItem["device_uuids"]
	vvDeviceUuids := interfaceToSliceString(vDeviceUuids)
	request1 := expandRequestComplianceCheckRunRunCompliance(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	response1, restyResp1, err := client.Compliance.RunCompliance(request1)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddWLCToFabricDomain", err,
			"Failure at AddWLCToFabricDomain, unexpected response", ""))
		return diags
	}
	if response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddWLCToFabricDomain", err,
			"Failure at AddWLCToFabricDomain, unexpected response", ""))
		return diags
	}
	taskId := response1.Response.TaskID
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
			diags = append(diags, diagError(
				"Failure when executing CreateApplicationPolicyQueuingProfile", err))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["device_uuids"] = vvDeviceUuids[0]
	d.SetId(joinResourceID(resourceMap))
	return resourceComplianceRead(ctx, d, m)
}

func resourceComplianceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["device_uuids"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceByID")
		vvID := vID

		response1, restyResp1, err := client.Devices.GetDeviceByID(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetDeviceByIDItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceComplianceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceComplianceRead(ctx, d, m)
}

func resourceComplianceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete BusinessSdaWireless on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}

func expandRequestComplianceCheckRunRunCompliance(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestComplianceRunCompliance {
	request := dnacentersdkgo.RequestComplianceRunCompliance{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trigger_full")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trigger_full")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trigger_full")))) {
		request.TriggerFull = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuids")))) {
		request.DeviceUUIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
