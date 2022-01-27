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
func dataSourceBusinessSdaWirelessControllerCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Fabric Wireless.

- Add WLC to Fabric Domain
`,

		ReadContext: dataSourceBusinessSdaWirelessControllerCreateRead,
		Schema: map[string]*schema.Schema{
			"device_name": &schema.Schema{
				Description: `EWLC Device Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"execution_id": &schema.Schema{
							Description: `Status of the job for wireless state change in fabric domain
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_status_url": &schema.Schema{
							Description: `executionStatusURL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"site_name_hierarchy": &schema.Schema{
				Description: `Site Name Hierarchy
`,
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceBusinessSdaWirelessControllerCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: AddWLCToFabricDomain")
		request1 := expandRequestBusinessSdaWirelessControllerCreateAddWLCToFabricDomain(ctx, "", d)

		response1, restyResp1, err := client.FabricWireless.AddWLCToFabricDomain(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AddWLCToFabricDomain", err,
				"Failure at AddWLCToFabricDomain, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenFabricWirelessAddWLCToFabricDomainItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting AddWLCToFabricDomain response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestBusinessSdaWirelessControllerCreateAddWLCToFabricDomain(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessAddWLCToFabricDomain {
	request := dnacentersdkgo.RequestFabricWirelessAddWLCToFabricDomain{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_name")))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenFabricWirelessAddWLCToFabricDomainItems(items *dnacentersdkgo.ResponseFabricWirelessAddWLCToFabricDomain) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	// for _, item := range *items {
	respItem := make(map[string]interface{})
	respItem["execution_id"] = items.ExecutionID
	respItem["execution_status_url"] = items.ExecutionStatusURL
	respItem["message"] = items.Message
	respItems = append(respItems, respItem)
	// }
	return respItems
}
