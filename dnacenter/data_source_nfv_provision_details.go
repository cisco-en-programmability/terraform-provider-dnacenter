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
func dataSourceNfvProvisionDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Checks the provisioning detail of an ENCS device including log information.
`,

		ReadContext: dataSourceNfvProvisionDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_ip": &schema.Schema{
				Description: `Device Ip`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"persistbapioutput": &schema.Schema{
				Description: `__persistbapioutput header parameter. Persist bapi sync response
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"runsync": &schema.Schema{
				Description: `__runsync header parameter. Enable this parameter to execute the API and return a response synchronously
			`,
				Type:         schema.TypeString,
				ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
				Optional:     true,
			},
			"runsynctimeout": &schema.Schema{
				Description: `__runsynctimeout header parameter. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
			`,
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}

func dataSourceNfvProvisionDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	vRunsync, okRunsync := d.GetOk("runsync")
	vRunsynctimeout, okRunsynctimeout := d.GetOk("runsynctimeout")
	vPersistbapioutput, okPersistbapioutput := d.GetOk("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: NfvProvisioningDetail")
		request1 := expandRequestNfvProvisionDetailsNfvProvisioningDetail(ctx, "", d)

		headerParams1 := dnacentersdkgo.NfvProvisioningDetailHeaderParams{}

		if okRunsync {
			headerParams1.Runsync = vRunsync.(string)
		}

		if okRunsynctimeout {
			headerParams1.Runsynctimeout = vRunsynctimeout.(string)
		}

		if okPersistbapioutput {
			headerParams1.Persistbapioutput = vPersistbapioutput.(string)
		}

		response1, restyResp1, err := client.SiteDesign.NfvProvisioningDetail(request1, &headerParams1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing NfvProvisioningDetail", err,
				"Failure at NfvProvisioningDetail, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignNfvProvisioningDetailItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting NfvProvisioningDetail response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestNfvProvisionDetailsNfvProvisioningDetail(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignNfvProvisioningDetail {
	request := dnacentersdkgo.RequestSiteDesignNfvProvisioningDetail{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_ip")))) {
		request.DeviceIP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func flattenSiteDesignNfvProvisioningDetailItem(item *dnacentersdkgo.ResponseSiteDesignNfvProvisioningDetail) []map[string]interface{} {
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
