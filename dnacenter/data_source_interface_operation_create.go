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
func dataSourceInterfaceOperationCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Clear mac-address on an individual port. In request body, operation needs to be specified as 'ClearMacAddress'. In the
future more possible operations will be added to this API
`,

		ReadContext: dataSourceInterfaceOperationCreateRead,
		Schema: map[string]*schema.Schema{
			"deployment_mode": &schema.Schema{
				Description: `deploymentMode query parameter. Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_uuid": &schema.Schema{
				Description: `interfaceUuid path parameter. Interface Id
`,
				Type:     schema.TypeString,
				Required: true,
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
			"operation": &schema.Schema{
				Description: `Operation`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"payload": &schema.Schema{
				Description: `Payload`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceInterfaceOperationCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vInterfaceUUID := d.Get("interface_uuid")
	vDeploymentMode, okDeploymentMode := d.GetOk("deployment_mode")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ClearMacAddressTable")
		vvInterfaceUUID := vInterfaceUUID.(string)
		request1 := expandRequestInterfaceOperationCreateClearMacAddressTable(ctx, "", d)
		queryParams1 := dnacentersdkgo.ClearMacAddressTableQueryParams{}

		if okDeploymentMode {
			queryParams1.DeploymentMode = vDeploymentMode.(string)
		}

		response1, restyResp1, err := client.Devices.ClearMacAddressTable(vvInterfaceUUID, request1, &queryParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ClearMacAddressTable", err,
				"Failure at ClearMacAddressTable, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesClearMacAddressTableItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ClearMacAddressTable response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestInterfaceOperationCreateClearMacAddressTable(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesClearMacAddressTable {
	request := dnacentersdkgo.RequestDevicesClearMacAddressTable{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operation")))) {
		request.Operation = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".payload")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".payload")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".payload")))) {
		request.Payload = expandRequestInterfaceOperationCreateClearMacAddressTablePayload(ctx, key+".payload.0", d)
	}
	return &request
}

func expandRequestInterfaceOperationCreateClearMacAddressTablePayload(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesClearMacAddressTablePayload {
	var request dnacentersdkgo.RequestDevicesClearMacAddressTablePayload
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenDevicesClearMacAddressTableItem(item *dnacentersdkgo.ResponseDevicesClearMacAddressTableResponse) []map[string]interface{} {
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
