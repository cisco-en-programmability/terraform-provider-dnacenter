package dnacenter

import (
	"context"
	"time"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSiteAssignCredential() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on NetworkSettings.

- Assign Device Credential to a site.
`,

		CreateContext: resourceSiteAssignCredentialCreate,
		ReadContext:   resourceSiteAssignCredentialRead,
		DeleteContext: resourceSiteAssignCredentialDelete,
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

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
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
							Description: `siteId path parameter. site id to assign credential.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"cli_id": &schema.Schema{
							Description: `Cli Id`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"http_read": &schema.Schema{
							Description: `Http Read`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"http_write": &schema.Schema{
							Description: `Http Write`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"snmp_v2_read_id": &schema.Schema{
							Description: `Snmp V2 Read Id`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"snmp_v2_write_id": &schema.Schema{
							Description: `Snmp V2 Write Id`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"snmp_v3_id": &schema.Schema{
							Description: `Snmp V3 Id`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSiteAssignCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vSiteID := resourceItem["site_id"]

	vPersistbapioutput := resourceItem["persistbapioutput"]

	vvSiteID := vSiteID.(string)
	request1 := expandRequestSiteAssignCredentialAssignDeviceCredentialToSite(ctx, "parameters.0", d)

	headerParams1 := dnacentersdkgo.AssignDeviceCredentialToSiteHeaderParams{}

	headerParams1.Persistbapioutput = vPersistbapioutput.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.NetworkSettings.AssignDeviceCredentialToSite(vvSiteID, request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AssignDeviceCredentialToSite", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AssignDeviceCredentialToSite", err,
				"Failure at AssignDeviceCredentialToSite execution", bapiError))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenNetworkSettingsAssignDeviceCredentialToSiteItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignDeviceCredentialToSite response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceSiteAssignCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSiteAssignCredentialDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSiteAssignCredentialAssignDeviceCredentialToSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsAssignDeviceCredentialToSite {
	request := dnacentersdkgo.RequestNetworkSettingsAssignDeviceCredentialToSite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_id")))) {
		request.CliID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2_read_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2_read_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2_read_id")))) {
		request.SNMPV2ReadID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2_write_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2_write_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2_write_id")))) {
		request.SNMPV2WriteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_read")))) {
		request.HTTPRead = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".http_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".http_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".http_write")))) {
		request.HTTPWrite = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v3_id")))) {
		request.SNMPV3ID = interfaceToString(v)
	}
	return &request
}

func flattenNetworkSettingsAssignDeviceCredentialToSiteItem(item *dnacentersdkgo.ResponseNetworkSettingsAssignDeviceCredentialToSite) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
