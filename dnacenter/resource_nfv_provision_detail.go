package dnacenter

import (
	"context"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/kuba-mazurkiewicz/dnacenter-go-sdk/v5/sdk"

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
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"begin_step": &schema.Schema{
							Description: `Begin Step`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"duration": &schema.Schema{
							Description: `Duration`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status_message": &schema.Schema{
							Description: `Status Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"task_nodes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_template_user_message_dto": &schema.Schema{
										Description: `Cli Template User Message D T O`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"duration": &schema.Schema{
										Description: `Duration`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"end_time": &schema.Schema{
										Description: `End Time`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"error_payload": &schema.Schema{
										Description: `Error Payload`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"next_task": &schema.Schema{
										Description: `Next Task`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"parent_task": &schema.Schema{
										Description: `Parent Task`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"payload": &schema.Schema{
										Description: `Payload`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"provisioned_names": &schema.Schema{
										Description: `Provisioned Names`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"start_time": &schema.Schema{
										Description: `Start Time`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"status": &schema.Schema{
										Description: `Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"status_message": &schema.Schema{
										Description: `Status Message`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"step_ran": &schema.Schema{
										Description: `Step Ran`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"target": &schema.Schema{
										Description: `Target`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"topology": &schema.Schema{
							Description: `Topology`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_ip": &schema.Schema{
							Description: `Device Ip`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNfvProvisionDetailCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNfvProvisionDetailNfvProvisioningDetail(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vDeviceIP := resourceItem["device_ip"]
	vvDeviceIP := interfaceToString(vDeviceIP)
	queryParamImport := dnacentersdkgo.GetDeviceDetailsByIPQueryParams{}
	queryParamImport.DeviceIP = vvDeviceIP
	item2, _, err := client.SiteDesign.GetDeviceDetailsByIP(&queryParamImport)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["device_ip"] = vvDeviceIP
		d.SetId(joinResourceID(resourceMap))
		return resourceNfvProvisionDetailRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.SiteDesign.NfvProvisioningDetail(request1, nil)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing NfvProvisioningDetail", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing NfvProvisioningDetail", err))
		return diags
	}
	executionId := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if statusIsFailure(response2.Status) {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing NfvProvisioningDetail", err))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetDeviceDetailsByIPQueryParams{}
	queryParamValidate.DeviceIP = vvDeviceIP
	item3, _, err := client.SiteDesign.GetDeviceDetailsByIP(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing NfvProvisioningDetail", err,
			"Failure at NfvProvisioningDetail, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["device_ip"] = vvDeviceIP

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
		log.Printf("[DEBUG] Selected method: GetDeviceDetailsByIP")
		queryParams1 := dnacentersdkgo.GetDeviceDetailsByIPQueryParams{}

		queryParams1.DeviceIP = vDeviceIP

		response1, restyResp1, err := client.SiteDesign.GetDeviceDetailsByIP(&queryParams1)

		if err != nil || response1 == nil {
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
