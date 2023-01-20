package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNfvProvisionDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Site Design.

- Checks the provisioning detail of an ENCS device including log information.
`,

		CreateContext: resourceNfvProvisionDetailCreate,
		ReadContext:   resourceNfvProvisionDetailRead,
		UpdateContext: resourceNfvProvisionDetailUpdate,
		DeleteContext: resourceNfvProvisionDetailDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_ip": &schema.Schema{
							Description: `Device Ip`,
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNfvProvisionDetailCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	//var diags diag.Diagnostics

	//resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNfvProvisionDetailNfvProvisioningDetail(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	/*resp1, restyResp1, err := client.SiteDesign.NfvProvisioningDetail(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing NfvProvisioningDetail", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing NfvProvisioningDetail", err))
		return diags
	}*/
	resourceMap := make(map[string]string)
	d.SetId(joinResourceID(resourceMap))
	return resourceNfvProvisionDetailRead(ctx, d, m)
}

func resourceNfvProvisionDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vDeviceIP := resourceMap["device_ip"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceDetailsByIP")
		queryParams1 := dnacentersdkgo.GetDeviceDetailsByIPQueryParams{}

		queryParams1.DeviceIP = vDeviceIP

		response1, restyResp1, _ := client.SiteDesign.GetDeviceDetailsByIP(&queryParams1)

		/*		if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetDeviceDetailsByIP", err,
					"Failure at GetDeviceDetailsByIP, unexpected response", ""))
				return diags
			}*/
		if response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetDeviceDetailsByIPItem(response1.ProvisionDetails)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceDetailsByIP response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNfvProvisionDetailUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceNfvProvisionDetailRead(ctx, d, m)
}

func resourceNfvProvisionDetailDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete NfvProvisionDetail on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestNfvProvisionDetailNfvProvisioningDetail(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignNfvProvisioningDetail {
	request := dnacentersdkgo.RequestSiteDesignNfvProvisioningDetail{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_ip")))) {
		request.DeviceIP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
