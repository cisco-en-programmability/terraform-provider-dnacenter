package dnacenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	//dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
	dnacentersdkgo "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessControllersWirelessMobilityGroupsMobilityProvision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action is used to provision/deploy wireless mobility into Cisco wireless controllers.
`,

		CreateContext: resourceWirelessControllersWirelessMobilityGroupsMobilityProvisionCreate,
		ReadContext:   resourceWirelessControllersWirelessMobilityGroupsMobilityProvisionRead,
		DeleteContext: resourceWirelessControllersWirelessMobilityGroupsMobilityProvisionDelete,
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

						"task_id": &schema.Schema{
							Description: `Asynchronous Task Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Asynchronous Task URL for further tracking
`,
							Type:     schema.TypeString,
							Computed: true,
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
						"data_link_encryption": &schema.Schema{
							Description: `A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"dtls_high_cipher": &schema.Schema{
							Description: `DTLS High Cipher.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"mac_address": &schema.Schema{
							Description: `Device mobility MAC Address. Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"management_ip": &schema.Schema{
							Description: `Self device wireless Management IP.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"mobility_group_name": &schema.Schema{
							Description: `Self device Group Name. Must be alphanumeric without {!,<,space,?/'} <br/> and maximum of 31 characters.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"mobility_peers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_series": &schema.Schema{
										Description: `Indicates peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"hash_key": &schema.Schema{
										Description: `SSC hash string must be 40 characters.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"member_mac_address": &schema.Schema{
										Description: `Peer device mobility MAC Address.  Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mobility_group_name": &schema.Schema{
										Description: `Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} <br/> and maximum of 31 characters.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"peer_device_name": &schema.Schema{
										Description: `Peer device Host Name.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"peer_ip": &schema.Schema{
										Description: `This indicates public ip address.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"peer_network_device_id": &schema.Schema{
										Description: `The possible values are: UNKNOWN or valid UUID of Network device Id. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device id.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"private_ip_address": &schema.Schema{
										Description: `This indicates private/management ip address.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"network_device_id": &schema.Schema{
							Description: `Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllersWirelessMobilityGroupsMobilityProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestWirelessControllersWirelessMobilityGroupsMobilityProvisionMobilityProvision(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.MobilityProvision(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing MobilityProvision", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing MobilityProvision", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing MobilityProvision", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessMobilityProvisionItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting MobilityProvision response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessControllersWirelessMobilityGroupsMobilityProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessControllersWirelessMobilityGroupsMobilityProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessControllersWirelessMobilityGroupsMobilityProvisionMobilityProvision(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessMobilityProvision {
	request := dnacentersdkgo.RequestWirelessMobilityProvision{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobility_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobility_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobility_group_name")))) {
		request.MobilityGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_ip")))) {
		request.ManagementIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dtls_high_cipher")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dtls_high_cipher")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dtls_high_cipher")))) {
		request.DtlsHighCipher = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_link_encryption")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_link_encryption")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_link_encryption")))) {
		request.DataLinkEncryption = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobility_peers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobility_peers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobility_peers")))) {
		request.MobilityPeers = expandRequestWirelessControllersWirelessMobilityGroupsMobilityProvisionMobilityProvisionMobilityPeersArray(ctx, key+".mobility_peers", d)
	}
	return &request
}

func expandRequestWirelessControllersWirelessMobilityGroupsMobilityProvisionMobilityProvisionMobilityPeersArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessMobilityProvisionMobilityPeers {
	request := []dnacentersdkgo.RequestWirelessMobilityProvisionMobilityPeers{}
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
		i := expandRequestWirelessControllersWirelessMobilityGroupsMobilityProvisionMobilityProvisionMobilityPeers(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessControllersWirelessMobilityGroupsMobilityProvisionMobilityProvisionMobilityPeers(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessMobilityProvisionMobilityPeers {
	request := dnacentersdkgo.RequestWirelessMobilityProvisionMobilityPeers{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".peer_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peer_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peer_ip")))) {
		request.PeerIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".private_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".private_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".private_ip_address")))) {
		request.PrivateIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".peer_device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peer_device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peer_device_name")))) {
		request.PeerDeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".peer_network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peer_network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peer_network_device_id")))) {
		request.PeerNetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobility_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobility_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobility_group_name")))) {
		request.MobilityGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_mac_address")))) {
		request.MemberMacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_series")))) {
		request.DeviceSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hash_key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hash_key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hash_key")))) {
		request.HashKey = interfaceToString(v)
	}
	return &request
}

func flattenWirelessMobilityProvisionItem(item *dnacentersdkgo.ResponseWirelessMobilityProvisionResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
