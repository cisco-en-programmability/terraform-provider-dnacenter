package dnacenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaAuthenticationProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on SDA.

- Updates an authentication profile based on user input.
`,

		CreateContext: resourceSdaAuthenticationProfilesCreate,
		ReadContext:   resourceSdaAuthenticationProfilesRead,
		UpdateContext: resourceSdaAuthenticationProfilesUpdate,
		DeleteContext: resourceSdaAuthenticationProfilesDelete,
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

						"authentication_order": &schema.Schema{
							Description: `First authentication method.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication_profile_name": &schema.Schema{
							Description: `The default host authentication template.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"dot1x_to_mab_fallback_timeout": &schema.Schema{
							Description: `802.1x Timeout.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"fabric_id": &schema.Schema{
							Description: `ID of the fabric this authentication profile is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the authentication profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_bpdu_guard_enabled": &schema.Schema{
							Description: `Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to "Closed Authentication".
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"number_of_hosts": &schema.Schema{
							Description: `Number of Hosts.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"pre_auth_acl": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"access_contracts": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"action": &schema.Schema{
													Description: `Contract behaviour.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"port": &schema.Schema{
													Description: `Port for the access contract.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol for the access contract.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"description": &schema.Schema{
										Description: `Description of this Pre-Authentication ACL.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"enabled": &schema.Schema{
										Description: `Enable/disable Pre-Authentication ACL.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"implicit_action": &schema.Schema{
										Description: `Implicit behaviour unless overridden.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"wake_on_lan": &schema.Schema{
							Description: `Wake on LAN.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaUpdateAuthenticationProfile`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"authentication_order": &schema.Schema{
										Description: `First authentication method.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"authentication_profile_name": &schema.Schema{
										Description: `The default host authentication template (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dot1x_to_mab_fallback_timeout": &schema.Schema{
										Description: `802.1x Timeout.
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric this authentication profile is assigned to (updating this field is not allowed). To update a global authentication profile, either remove this property or set its value to null.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the authentication profile (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_bpdu_guard_enabled": &schema.Schema{
										Description: `Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to "Closed Authentication" (defaults to true).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"number_of_hosts": &schema.Schema{
										Description: `Number of Hosts.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pre_auth_acl": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"access_contracts": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"action": &schema.Schema{
																Description: `Contract behaviour.
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Port for the access contract. The port can only be used once in the Access Contract list.
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol for the access contract. "TCP" and "TCP_UDP" are only allowed when the contract port is "domain".
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"description": &schema.Schema{
													Description: `Description of this Pre-Authentication ACL.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"enabled": &schema.Schema{
													Description: `Enable/disable Pre-Authentication ACL.
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"implicit_action": &schema.Schema{
													Description: `Implicit behaviour unless overridden (defaults to "DENY").
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"wake_on_lan": &schema.Schema{
										Description: `Wake on LAN.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
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

func resourceSdaAuthenticationProfilesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["id"])
	resourceMap["name"] = interfaceToString(resourceItem["authentication_profile_name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSdaAuthenticationProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	vvName := resourceMap["name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAuthenticationProfiles")
		queryParams1 := dnacentersdkgo.GetAuthenticationProfilesQueryParams{}

		queryParams1.AuthenticationProfileName = vvName

		item1, err := searchSdaGetAuthenticationProfiles(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used

		items := []dnacentersdkgo.ResponseSdaGetAuthenticationProfilesResponse{
			*item1,
		}
		vItem1 := flattenSdaGetAuthenticationProfilesItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAuthenticationProfiles search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaAuthenticationProfilesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	if d.HasChange("parameters") {
		request1 := expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfile(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateAuthenticationProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateAuthenticationProfile", err, restyResp1.String(),
					"Failure at UpdateAuthenticationProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateAuthenticationProfile", err,
				"Failure at UpdateAuthenticationProfile, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateAuthenticationProfile", err))
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
					"Failure when executing UpdateAuthenticationProfile", err1))
				return diags
			}
		}

	}

	return resourceSdaAuthenticationProfilesRead(ctx, d, m)
}

func resourceSdaAuthenticationProfilesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SdaAuthenticationProfiles", err, "Delete method is not supported",
		"Failure at SdaAuthenticationProfilesDelete, unexpected response", ""))
	return diags
}
func expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfile(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSdaUpdateAuthenticationProfile {
	request := dnacentersdkgo.RequestSdaUpdateAuthenticationProfile{}
	if v := expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfile {
	request := []dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfile{}
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
		i := expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfile {
	request := dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_profile_name")))) {
		request.AuthenticationProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_order")))) {
		request.AuthenticationOrder = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dot1x_to_mab_fallback_timeout")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dot1x_to_mab_fallback_timeout")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dot1x_to_mab_fallback_timeout")))) {
		request.Dot1XToMabFallbackTimeout = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wake_on_lan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wake_on_lan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wake_on_lan")))) {
		request.WakeOnLan = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_hosts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_hosts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_hosts")))) {
		request.NumberOfHosts = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_bpdu_guard_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_bpdu_guard_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_bpdu_guard_enabled")))) {
		request.IsBpduGuardEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pre_auth_acl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pre_auth_acl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pre_auth_acl")))) {
		request.PreAuthACL = expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemPreAuthACL(ctx, key+".pre_auth_acl.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemPreAuthACL(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfilePreAuthACL {
	request := dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfilePreAuthACL{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enabled")))) {
		request.Enabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".implicit_action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".implicit_action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".implicit_action")))) {
		request.ImplicitAction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".access_contracts")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".access_contracts")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".access_contracts")))) {
		request.AccessContracts = expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemPreAuthACLAccessContractsArray(ctx, key+".access_contracts", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemPreAuthACLAccessContractsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfilePreAuthACLAccessContracts {
	request := []dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfilePreAuthACLAccessContracts{}
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
		i := expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemPreAuthACLAccessContracts(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaAuthenticationProfilesUpdateAuthenticationProfileItemPreAuthACLAccessContracts(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfilePreAuthACLAccessContracts {
	request := dnacentersdkgo.RequestItemSdaUpdateAuthenticationProfilePreAuthACLAccessContracts{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".action")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".action")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".action")))) {
		request.Action = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetAuthenticationProfiles(m interface{}, queryParams dnacentersdkgo.GetAuthenticationProfilesQueryParams, vID string) (*dnacentersdkgo.ResponseSdaGetAuthenticationProfilesResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseSdaGetAuthenticationProfilesResponse
	var ite *dnacentersdkgo.ResponseSdaGetAuthenticationProfiles
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetAuthenticationProfiles(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Sda.GetAuthenticationProfiles(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.AuthenticationProfileName != "" {
		ite, _, err = client.Sda.GetAuthenticationProfiles(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.AuthenticationProfileName == queryParams.AuthenticationProfileName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
