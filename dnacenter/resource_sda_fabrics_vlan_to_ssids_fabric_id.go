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

func resourceSdaFabricsVLANToSSIDsFabricID() *schema.Resource {
	return &schema.Resource{
		Description: `It manages read and update operations on Fabric Wireless.

- Add, update, or remove SSID mappings to a VLAN. If the payload doesn't contain a 'vlanName' which has SSIDs mapping
done earlier then all the mapped SSIDs of the 'vlanName' is cleared. The request must include all SSIDs currently mapped
to a VLAN, as determined by the response from the GET operation for the same fabricId used in the request. If an
already-mapped SSID is not included in the payload, its mapping will be removed by this API. Conversely, if a new SSID
is provided, it will be added to the Mapping. Ensure that any new SSID added is a Fabric SSID. This resource can also be
used to add a VLAN and associate the relevant SSIDs with it. The 'vlanName' must be 'Fabric Wireless Enabled' and should
be part of the Fabric Site representing 'Fabric ID' specified in the API request.
`,

		CreateContext: resourceSdaFabricsVLANToSSIDsFabricIDCreate,
		ReadContext:   resourceSdaFabricsVLANToSSIDsFabricIDRead,
		UpdateContext: resourceSdaFabricsVLANToSSIDsFabricIDUpdate,
		DeleteContext: resourceSdaFabricsVLANToSSIDsFabricIDDelete,
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

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Name of the SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_group_tag": &schema.Schema{
										Description: `Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"vlan_name": &schema.Schema{
							Description: `Vlan Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN`,
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

									"fabric_id": &schema.Schema{
										Description: `fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site
`,
										Type:     schema.TypeString,
										Required: true,
									},
									"ssid_details": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name of the SSID
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"security_group_tag": &schema.Schema{
													Description: `Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"vlan_name": &schema.Schema{
										Description: `Vlan Name`,
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

func resourceSdaFabricsVLANToSSIDsFabricIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	resourceMap := make(map[string]string)
	// TODO: Add the path params to `item` schema
	//       & return it individually
	resourceMap["id"] = interfaceToString(resourceItem["fabric_id"])
	resourceMap["name"] = interfaceToString(resourceItem["vlan_name"])
	d.SetId(joinResourceID(resourceMap))
	return diags
}

func resourceSdaFabricsVLANToSSIDsFabricIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vName := resourceMap["name"]
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite")

		queryParams1 := dnacentersdkgo.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams{}
		item1, err := searchFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(m, queryParams1, vID, vName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}

		items := []dnacentersdkgo.ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponse{
			*item1,
		}

		// Review flatten function used
		vItem1 := flattenFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaFabricsVLANToSSIDsFabricIDUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvFabricID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLAN(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.FabricWireless.AddUpdateOrRemoveSSIDMappingToAVLAN(vvFabricID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing AddUpdateOrRemoveSSIDMappingToAVLAN", err, restyResp1.String(),
					"Failure at AddUpdateOrRemoveSSIDMappingToAVLAN, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing AddUpdateOrRemoveSSIDMappingToAVLAN", err,
				"Failure at AddUpdateOrRemoveSSIDMappingToAVLAN, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing AddUpdateOrRemoveSSIDMappingToAVLAN", err))
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
					"Failure when executing AddUpdateOrRemoveSSIDMappingToAVLAN", err1))
				return diags
			}
		}

	}

	return resourceSdaFabricsVLANToSSIDsFabricIDRead(ctx, d, m)
}

func resourceSdaFabricsVLANToSSIDsFabricIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SdaFabricsVLANToSSIDsFabricID", err, "Delete method is not supported",
		"Failure at SdaFabricsVLANToSSIDsFabricIDDelete, unexpected response", ""))
	return diags
}
func expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLAN(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN {
	request := dnacentersdkgo.RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN{}
	if v := expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN {
	request := []dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN{}
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
		i := expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN {
	request := dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid_details")))) {
		request.SSIDDetails = expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItemSSIDDetailsArray(ctx, key+".ssid_details", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItemSSIDDetailsArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANSSIDDetails {
	request := []dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANSSIDDetails{}
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
		i := expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItemSSIDDetails(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricsVLANToSSIDsFabricIDAddUpdateOrRemoveSSIDMappingToAVLANItemSSIDDetails(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANSSIDDetails {
	request := dnacentersdkgo.RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANSSIDDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_group_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_group_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_group_tag")))) {
		request.SecurityGroupTag = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(m interface{}, queryParams dnacentersdkgo.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams, vID string, vName string) (*dnacentersdkgo.ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponse, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponse
	// var ite *dnacentersdkgo.ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.FabricWireless.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(vID, nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vName == item.VLANName {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.FabricWireless.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(vID, &queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	}
	return foundItem, err
}
