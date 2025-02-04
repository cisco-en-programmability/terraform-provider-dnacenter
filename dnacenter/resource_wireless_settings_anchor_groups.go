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

func resourceWirelessSettingsAnchorGroups() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Wireless.

- This resource allows the user to create an AnchorGroup
`,

		CreateContext: resourceWirelessSettingsAnchorGroupsCreate,
		ReadContext:   resourceWirelessSettingsAnchorGroupsRead,
		UpdateContext: resourceWirelessSettingsAnchorGroupsUpdate,
		DeleteContext: resourceWirelessSettingsAnchorGroupsDelete,
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

						"anchor_group_name": &schema.Schema{
							Description: `Anchor Group Name. Max length is 32 characters
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Anchor Profile unique ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"mobility_anchors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"anchor_priority": &schema.Schema{
										Description: `This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_name": &schema.Schema{
										Description: `Peer Host Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `This indicates Mobility public IP address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"managed_anchor_wlc": &schema.Schema{
										Description: `This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mobility_group_name": &schema.Schema{
										Description: `Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"peer_device_type": &schema.Schema{
										Description: `Indicates peer device mobility belongs to AireOS or IOS-XE family.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_ip": &schema.Schema{
										Description: `This indicates private management IP address
`,
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

						"anchor_group_name": &schema.Schema{
							Description: `Anchor Group Name. Max length is 32 characters
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mobility_anchors": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"anchor_priority": &schema.Schema{
										Description: `This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"device_name": &schema.Schema{
										Description: `Peer Host Name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `This indicates Mobility public IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"managed_anchor_wlc": &schema.Schema{
										Description: `This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"mobility_group_name": &schema.Schema{
										Description: `Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"peer_device_type": &schema.Schema{
										Description: `Indicates peer device mobility belongs to AireOS or IOS-XE family.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"private_ip": &schema.Schema{
										Description: `This indicates private management IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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

func resourceWirelessSettingsAnchorGroupsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	request1 := expandRequestWirelessSettingsAnchorGroupsCreateAnchorGroup(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	item2, _, err := client.Wireless.GetAnchorGroups()
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessSettingsAnchorGroupsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Wireless.CreateAnchorGroup(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateAnchorGroup", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateAnchorGroup", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateAnchorGroup", err))
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
				"Failure when executing CreateAnchorGroup", err1))
			return diags
		}
	}
	item3, _, err := client.Wireless.GetAnchorGroups()
	if err != nil || item3 != nil {
		resourceMap := make(map[string]string)
		d.SetId(joinResourceID(resourceMap))
		return resourceWirelessSettingsAnchorGroupsRead(ctx, d, m)
	}
	item4, _, err := client.Wireless.GetAnchorGroups()
	if err != nil || item4 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateAnchorGroup", err,
			"Failure at CreateAnchorGroup, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)

	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessSettingsAnchorGroupsRead(ctx, d, m)
}

func resourceWirelessSettingsAnchorGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAnchorGroups")

		response1, restyResp1, err := client.Wireless.GetAnchorGroups()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAnchorGroupsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAnchorGroups response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceWirelessSettingsAnchorGroupsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceWirelessSettingsAnchorGroupsRead(ctx, d, m)
}

func resourceWirelessSettingsAnchorGroupsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete WirelessSettingsAnchorGroups on Dna Center
	//       Returning empty diags to delete it on Terraform
	return diags
}
func expandRequestWirelessSettingsAnchorGroupsCreateAnchorGroup(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateAnchorGroup {
	request := dnacentersdkgo.RequestWirelessCreateAnchorGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".anchor_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".anchor_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".anchor_group_name")))) {
		request.AnchorGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobility_anchors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobility_anchors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobility_anchors")))) {
		request.MobilityAnchors = expandRequestWirelessSettingsAnchorGroupsCreateAnchorGroupMobilityAnchorsArray(ctx, key+".mobility_anchors", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsAnchorGroupsCreateAnchorGroupMobilityAnchorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestWirelessCreateAnchorGroupMobilityAnchors {
	request := []dnacentersdkgo.RequestWirelessCreateAnchorGroupMobilityAnchors{}
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
		i := expandRequestWirelessSettingsAnchorGroupsCreateAnchorGroupMobilityAnchors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestWirelessSettingsAnchorGroupsCreateAnchorGroupMobilityAnchors(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateAnchorGroupMobilityAnchors {
	request := dnacentersdkgo.RequestWirelessCreateAnchorGroupMobilityAnchors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_name")))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".anchor_priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".anchor_priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".anchor_priority")))) {
		request.AnchorPriority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".managed_anchor_wlc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".managed_anchor_wlc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".managed_anchor_wlc")))) {
		request.ManagedAnchorWlc = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".peer_device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".peer_device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".peer_device_type")))) {
		request.PeerDeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mobility_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mobility_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mobility_group_name")))) {
		request.MobilityGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".private_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".private_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".private_ip")))) {
		request.PrivateIP = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
