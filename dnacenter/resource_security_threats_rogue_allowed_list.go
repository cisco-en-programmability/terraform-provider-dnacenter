package dnacenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecurityThreatsRogueAllowedList() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Devices.

- Intent API to add the threat mac address to allowed list.

- Intent API to remove the threat mac address from allowed list.
`,

		CreateContext: resourceSecurityThreatsRogueAllowedListCreate,
		ReadContext:   resourceSecurityThreatsRogueAllowedListRead,
		UpdateContext: resourceSecurityThreatsRogueAllowedListUpdate,
		DeleteContext: resourceSecurityThreatsRogueAllowedListDelete,
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

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"last_modified": &schema.Schema{
							Description: `Last Modified`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestDevicesAddAllowedMacAddress`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"mac_address": &schema.Schema{
							Description: `Mac Address`,
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

func resourceSecurityThreatsRogueAllowedListCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSecurityThreatsRogueAllowedListAddAllowedMacAddress(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vMacAddress := resourceItem["mac_address"]
	vvMacAddress := interfaceToString(vMacAddress)
	queryParamImport := dnacentersdkgo.GetAllowedMacAddressQueryParams{}
	item2, err := searchDevicesGetAllowedMacAddress(m, queryParamImport, vvMacAddress)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["mac_address"] = item2.MacAddress
		d.SetId(joinResourceID(resourceMap))
		return resourceSecurityThreatsRogueAllowedListRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Devices.AddAllowedMacAddress(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddAllowedMacAddress", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddAllowedMacAddress", err))
		return diags
	}
	// TODO REVIEW
	queryParamValidate := dnacentersdkgo.GetAllowedMacAddressQueryParams{}
	item3, err := searchDevicesGetAllowedMacAddress(m, queryParamValidate, vvMacAddress)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddAllowedMacAddress", err,
			"Failure at AddAllowedMacAddress, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["mac_address"] = item3.MacAddress
	d.SetId(joinResourceID(resourceMap))
	return resourceSecurityThreatsRogueAllowedListRead(ctx, d, m)
}

func resourceSecurityThreatsRogueAllowedListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvMacAddress := resourceMap["mac_address"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllowedMacAddress")
		queryParams1 := dnacentersdkgo.GetAllowedMacAddressQueryParams{}

		// has_unknown_response: None

		item1, err := searchDevicesGetAllowedMacAddress(m, queryParams1, vvMacAddress)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		var items dnacentersdkgo.ResponseDevicesGetAllowedMacAddress
		items = append(items, *item1)
		// Review flatten function used
		vItem1 := flattenDevicesGetAllowedMacAddressItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedMacAddress search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSecurityThreatsRogueAllowedListUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSecurityThreatsRogueAllowedListRead(ctx, d, m)
}

func resourceSecurityThreatsRogueAllowedListDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvMacAddress := resourceMap["mac_address"]

	response1, restyResp1, err := client.Devices.RemoveAllowedMacAddress(vvMacAddress)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing RemoveAllowedMacAddress", err, restyResp1.String(),
				"Failure at RemoveAllowedMacAddress, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing RemoveAllowedMacAddress", err,
			"Failure at RemoveAllowedMacAddress, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSecurityThreatsRogueAllowedListAddAllowedMacAddress(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestDevicesAddAllowedMacAddress {
	request := dnacentersdkgo.RequestDevicesAddAllowedMacAddress{}
	if v := expandRequestSecurityThreatsRogueAllowedListAddAllowedMacAddressItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSecurityThreatsRogueAllowedListAddAllowedMacAddressItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]dnacentersdkgo.RequestItemDevicesAddAllowedMacAddress {
	request := []dnacentersdkgo.RequestItemDevicesAddAllowedMacAddress{}
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
		i := expandRequestSecurityThreatsRogueAllowedListAddAllowedMacAddressItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSecurityThreatsRogueAllowedListAddAllowedMacAddressItem(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestItemDevicesAddAllowedMacAddress {
	request := dnacentersdkgo.RequestItemDevicesAddAllowedMacAddress{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".category")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".category")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".category")))) {
		request.Category = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchDevicesGetAllowedMacAddress(m interface{}, queryParams dnacentersdkgo.GetAllowedMacAddressQueryParams, vID string) (*dnacentersdkgo.ResponseItemDevicesGetAllowedMacAddress, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemDevicesGetAllowedMacAddress
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Devices.GetAllowedMacAddress(nil)
		maxPageSize := len(*nResponse)
		for len(*nResponse) > 0 {
			for _, item := range *nResponse {
				if vID == item.MacAddress {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Devices.GetAllowedMacAddress(&queryParams)
		}
		return nil, err
	}
	return foundItem, err
}
