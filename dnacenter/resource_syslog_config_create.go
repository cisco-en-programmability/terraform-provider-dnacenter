package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSyslogConfigCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Event Management.

- Create Syslog Destination
`,

		CreateContext: resourceSyslogConfigCreateCreate,
		ReadContext:   resourceSyslogConfigCreateRead,
		DeleteContext: resourceSyslogConfigCreateDelete,
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

						"api_status": &schema.Schema{
							Description: `Api Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"error_message": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"errors": &schema.Schema{
										Description: `Errors`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"status_message": &schema.Schema{
							Description: `Status Message`,
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
						"config_id": &schema.Schema{
							Description: `Required only for update syslog configuration
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"host": &schema.Schema{
							Description: `Host`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"port": &schema.Schema{
							Description: `Port`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"protocol": &schema.Schema{
							Description: `Protocol`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSyslogConfigCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSyslogConfigCreateCreateSyslogDestination(ctx, "parameters.0", d)

	response1, restyResp1, err := client.EventManagement.CreateSyslogDestination(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateSyslogDestination", err,
			"Failure at CreateSyslogDestination, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	//Analizar verificacion.

	vItem1 := flattenEventManagementCreateSyslogDestinationItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreateSyslogDestination response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceSyslogConfigCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSyslogConfigCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSyslogConfigCreateCreateSyslogDestination(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestEventManagementCreateSyslogDestination {
	request := dnacentersdkgo.RequestEventManagementCreateSyslogDestination{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host")))) {
		request.Host = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToString(v)
	}
	return &request
}

func flattenEventManagementCreateSyslogDestinationItem(item *dnacentersdkgo.ResponseEventManagementCreateSyslogDestination) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["error_message"] = flattenEventManagementCreateSyslogDestinationItemErrorMessage(item.ErrorMessage)
	respItem["api_status"] = item.APIStatus
	respItem["status_message"] = item.StatusMessage
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEventManagementCreateSyslogDestinationItemErrorMessage(item *dnacentersdkgo.ResponseEventManagementCreateSyslogDestinationErrorMessage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["errors"] = item.Errors

	return []map[string]interface{}{
		respItem,
	}

}
