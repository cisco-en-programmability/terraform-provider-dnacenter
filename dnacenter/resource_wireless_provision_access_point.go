package dnacenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessProvisionAccessPoint() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.
		- Access Point Provision and ReProvision
`,

		CreateContext: resourceWirelessProvisionAccessPointCreate,
		ReadContext:   resourceWirelessProvisionAccessPointRead,
		DeleteContext: resourceWirelessProvisionAccessPointDelete,

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Execution Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"execution_url": &schema.Schema{
							Description: `Execution Url`,
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
						"persistbapioutput": &schema.Schema{
							Description: `__persistbapioutput header parameter. Persist bapi sync response
						`,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestWirelessAPProvision`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_ap_group_name": &schema.Schema{
										Description: `Custom AP group name
			`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"custom_flex_group_name": &schema.Schema{
										Description: `["Custom flex group name"]
			`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_name": &schema.Schema{
										Description: `Device name
			`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"rf_profile": &schema.Schema{
										Description: `Radio frequency profile name
			`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"site_id": &schema.Schema{
										Description: `Site name hierarchy(ex: Global/...)
			`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"site_name_hierarchy": &schema.Schema{
										Description: `Site name hierarchy(ex: Global/...)
			`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"type": &schema.Schema{
										Description: `ApWirelessConfiguration
			`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
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

func resourceWirelessProvisionAccessPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vPersistbapioutput, okPersistbapioutput := d.GetOk("parameters.0.persistbapioutput")

	log.Printf("[DEBUG] Selected method 1: ApProvision")
	request1 := expandRequestWirelessProvisionAccessPointApProvision(ctx, "parameters.0", d)
	headerParams1 := dnacentersdkgo.ApProvisionHeaderParams{}

	if okPersistbapioutput {
		headerParams1.Persistbapioutput = vPersistbapioutput.(string)
	}
	response1, restyResp1, err := client.Wireless.ApProvision(request1, &headerParams1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ApProvision", err,
			"Failure at ApProvision, unexpected response", ""))
		return diags
	}

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
				"Failure when executing ApProvision", err,
				"Failure at ApProvision execution", bapiError))
			return diags
		}
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenWirelessApProvisionItems(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ApProvision response",
			err))
		return diags
	}
	log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	d.SetId(getUnixTimeString())
	return resourceWirelessProvisionAccessPointRead(ctx, d, m)
}

func resourceWirelessProvisionAccessPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	return diags
}

func resourceWirelessProvisionAccessPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessProvisionAccessPointApProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessApProvision {
	request := dnacentersdkgo.RequestWirelessApProvision{}
	if v := expandRequestWirelessProvisionAccessPointApProvisionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProvisionAccessPointApProvisionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemWirelessApProvision {
	request := []dnacentersdkgo.RequestItemWirelessApProvision{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestWirelessProvisionAccessPointApProvisionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestWirelessProvisionAccessPointApProvisionItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemWirelessApProvision {
	request := dnacentersdkgo.RequestItemWirelessApProvision{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_profile")))) {
		request.RfProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_name")))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_ap_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_ap_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_ap_group_name")))) {
		request.CustomApGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_flex_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_flex_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_flex_group_name")))) {
		request.CustomFlexGroupName = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenWirelessApProvisionItems(items *dnacentersdkgo.ResponseWirelessApProvision) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	// for _, item := range *items {
	respItem := make(map[string]interface{})
	respItem["execution_id"] = items.ExecutionID
	respItem["execution_url"] = items.ExecutionURL
	respItem["message"] = items.Message
	respItems = append(respItems, respItem)
	// }
	return respItems
}
