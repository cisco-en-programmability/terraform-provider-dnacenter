package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceSiteAssignCredential() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Network Settings.

- Assign Device Credential To Site
`,

		ReadContext: dataSourceSiteAssignCredentialRead,
		Schema: map[string]*schema.Schema{
			"persistbapioutput": &schema.Schema{
				Description: `__persistbapioutput header parameter. Persist bapi sync response
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. site id to assign credential.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"cli_id": &schema.Schema{
				Description: `Cli Id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"http_read": &schema.Schema{
				Description: `Http Read`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"http_write": &schema.Schema{
				Description: `Http Write`,
				Type:        schema.TypeString,
				Optional:    true,
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
			"snmp_v2_read_id": &schema.Schema{
				Description: `Snmp V2 Read Id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snmp_v2_write_id": &schema.Schema{
				Description: `Snmp V2 Write Id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"snmp_v3_id": &schema.Schema{
				Description: `Snmp V3 Id`,
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceSiteAssignCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AssignCredentialToSite")
		vvSiteID := vSiteID.(string)
		request1 := expandRequestSiteAssignCredentialAssignCredentialToSite(ctx, "", d)

		headerParams1 := dnacentersdkgo.AssignCredentialToSiteHeaderParams{}

		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.AssignCredentialToSite(vvSiteID, request1, &headerParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AssignCredentialToSite", err,
				"Failure at AssignCredentialToSite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsAssignCredentialToSiteItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AssignCredentialToSite response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestSiteAssignCredentialAssignCredentialToSite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsAssignCredentialToSite {
	request := dnacentersdkgo.RequestNetworkSettingsAssignCredentialToSite{}
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
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenNetworkSettingsAssignCredentialToSiteItem(item *dnacentersdkgo.ResponseNetworkSettingsAssignCredentialToSite) []map[string]interface{} {
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
