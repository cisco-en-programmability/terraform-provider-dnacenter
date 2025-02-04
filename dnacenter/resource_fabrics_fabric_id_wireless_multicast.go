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

func resourceFabricsFabricIDWirelessMulticast() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Fabric Wireless.

- Updates the Software-Defined Access (SDA) Wireless Multicast setting for a specified fabric site. This resource allows
you to enable or disable the multicast feature. For optimal performance, ensure wired multicast is also enabled.
`,

		CreateContext: resourceFabricsFabricIDWirelessMulticastCreate,
		ReadContext:   resourceFabricsFabricIDWirelessMulticastRead,
		UpdateContext: resourceFabricsFabricIDWirelessMulticastUpdate,
		DeleteContext: resourceFabricsFabricIDWirelessMulticastDelete,
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

						"multicast_enabled": &schema.Schema{
							Description: `The setting indicates whether multicast is enabled (true) or disabled (false).
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
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

						"fabric_id": &schema.Schema{
							Description: `fabricId path parameter. The unique identifier of the fabric site for which the multicast setting is being requested. The identifier should be in the format of a UUID. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"multicast_enabled": &schema.Schema{
							Description: `Multicast Enabled`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceFabricsFabricIDWirelessMulticastCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["name"])
	resourceMap["fabric_id"] = interfaceToString(resourceItem["fabric_id"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceFabricsFabricIDWirelessMulticastRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFabricID := resourceMap["fabric_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSdaWirelessMulticast")
		vvFabricID := vFabricID

		response1, restyResp1, err := client.FabricWireless.GetSdaWirelessMulticast(vvFabricID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenFabricWirelessGetSdaWirelessMulticastItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSdaWirelessMulticast response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceFabricsFabricIDWirelessMulticastUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vFabricID := resourceMap["fabric_id"]

	if d.HasChange("parameters") {
		request1 := expandRequestFabricsFabricIDWirelessMulticastUpdateSdaWirelessMulticast(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.FabricWireless.UpdateSdaWirelessMulticast(vFabricID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSdaWirelessMulticast", err, restyResp1.String(),
					"Failure at UpdateSdaWirelessMulticast, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSdaWirelessMulticast", err,
				"Failure at UpdateSdaWirelessMulticast, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateSdaWirelessMulticast", err))
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
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateSdaWirelessMulticast", err1))
				return diags
			}
		}

	}

	return resourceFabricsFabricIDWirelessMulticastRead(ctx, d, m)
}

func resourceFabricsFabricIDWirelessMulticastDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete FabricsFabricIDWirelessMulticast on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestFabricsFabricIDWirelessMulticastUpdateSdaWirelessMulticast(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessUpdateSdaWirelessMulticast {
	request := dnacentersdkgo.RequestFabricWirelessUpdateSdaWirelessMulticast{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".multicast_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".multicast_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".multicast_enabled")))) {
		request.MulticastEnabled = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
