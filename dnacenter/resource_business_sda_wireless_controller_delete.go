package dnacenter

import (
	"context"
	"time"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceBusinessSdaWirelessControllerDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Fabric Wireless.

- Remove WLC from Fabric Domain
`,

		CreateContext: resourceBusinessSdaWirelessControllerDeleteCreate,
		ReadContext:   resourceBusinessSdaWirelessControllerDeleteRead,
		DeleteContext: resourceBusinessSdaWirelessControllerDeleteDelete,
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
						"device_ipaddress": &schema.Schema{
							Description: `deviceIPAddress query parameter. Device Management IP Address
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"persistbapioutput": &schema.Schema{
							Description: `deviceIPAddress query parameter. Device Management IP Address
`,
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							Default:      "false",
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
						},
					},
				},
			},
		},
	}
}

func resourceBusinessSdaWirelessControllerDeleteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vDeviceIPAddress := resourceItem["device_ipaddress"]

	vPersistbapioutput := resourceItem["persistbapioutput"]

	headerParams1 := dnacentersdkgo.RemoveWLCFromFabricDomainHeaderParams{}
	queryParams1 := dnacentersdkgo.RemoveWLCFromFabricDomainQueryParams{}

	queryParams1.DeviceIPAddress = vDeviceIPAddress.(string)

	headerParams1.Persistbapioutput = vPersistbapioutput.(string)

	response1, restyResp1, err := client.FabricWireless.RemoveWLCFromFabricDomain(&headerParams1, &queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing RemoveWLCFromFabricDomain", err))
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
				"Failure when executing RemoveWLCFromFabricDomain", err,
				"Failure at RemoveWLCFromFabricDomain execution", bapiError))
			return diags
		}
	}

	vItem1 := flattenFabricWirelessRemoveWLCFromFabricDomainItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RemoveWLCFromFabricDomain response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceBusinessSdaWirelessControllerDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceBusinessSdaWirelessControllerDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenFabricWirelessRemoveWLCFromFabricDomainItem(item *dnacentersdkgo.ResponseFabricWirelessRemoveWLCFromFabricDomain) []map[string]interface{} {
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
