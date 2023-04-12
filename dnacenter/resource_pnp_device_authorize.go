package dnacenter

import (
	"context"

	"reflect"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePnpDeviceAuthorize() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Cisco DNA Center System.

- Authorizes one of more devices. A device can only be authorized if Authorization is set in Device Settings.
`,

		CreateContext: resourcePnpDeviceAuthorizeCreate,
		ReadContext:   resourcePnpDeviceAuthorizeRead,
		DeleteContext: resourcePnpDeviceAuthorizeDelete,
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

						"json_array_response": &schema.Schema{
							Description: `Json Array Response`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"json_response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"empty": &schema.Schema{
										Description: `Empty`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status_code": &schema.Schema{
							Description: `Status Code`,
							Type:        schema.TypeFloat,
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
						"device_id_list": &schema.Schema{
							Description: `Device Id List`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourcePnpDeviceAuthorizeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpDeviceAuthorizeAuthorizeDevice(ctx, "parameters.0", d)

	response1, restyResp1, err := client.CiscoDnaCenterSystem.AuthorizeDevice(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when setting CreateWebhookDestination response",
			err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	//REVIEW: '- Analizar como se puede comprobar la ejecucion.'
	//Analizar verificacion.

	vItem1 := flattenCiscoDnaCenterSystemAuthorizeDeviceItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AuthorizeDevice response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourcePnpDeviceAuthorizeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpDeviceAuthorizeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestPnpDeviceAuthorizeAuthorizeDevice(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestCiscoDnaCenterSystemAuthorizeDevice {
	request := dnacentersdkgo.RequestCiscoDnaCenterSystemAuthorizeDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id_list")))) {
		request.DeviceIDList = interfaceToSliceString(v)
	}
	return &request
}

func flattenCiscoDnaCenterSystemAuthorizeDeviceItem(item *dnacentersdkgo.ResponseCiscoDnaCenterSystemAuthorizeDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["json_response"] = flattenCiscoDnaCenterSystemAuthorizeDeviceItemJSONResponse(item.JSONResponse)
	respItem["message"] = item.Message
	respItem["status_code"] = item.StatusCode
	respItem["json_array_response"] = item.JSONArrayResponse
	return []map[string]interface{}{
		respItem,
	}
}

func flattenCiscoDnaCenterSystemAuthorizeDeviceItemJSONResponse(item *dnacentersdkgo.ResponseCiscoDnaCenterSystemAuthorizeDeviceJSONResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["empty"] = boolPtrToString(item.Empty)

	return []map[string]interface{}{
		respItem,
	}

}
