package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalCredentialV2() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Discovery.

- API to update device credentials. Multiple credentials can be passed at once, but only a single credential of a given
type can be passed at once. Please refer sample Request Body for more information.

- API to create new global credentials. Multiple credentials of various types can be passed at once. Please refer sample
Request Body for more information.

- Delete a global credential. Only 'id' of the credential has to be passed.
`,

		CreateContext: resourceGlobalCredentialV2Create,
		ReadContext:   resourceGlobalCredentialV2Read,
		UpdateContext: resourceGlobalCredentialV2Update,
		DeleteContext: resourceGlobalCredentialV2Delete,
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

						"cli_credential": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"enable_password": &schema.Schema{
										Description: `Enable Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"https_read": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"secure": &schema.Schema{
										Description: `Secure`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"https_write": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Sensitive:   true,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"secure": &schema.Schema{
										Description: `Secure`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"snmp_v2c_read": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"read_community": &schema.Schema{
										Description: `Read Community`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"snmp_v2c_write": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"write_community": &schema.Schema{
										Description: `Write Community`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"snmp_v3": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_password": &schema.Schema{
										Description: `Auth Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"comments": &schema.Schema{
										Description: `Comments`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"credential_type": &schema.Schema{
										Description: `Credential Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_tenant_id": &schema.Schema{
										Description: `Instance Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"privacy_password": &schema.Schema{
										Description: `Privacy Password`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"privacy_type": &schema.Schema{
										Description: `Privacy Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"snmp_mode": &schema.Schema{
										Description: `Snmp Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
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

						"cli_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"enable_password": &schema.Schema{
										Description: `Enable Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"https_read": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"https_write": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `id path parameter. Global Credential id	
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"snmp_v2c_read": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"read_community": &schema.Schema{
										Description: `Read Community`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"snmp_v2c_write": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"write_community": &schema.Schema{
										Description: `Write Community`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"snmp_v3": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_password": &schema.Schema{
										Description: `Auth Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"privacy_password": &schema.Schema{
										Description: `Privacy Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"privacy_type": &schema.Schema{
										Description: `Privacy Type`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"snmp_mode": &schema.Schema{
										Description: `Snmp Mode`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceGlobalCredentialV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client := m.(*dnacentersdkgo.Client)

	// var diags diag.Diagnostics

	// resourceItem := *getResourceItem(d.Get("parameters"))
	// request1 := expandRequestGlobalCredentialV2CreateGlobalCredentialsV2(ctx, "parameters.0", d)
	// log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	// //review new resource.
	// vID := resourceItem["id"]
	// vvID := interfaceToString(vID)
	// item2, _, err := client.Discovery.GetAllGlobalCredentialsV2(&queryParamImport)
	// if err == nil && item2 != nil {
	// 	resourceMap := make(map[string]string)
	// 	d.SetId(joinResourceID(resourceMap))
	// 	return resourceGlobalCredentialV2Read(ctx, d, m)
	// }
	// resp1, restyResp1, err := client.Discovery.CreateGlobalCredentialsV2(request1)
	// if err != nil || resp1 == nil {
	// 	if restyResp1 != nil {
	// 		diags = append(diags, diagErrorWithResponse(
	// 			"Failure when executing CreateGlobalCredentialsV2", err, restyResp1.String()))
	// 		return diags
	// 	}
	// 	diags = append(diags, diagError(
	// 		"Failure when executing CreateGlobalCredentialsV2", err))
	// 	return diags
	// }
	// if resp1.Response == nil {
	// 	diags = append(diags, diagError(
	// 		"Failure when executing CreateGlobalCredentialsV2", err))
	// 	return diags
	// }
	// taskId := resp1.Response.TaskID
	// log.Printf("[DEBUG] TASKID => %s", taskId)
	// if taskId != "" {
	// 	time.Sleep(5 * time.Second)
	// 	response2, restyResp2, err := client.Task.GetTaskByID(taskId)
	// 	if err != nil || response2 == nil {
	// 		if restyResp2 != nil {
	// 			log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
	// 		}
	// 		diags = append(diags, diagErrorWithAlt(
	// 			"Failure when executing GetTaskByID", err,
	// 			"Failure at GetTaskByID, unexpected response", ""))
	// 		return diags
	// 	}
	// 	if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
	// 		log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
	// 		errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
	// 		err1 := errors.New(errorMsg)
	// 		diags = append(diags, diagError(
	// 			"Failure when executing CreateGlobalCredentialsV2", err1))
	// 		return diags
	// 	}
	// }
	// item3, _, err := client.Discovery.GetAllGlobalCredentialsV2(&queryParamValidate)
	// if err != nil || item3 == nil {
	// 	diags = append(diags, diagErrorWithAlt(
	// 		"Failure when executing CreateGlobalCredentialsV2", err,
	// 		"Failure at CreateGlobalCredentialsV2, unexpected response", ""))
	// 	return diags
	// }

	// resourceMap := make(map[string]string)

	// d.SetId(joinResourceID(resourceMap))
	return resourceGlobalCredentialV2Read(ctx, d, m)
}

func resourceGlobalCredentialV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	// resourceID := d.Id()
	// resourceMap := separateResourceID(resourceID)

	// selectedMethod := 1
	// if selectedMethod == 1 {
	// 	log.Printf("[DEBUG] Selected method: GetAllGlobalCredentialsV2")

	// 	response1, restyResp1, err := client.Discovery.GetAllGlobalCredentialsV2()

	// 	if err != nil || response1 == nil {
	// 		if restyResp1 != nil {
	// 			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
	// 		}
	// 		d.SetId("")
	// 		return diags
	// 	}

	// 	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	// 	vItem1 := flattenDiscoveryGetAllGlobalCredentialsV2Item(response1.Response)
	// 	if err := d.Set("item", vItem1); err != nil {
	// 		diags = append(diags, diagError(
	// 			"Failure when setting GetAllGlobalCredentialsV2 response",
	// 			err))
	// 		return diags
	// 	}

	// 	return diags

	// }
	return diags
}

func resourceGlobalCredentialV2Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Discovery.UpdateGlobalCredentialsV2(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGlobalCredentialsV2", err, restyResp1.String(),
					"Failure at UpdateGlobalCredentialsV2, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalCredentialsV2", err,
				"Failure at UpdateGlobalCredentialsV2, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateGlobalCredentialsV2", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateGlobalCredentialsV2", err1))
				return diags
			}
		}

	}

	return resourceGlobalCredentialV2Read(ctx, d, m)
}

func resourceGlobalCredentialV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	// resourceID := d.Id()
	// resourceMap := separateResourceID(resourceID)

	// response1, restyResp1, err := client.Discovery.DeleteGlobalCredentialV2(vvID)
	// if err != nil || response1 == nil {
	// 	if restyResp1 != nil {
	// 		log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
	// 		diags = append(diags, diagErrorWithAltAndResponse(
	// 			"Failure when executing DeleteGlobalCredentialV2", err, restyResp1.String(),
	// 			"Failure at DeleteGlobalCredentialV2, unexpected response", ""))
	// 		return diags
	// 	}
	// 	diags = append(diags, diagErrorWithAlt(
	// 		"Failure when executing DeleteGlobalCredentialV2", err,
	// 		"Failure at DeleteGlobalCredentialV2, unexpected response", ""))
	// 	return diags
	// }

	// if response1.Response == nil {
	// 	diags = append(diags, diagError(
	// 		"Failure when executing DeleteGlobalCredentialV2", err))
	// 	return diags
	// }
	// taskId := response1.Response.TaskID
	// log.Printf("[DEBUG] TASKID => %s", taskId)
	// if taskId != "" {
	// 	time.Sleep(5 * time.Second)
	// 	response2, restyResp2, err := client.Task.GetTaskByID(taskId)
	// 	if err != nil || response2 == nil {
	// 		if restyResp2 != nil {
	// 			log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
	// 		}
	// 		diags = append(diags, diagErrorWithAlt(
	// 			"Failure when executing GetTaskByID", err,
	// 			"Failure at GetTaskByID, unexpected response", ""))
	// 		return diags
	// 	}
	// 	if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
	// 		log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
	// 		errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
	// 		err1 := errors.New(errorMsg)
	// 		diags = append(diags, diagError(
	// 			"Failure when executing DeleteGlobalCredentialV2", err1))
	// 		return diags
	// 	}
	// }

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2 {
	request := dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_credential")))) {
		request.CliCredential = expandRequestGlobalCredentialV2CreateGlobalCredentialsV2CliCredentialArray(ctx, key+".cli_credential", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_read")))) {
		request.SNMPV2CRead = expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CReadArray(ctx, key+".snmp_v2c_read", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_write")))) {
		request.SNMPV2CWrite = expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CWriteArray(ctx, key+".snmp_v2c_write", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v3")))) {
		request.SNMPV3 = expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV3Array(ctx, key+".snmp_v3", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_read")))) {
		request.HTTPSRead = expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSReadArray(ctx, key+".https_read", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_write")))) {
		request.HTTPSWrite = expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSWriteArray(ctx, key+".https_write", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2CliCredentialArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2CliCredential {
	request := []dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2CliCredential{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestGlobalCredentialV2CreateGlobalCredentialsV2CliCredential(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2CliCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2CliCredential {
	request := dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2CliCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password")))) {
		request.EnablePassword = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CReadArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CRead {
	request := []dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CRead{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CRead(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CRead {
	request := dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".read_community")))) {
		request.ReadCommunity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CWriteArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CWrite {
	request := []dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CWrite{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CWrite(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV2CWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CWrite {
	request := dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV2CWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".write_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".write_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".write_community")))) {
		request.WriteCommunity = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV3Array(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV3 {
	request := []dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV3{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV3(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2SNMPV3(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV3 {
	request := dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2SNMPV3{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_type")))) {
		request.PrivacyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_password")))) {
		request.PrivacyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_password")))) {
		request.AuthPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSReadArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSRead {
	request := []dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSRead{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSRead(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSRead {
	request := dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSWriteArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSWrite {
	request := []dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSWrite{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSWrite(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2CreateGlobalCredentialsV2HTTPSWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSWrite {
	request := dnacentersdkgo.RequestDiscoveryCreateGlobalCredentialsV2HTTPSWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2 {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_credential")))) {
		request.CliCredential = expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2CliCredential(ctx, key+".cli_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_read")))) {
		request.SNMPV2CRead = expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2SNMPV2CRead(ctx, key+".snmp_v2c_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_write")))) {
		request.SNMPV2CWrite = expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2SNMPV2CWrite(ctx, key+".snmp_v2c_write.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v3")))) {
		request.SNMPV3 = expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2SNMPV3(ctx, key+".snmp_v3.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_read")))) {
		request.HTTPSRead = expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2HTTPSRead(ctx, key+".https_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_write")))) {
		request.HTTPSWrite = expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2HTTPSWrite(ctx, key+".https_write.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2CliCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2CliCredential {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2CliCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_password")))) {
		request.EnablePassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2SNMPV2CRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CRead {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".read_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".read_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".read_community")))) {
		request.ReadCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2SNMPV2CWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CWrite {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2SNMPV2CWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".write_community")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".write_community")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".write_community")))) {
		request.WriteCommunity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2SNMPV3(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2SNMPV3 {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2SNMPV3{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_password")))) {
		request.AuthPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_mode")))) {
		request.SNMPMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_password")))) {
		request.PrivacyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".privacy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".privacy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".privacy_type")))) {
		request.PrivacyType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2HTTPSRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2HTTPSRead {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2HTTPSRead{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalCredentialV2UpdateGlobalCredentialsV2HTTPSWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2HTTPSWrite {
	request := dnacentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV2HTTPSWrite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
