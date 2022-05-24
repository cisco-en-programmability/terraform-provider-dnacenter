package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceCredential() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to create device credentials.

- API to update device credentials.

- Delete device credential.
`,

		CreateContext: resourceDeviceCredentialCreate,
		ReadContext:   resourceDeviceCredentialRead,
		UpdateContext: resourceDeviceCredentialUpdate,
		DeleteContext: resourceDeviceCredentialDelete,
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
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. global credential id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"cli_credential": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Description: `Name or description for CLI credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"enable_password": &schema.Schema{
													Description: `Enable password for CLI credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"password": &schema.Schema{
													Description: `Password for CLI credential
`,
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"username": &schema.Schema{
													Description: `User name for CLI credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"https_read": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"name": &schema.Schema{
													Description: `Name or description of http read credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"password": &schema.Schema{
													Description: `Password for http read credential
`,
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"port": &schema.Schema{
													Description: `Port for http read credential

ERROR: Different types for param port schema.TypeFloat schema.TypeString`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"username": &schema.Schema{
													Description: `User name of the http read credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"https_write": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"name": &schema.Schema{
													Description: `Name or description of http write credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"password": &schema.Schema{
													Description: `Password for http write credential
`,
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"port": &schema.Schema{
													Description: `Port for http write credential

ERROR: Different types for param port schema.TypeFloat schema.TypeString`,
													Type:     schema.TypeFloat,
													Optional: true,
												},
												"username": &schema.Schema{
													Description: `User name of the http write credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"snmp_v2c_read": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Description: `Description for snmp v2 read
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"read_community": &schema.Schema{
													Description: `Ready community for snmp v2 read credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"snmp_v2c_write": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Description: `Description for snmp v2 write
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"write_community": &schema.Schema{
													Description: `Write community for snmp v2 write credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"snmp_v3": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"auth_password": &schema.Schema{
													Description: `Authentication password for snmpv3 credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"auth_type": &schema.Schema{
													Description: `Authentication type for snmpv3 credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"description": &schema.Schema{
													Description: `Name or description for SNMPV3 credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"privacy_password": &schema.Schema{
													Description: `Privacy password for snmpv3 credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"privacy_type": &schema.Schema{
													Description: `Privacy type for snmpv3 credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"snmp_mode": &schema.Schema{
													Description: `Mode for snmpv3 credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
												"username": &schema.Schema{
													Description: `User name for SNMPv3 credential
`,
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
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

func resourceDeviceCredentialCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestDeviceCredentialCreateDeviceCredentials(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	resp1, restyResp1, err := client.NetworkSettings.CreateDeviceCredentials(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateDeviceCredentials", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateDeviceCredentials", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceDeviceCredentialRead(ctx, d, m)
}

func resourceDeviceCredentialRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetDeviceCredentialDetails")
		queryParams1 := dnacentersdkgo.GetDeviceCredentialDetailsQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID
		}

		response1, restyResp1, err := client.NetworkSettings.GetDeviceCredentialDetails(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDeviceCredentialDetails", err,
				"Failure at GetDeviceCredentialDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetDeviceCredentialDetailsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCredentialDetails response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceDeviceCredentialUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]

	queryParams1 := dnacentersdkgo.GetDeviceCredentialDetailsQueryParams
	queryParams1.SiteID = vSiteID
	item, err := searchNetworkSettingsGetDeviceCredentialDetails(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDeviceCredentialDetails", err,
			"Failure at GetDeviceCredentialDetails, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestDeviceCredentialUpdateDeviceCredentials(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.NetworkSettings.UpdateDeviceCredentials(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateDeviceCredentials", err, restyResp1.String(),
					"Failure at UpdateDeviceCredentials, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceCredentials", err,
				"Failure at UpdateDeviceCredentials, unexpected response", ""))
			return diags
		}
	}

	return resourceDeviceCredentialRead(ctx, d, m)
}

func resourceDeviceCredentialDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]

	queryParams1 := dnacentersdkgo.GetDeviceCredentialDetailsQueryParams
	queryParams1.SiteID = vSiteID
	item, err := searchNetworkSettingsGetDeviceCredentialDetails(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDeviceCredentialDetails", err,
			"Failure at GetDeviceCredentialDetails, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.NetworkSettings.GetDeviceCredentialDetails(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkSettingsGetDeviceCredentialDetails(m, getResp1, nil)
		item1, err := searchNetworkSettingsGetDeviceCredentialDetails(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vID != item1.ID {
			vvID = item1.ID
		} else {
			vvID = vID
		}
	}
	response1, restyResp1, err := client.NetworkSettings.DeleteDeviceCredential(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDeviceCredential", err, restyResp1.String(),
				"Failure at DeleteDeviceCredential, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDeviceCredential", err,
			"Failure at DeleteDeviceCredential, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestDeviceCredentialCreateDeviceCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentials {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentials{}
	request.Settings = expandRequestDeviceCredentialCreateDeviceCredentialsSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettings {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_credential")))) {
		request.CliCredential = expandRequestDeviceCredentialCreateDeviceCredentialsSettingsCliCredentialArray(ctx, key+".cli_credential", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_read")))) {
		request.SNMPV2CRead = expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CReadArray(ctx, key+".snmp_v2c_read", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_write")))) {
		request.SNMPV2CWrite = expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CWriteArray(ctx, key+".snmp_v2c_write", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v3")))) {
		request.SNMPV3 = expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV3Array(ctx, key+".snmp_v3", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_read")))) {
		request.HTTPSRead = expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSReadArray(ctx, key+".https_read", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_write")))) {
		request.HTTPSWrite = expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSWriteArray(ctx, key+".https_write", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsCliCredentialArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceCredentialCreateDeviceCredentialsSettingsCliCredential(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsCliCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential{}
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

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CReadArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CRead(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead{}
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

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CWriteArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CWrite(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV2CWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite{}
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

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV3Array(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3 {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV3(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsSNMPV3(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3 {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3{}
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

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSReadArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSRead(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead{}
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
		request.Port = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSWriteArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite {
	request := []dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSWrite(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialCreateDeviceCredentialsSettingsHTTPSWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite {
	request := dnacentersdkgo.RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite{}
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
		request.Port = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialUpdateDeviceCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentials {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentials{}
	request.Settings = expandRequestDeviceCredentialUpdateDeviceCredentialsSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialUpdateDeviceCredentialsSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_credential")))) {
		request.CliCredential = expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsCliCredential(ctx, key+".cli_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_read")))) {
		request.SNMPV2CRead = expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsSNMPV2CRead(ctx, key+".snmp_v2c_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_write")))) {
		request.SNMPV2CWrite = expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsSNMPV2CWrite(ctx, key+".snmp_v2c_write.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v3")))) {
		request.SNMPV3 = expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsSNMPV3(ctx, key+".snmp_v3.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_read")))) {
		request.HTTPSRead = expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsHTTPSRead(ctx, key+".https_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_write")))) {
		request.HTTPSWrite = expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsHTTPSWrite(ctx, key+".https_write.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsCliCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsCliCredential {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsCliCredential{}
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

func expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsSNMPV2CRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CRead {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CRead{}
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

func expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsSNMPV2CWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CWrite {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CWrite{}
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

func expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsSNMPV3(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV3 {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV3{}
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

func expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsHTTPSRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSRead {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSRead{}
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
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialUpdateDeviceCredentialsSettingsHTTPSWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSWrite {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSWrite{}
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
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
