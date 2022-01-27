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
func dataSourceLicenseDeviceRegistration() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Licenses.

- Register device(s) in CSSM(Cisco Smart Software Manager).
`,

		ReadContext: dataSourceLicenseDeviceRegistrationRead,
		Schema: map[string]*schema.Schema{
			"virtual_account_name": &schema.Schema{
				Description: `virtual_account_name path parameter. Name of virtual account
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"device_uuids": &schema.Schema{
				Description: `Comma separated device ids
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"task_id": &schema.Schema{
							Description: `Task id of process
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task URL of process
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLicenseDeviceRegistrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vVirtualAccountName := d.Get("virtual_account_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: DeviceRegistration")
		vvVirtualAccountName := vVirtualAccountName.(string)
		request1 := expandRequestLicenseDeviceRegistrationDeviceRegistration(ctx, "", d)

		response1, restyResp1, err := client.Licenses.DeviceRegistration(vvVirtualAccountName, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeviceRegistration", err,
				"Failure at DeviceRegistration, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesDeviceRegistrationItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DeviceRegistration response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestLicenseDeviceRegistrationDeviceRegistration(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestLicensesDeviceRegistration {
	request := dnacentersdkgo.RequestLicensesDeviceRegistration{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuids")))) {
		request.DeviceUUIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenLicensesDeviceRegistrationItem(item *dnacentersdkgo.ResponseLicensesDeviceRegistrationResponse) []map[string]interface{} {
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
