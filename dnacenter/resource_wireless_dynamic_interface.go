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

func resourceWirelessDynamicInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Wireless.

- API to create or update an dynamic interface

- Delete a dynamic interface
`,

		CreateContext: resourceWirelessDynamicInterfaceCreate,
		ReadContext:   resourceWirelessDynamicInterfaceRead,
		UpdateContext: resourceWirelessDynamicInterfaceUpdate,
		DeleteContext: resourceWirelessDynamicInterfaceDelete,
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

						"interface_name": &schema.Schema{
							Description: `dynamic-interface name
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"vlan_id": &schema.Schema{
							Description: `Vlan Id
`,
							Type:     schema.TypeFloat,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessDynamicInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestWirelessDynamicInterfaceCreateUpdateDynamicInterface(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vInterfaceName, okInterfaceName := resourceItem["interface_name"]
	vvInterfaceName := interfaceToString(vInterfaceName)
	resp1, restyResp1, err := client.Wireless.CreateUpdateDynamicInterface(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateUpdateDynamicInterface", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateUpdateDynamicInterface", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["interface_name"] = vvInterfaceName
	d.SetId(joinResourceID(resourceMap))
	return resourceWirelessDynamicInterfaceRead(ctx, d, m)
}

func resourceWirelessDynamicInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vInterfaceName := resourceMap["interface_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDynamicInterface")
		queryParams1 := dnacentersdkgo.GetDynamicInterfaceQueryParams{}

		if okInterfaceName {
			queryParams1.InterfaceName = vInterfaceName
		}

		response1, restyResp1, err := client.Wireless.GetDynamicInterface(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetDynamicInterface", err,
				"Failure at GetDynamicInterface, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsWirelessGetDynamicInterface(m, response1, nil)
		item1, err := searchWirelessGetDynamicInterface(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when searching item from GetDynamicInterface response", err,
				"Failure when searching item from GetDynamicInterface, unexpected response", ""))
			return diags
		}
		// Review flatten function used
		vItem1 := flattenWirelessGetDynamicInterfaceByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDynamicInterface search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceWirelessDynamicInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceWirelessDynamicInterfaceRead(ctx, d, m)
}

func resourceWirelessDynamicInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vInterfaceName := resourceMap["interface_name"]

	queryParams1 := dnacentersdkgo.GetDynamicInterfaceQueryParams
	queryParams1.InterfaceName = vInterfaceName
	item, err := searchWirelessGetDynamicInterface(m, queryParams1)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetDynamicInterface", err,
			"Failure at GetDynamicInterface, unexpected response", ""))
		return diags
	}

	selectedMethod := 1
	var vvID string
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 1 {

		getResp1, _, err := client.Wireless.GetDynamicInterface(nil)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsWirelessGetDynamicInterface(m, getResp1, nil)
		item1, err := searchWirelessGetDynamicInterface(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vName != item1.Name {
			vvName = item1.Name
		} else {
			vvName = vName
		}
	}
	restyResp1, err := client.Wireless.DeleteDynamicInterface(vvInterfaceName)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteDynamicInterface", err, restyResp1.String(),
				"Failure at DeleteDynamicInterface, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteDynamicInterface", err,
			"Failure at DeleteDynamicInterface, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestWirelessDynamicInterfaceCreateUpdateDynamicInterface(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestWirelessCreateUpdateDynamicInterface {
	request := dnacentersdkgo.RequestWirelessCreateUpdateDynamicInterface{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToFloat64Ptr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchWirelessGetDynamicInterface(m interface{}, queryParams dnacentersdkgo.GetDynamicInterfaceQueryParams) (*dnacentersdkgo.ResponseItemWirelessGetDynamicInterface, error) {
	client := m.(*dnacentersdkgo.Client)
	var err error
	var foundItem *dnacentersdkgo.ResponseItemWirelessGetDynamicInterface
	var ite *dnacentersdkgo.ResponseWirelessGetDynamicInterface
	ite, _, err = client.Wireless.GetDynamicInterface(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *dnacentersdkgo.ResponseItemWirelessGetDynamicInterface
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
