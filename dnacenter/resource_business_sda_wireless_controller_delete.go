package dnacenter

import (
	"context"
	"time"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBusinessSdaWirelessControllerDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Fabric Wireless.
		- Remove WLC from Fabric Domain
`,

		CreateContext: resourceBusinessSdaWirelessControllerDeleteCreate,
		ReadContext:   resourceBusinessSdaWirelessControllerDeleteRead,
		DeleteContext: resourceBusinessSdaWirelessControllerDeleteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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
	vDeviceIpAddress := resourceItem["device_ipaddress"]
	vvDeviceIpAddress := interfaceToString(vDeviceIpAddress)

	queryParams1 := dnacentersdkgo.RemoveWLCFromFabricDomainQueryParams{}

	queryParams1.DeviceIPAddress = vvDeviceIpAddress

	response1, restyResp1, err := client.FabricWireless.RemoveWLCFromFabricDomain(&queryParams1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RemoveWLCFromFabricDomain", err,
			"Failure at RemoveWLCFromFabricDomain, unexpected response", ""))
		return diags
	}
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenFabricWirelessRemoveWLCFromFabricDomainItem(response1)
	if err := d.Set("item", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RemoveWLCFromFabricDomain response",
			err))
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
				"Failure when executing RemoveWLCFromFabricDomain", err,
				"Failure at RemoveWLCFromFabricDomain execution", bapiError))
			return diags
		}
	}

	d.SetId(getUnixTimeString())
	return resourceBusinessSdaWirelessControllerDeleteRead(ctx, d, m)
}

func resourceBusinessSdaWirelessControllerDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceBusinessSdaWirelessControllerDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
