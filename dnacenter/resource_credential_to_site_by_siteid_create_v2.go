package dnacenter

import (
	"context"

	"errors"

	"time"

	"reflect"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceCredentialToSiteBySiteidCreateV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Settings.

- API to assign Device Credential to a site.
`,

		CreateContext: resourceCredentialToSiteBySiteidCreateV2Create,
		ReadContext:   resourceCredentialToSiteBySiteidCreateV2Read,
		DeleteContext: resourceCredentialToSiteBySiteidCreateV2Delete,
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

						"task_id": &schema.Schema{
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site Id to assign credential.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"cli_id": &schema.Schema{
							Description: `CLI Credential Id
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"http_read": &schema.Schema{
							Description: `HTTP(S) Read Credential Id
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"http_write": &schema.Schema{
							Description: `HTTP(S) Write Credential Id
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"snmp_v2_read_id": &schema.Schema{
							Description: `SNMPv2c Read Credential Id
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"snmp_v2_write_id": &schema.Schema{
							Description: `SNMPv2c Write Credential Id
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"snmp_v3_id": &schema.Schema{
							Description: `SNMPv3 Credential Id
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceCredentialToSiteBySiteidCreateV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vSiteID := resourceItem["site_id"]

	vvSiteID := vSiteID.(string)
	request1 := expandRequestCredentialToSiteBySiteidCreateV2AssignDeviceCredentialToSiteV2(ctx, "parameters.0", d)

	response1, restyResp1, err := client.NetworkSettings.AssignDeviceCredentialToSiteV2(vvSiteID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AssignDeviceCredentialToSiteV2", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignDeviceCredentialToSiteV2", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing AssignDeviceCredentialToSiteV2", err1))
			return diags
		}
	}

	vItem1 := flattenNetworkSettingsAssignDeviceCredentialToSiteV2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignDeviceCredentialToSiteV2 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceCredentialToSiteBySiteidCreateV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceCredentialToSiteBySiteidCreateV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestCredentialToSiteBySiteidCreateV2AssignDeviceCredentialToSiteV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsAssignDeviceCredentialToSiteV2 {
	request := dnacentersdkgo.RequestNetworkSettingsAssignDeviceCredentialToSiteV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_id")))) {
		request.CliID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2_read_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2_read_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2_read_id")))) {
		request.SNMPV2ReadID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2_write_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2_write_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2_write_id")))) {
		request.SNMPV2WriteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v3_id")))) {
		request.SNMPV3ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_read")))) {
		request.HTTPRead = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_write")))) {
		request.HTTPWrite = interfaceToString(v)
	}
	return &request
}

func flattenNetworkSettingsAssignDeviceCredentialToSiteV2Item(item *dnacentersdkgo.ResponseNetworkSettingsAssignDeviceCredentialToSiteV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
