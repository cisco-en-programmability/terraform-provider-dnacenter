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
func dataSourceDeviceConfigurationsExport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Archive.

- Export Device configurations to an encrypted zip file.
`,

		ReadContext: dataSourceDeviceConfigurationsExportRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `Device Id`,
				Type:        schema.TypeList,
				Optional:    true,
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
			"password": &schema.Schema{
				Description: `Password`,
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func dataSourceDeviceConfigurationsExportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ExportDeviceConfigurations")
		request1 := expandRequestDeviceConfigurationsExportExportDeviceConfigurations(ctx, "", d)

		response1, restyResp1, err := client.ConfigurationArchive.ExportDeviceConfigurations(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ExportDeviceConfigurations", err,
				"Failure at ExportDeviceConfigurations, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationArchiveExportDeviceConfigurationsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ExportDeviceConfigurations response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestDeviceConfigurationsExportExportDeviceConfigurations(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestConfigurationArchiveExportDeviceConfigurations {
	request := dnacentersdkgo.RequestConfigurationArchiveExportDeviceConfigurations{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenConfigurationArchiveExportDeviceConfigurationsItem(item *dnacentersdkgo.ResponseConfigurationArchiveExportDeviceConfigurationsResponse) []map[string]interface{} {
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
