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
func dataSourceDeviceCredentialUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Network Settings.

- API to update device credentials.
`,

		ReadContext: dataSourceDeviceCredentialUpdateRead,
		Schema: map[string]*schema.Schema{
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
			"settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cli_credential": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"enable_password": &schema.Schema{
										Description: `Enable Password`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
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
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
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
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
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
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"read_community": &schema.Schema{
										Description: `Read Community`,
										Type:        schema.TypeString,
										Optional:    true,
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
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"write_community": &schema.Schema{
										Description: `Write Community`,
										Type:        schema.TypeString,
										Optional:    true,
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
										Description: `Auth Password`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"privacy_password": &schema.Schema{
										Description: `Privacy Password`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"privacy_type": &schema.Schema{
										Description: `Privacy Type`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"snmp_mode": &schema.Schema{
										Description: `Snmp Mode`,
										Type:        schema.TypeString,
										Optional:    true,
									},
									"username": &schema.Schema{
										Description: `Username`,
										Type:        schema.TypeString,
										Optional:    true,
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

func dataSourceDeviceCredentialUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: UpdateDeviceCredentials")
		request1 := expandRequestDeviceCredentialUpdateUpdateDeviceCredentials(ctx, "", d)

		response1, restyResp1, err := client.NetworkSettings.UpdateDeviceCredentials(request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateDeviceCredentials", err,
				"Failure at UpdateDeviceCredentials, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsUpdateDeviceCredentialsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting UpdateDeviceCredentials response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentials(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentials {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentials{}
	request.Settings = expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettings(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettings {
	request := dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cli_credential")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cli_credential")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cli_credential")))) {
		request.CliCredential = expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsCliCredential(ctx, key+".cli_credential.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_read")))) {
		request.SNMPV2CRead = expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsSNMPV2CRead(ctx, key+".snmp_v2c_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v2c_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v2c_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v2c_write")))) {
		request.SNMPV2CWrite = expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsSNMPV2CWrite(ctx, key+".snmp_v2c_write.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".snmp_v3")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".snmp_v3")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".snmp_v3")))) {
		request.SNMPV3 = expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsSNMPV3(ctx, key+".snmp_v3.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_read")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_read")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_read")))) {
		request.HTTPSRead = expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsHTTPSRead(ctx, key+".https_read.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".https_write")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".https_write")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".https_write")))) {
		request.HTTPSWrite = expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsHTTPSWrite(ctx, key+".https_write.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsCliCredential(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsCliCredential {
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

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsSNMPV2CRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CRead {
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

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsSNMPV2CWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CWrite {
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

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsSNMPV3(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV3 {
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

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsHTTPSRead(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSRead {
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

func expandRequestDeviceCredentialUpdateUpdateDeviceCredentialsSettingsHTTPSWrite(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSWrite {
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

func flattenNetworkSettingsUpdateDeviceCredentialsItem(item *dnacentersdkgo.ResponseNetworkSettingsUpdateDeviceCredentials) []map[string]interface{} {
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
