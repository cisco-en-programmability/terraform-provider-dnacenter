package dnacenter

import (
	"context"

	"log"

	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLanAutomationDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on LAN Automation.
- Invoke this API to stop LAN Automation for the given site.
`,

		CreateContext: resourceLanAutomationDeleteCreate,
		ReadContext:   resourceLanAutomationDeleteRead,
		DeleteContext: resourceLanAutomationDeleteDelete,
		Schema: map[string]*schema.Schema{
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Description: `id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.
			`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"detail": &schema.Schema{
							Description: `Detailed information of the error code.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"error_code": &schema.Schema{
							Description: `Error code value.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Description: `Description of the error code.
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

func resourceLanAutomationDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLanAutomationDeleteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LanAutomation")

		response1, restyResp1, err := client.LanAutomation.LanAutomation(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing LanAutomation", err,
				"Failure at LanAutomation, unexpected response", ""))
			return diags
		}

		//Analizar
		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLanAutomationLanAutomationItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomation response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags
	}
	return diags
}

func resourceLanAutomationDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenLanAutomationLanAutomationItem(item *dnacentersdkgo.ResponseLanAutomationLanAutomationResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["error_code"] = item.ErrorCode
	respItem["message"] = item.Message
	respItem["detail"] = item.Detail
	return []map[string]interface{}{
		respItem,
	}
}
