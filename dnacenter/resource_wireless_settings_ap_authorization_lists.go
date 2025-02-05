package dnacenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWirelessSettingsApAuthorizationLists() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Wireless.

- This resource allows the user to create an AP Authorization List.
`,

		CreateContext: resourceWirelessSettingsApAuthorizationListsCreate,
		ReadContext:   resourceWirelessSettingsApAuthorizationListsRead,
		UpdateContext: resourceWirelessSettingsApAuthorizationListsUpdate,
		DeleteContext: resourceWirelessSettingsApAuthorizationListsDelete,
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

						"ap_authorization_list_name": &schema.Schema{
							Description: `Ap Authorization List Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"local_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_mac_entries": &schema.Schema{
										Description: `AP Mac Addresses`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ap_serial_number_entries": &schema.Schema{
										Description: `AP Serial Number Entries`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"remote_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_servers": &schema.Schema{
										Description: `AAA Servers`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"authorize_ap_with_mac": &schema.Schema{
										Description: `Authorize AP With Mac`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"authorize_ap_with_serial_number": &schema.Schema{
										Description: `Authorize AP With Serial Number`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
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

						"ap_authorization_list_name": &schema.Schema{
							Description: `AP Authorization List Name. For a AP Authorization List to be created successfully, either Local Authorization or Remote Authorization is mandatory.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"local_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_mac_entries": &schema.Schema{
										Description: `List of Access Point's Ethernet MAC addresses. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"ap_serial_number_entries": &schema.Schema{
										Description: `List of Access Point's Serial Numbers.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"remote_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_servers": &schema.Schema{
										Description: `List of Authorization server IpAddresses. Obtain the AAA servers by using the API GET call '/dna/intent/api/v1/authentication-policy-servers'.
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"authorize_ap_with_mac": &schema.Schema{
										Description: `True if AP Authorization List should authorise APs With MAC addresses, else False. (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"authorize_ap_with_serial_number": &schema.Schema{
										Description: `True if server IpAddresses are added and AP Authorization List should authorise APs With Serial Numbers, else False (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)
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

func resourceWirelessSettingsApAuthorizationListsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessSettingsApAuthorizationListsCreateApAuthorizationList(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vApAuthorizationListName := resourceItem["ap_authorization_list_name"]
	vvApAuthorizationListName := interfaceToString(vApAuthorizationListName)

	queryParamImport := dnacentersdkgo.GetApAuthorizationListsQueryParams{}
	queryParamImport.ApAuthorizationListName = vvApAuthorizationListName
	item2, _, err := client.Wireless.GetApAuthorizationLists(&queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["ap_authorization_list_name"] = item2.Response.ApAuthorizationListName
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessSettingsApAuthorizationListsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Wireless.CreateApAuthorizationList(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApAuthorizationList", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApAuthorizationList", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApAuthorizationList", err))
		return diags
	}
	taskId := resp1.Response.TaskID
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
				"Failure when executing CreateApAuthorizationList", err1))
			return diags
		}
	}
	queryParamValidate := dnacentersdkgo.GetApAuthorizationListsQueryParams{}
	queryParamValidate.ApAuthorizationListName = vvApAuthorizationListName
	item3, _, err := client.Wireless.GetApAuthorizationLists(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateApAuthorizationList", err,
			"Failure at CreateApAuthorizationList, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["ap_authorization_list_name"] = item3.Response.ApAuthorizationListName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessSettingsApAuthorizationListsRead(ctx, d, m)
}

func resourceWirelessSettingsApAuthorizationListsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvApAuthorizationListName := resourceMap["ap_authorization_list_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApAuthorizationLists")
		queryParams1 := dnacentersdkgo.GetApAuthorizationListsQueryParams{}
		queryParams1.ApAuthorizationListName = vvApAuthorizationListName
		response1, restyResp1, err := client.Wireless.GetApAuthorizationLists(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetApAuthorizationListsItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApAuthorizationLists response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceWirelessSettingsApAuthorizationListsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceWirelessSettingsApAuthorizationListsRead(ctx, d, m)
}

func resourceWirelessSettingsApAuthorizationListsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete WirelessSettingsApAuthorizationLists on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestWirelessSettingsApAuthorizationListsCreateApAuthorizationList(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApAuthorizationList {
	request := dnacentersdkgo.RequestWirelessCreateApAuthorizationList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_authorization_list_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_authorization_list_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_authorization_list_name")))) {
		request.ApAuthorizationListName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".local_authorization")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".local_authorization")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".local_authorization")))) {
		request.LocalAuthorization = expandRequestWirelessSettingsApAuthorizationListsCreateApAuthorizationListLocalAuthorization(ctx, key+".local_authorization.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".remote_authorization")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".remote_authorization")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".remote_authorization")))) {
		request.RemoteAuthorization = expandRequestWirelessSettingsApAuthorizationListsCreateApAuthorizationListRemoteAuthorization(ctx, key+".remote_authorization.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsApAuthorizationListsCreateApAuthorizationListLocalAuthorization(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApAuthorizationListLocalAuthorization {
	request := dnacentersdkgo.RequestWirelessCreateApAuthorizationListLocalAuthorization{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_mac_entries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_mac_entries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_mac_entries")))) {
		request.ApMacEntries = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_serial_number_entries")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_serial_number_entries")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_serial_number_entries")))) {
		request.ApSerialNumberEntries = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsApAuthorizationListsCreateApAuthorizationListRemoteAuthorization(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateApAuthorizationListRemoteAuthorization {
	request := dnacentersdkgo.RequestWirelessCreateApAuthorizationListRemoteAuthorization{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aaa_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aaa_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aaa_servers")))) {
		request.AAAServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authorize_ap_with_mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authorize_ap_with_mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authorize_ap_with_mac")))) {
		request.AuthorizeApWithMac = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authorize_ap_with_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authorize_ap_with_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authorize_ap_with_serial_number")))) {
		request.AuthorizeApWithSerialNumber = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
