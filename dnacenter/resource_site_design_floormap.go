package dnacenter

import (
	"context"
	"reflect"

	"log"

	dnacentersdkgo "github.com/cisco-en-programmability/dnacenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSiteDesignFloormap() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Site Design.

- Service to create a floor map with callback

- Service to delete an (empty) floor map with callback

- Service to create a floor map with callback
`,

		CreateContext: resourceSiteDesignFloormapCreate,
		ReadContext:   resourceSiteDesignFloormapRead,
		UpdateContext: resourceSiteDesignFloormapUpdate,
		DeleteContext: resourceSiteDesignFloormapDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"floor_id": &schema.Schema{
							Description: `floorId path parameter. Group ID of the floor to be modified
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSiteDesignFloormapCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSiteDesignFloormapCreateFloormap(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vFloorID, okFloorID := resourceItem["floor_id"]
	vvFloorID := interfaceToString(vFloorID)
	if okFloorID && vvFloorID != "" {
		getResponse1, err := client.SiteDesign.ListSpecifiedFloormaps(vvFloorID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["floor_id"] = vvFloorID
			d.SetId(joinResourceID(resourceMap))
			return resourceSiteDesignFloormapRead(ctx, d, m)
		}
	}
	restyResp1, err := client.SiteDesign.CreateFloormap(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateFloormap", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateFloormap", err))
		return diags
	}

	log.Printf("[DEBUG]", vvFloorID)
	resourceMap := make(map[string]string)
	resourceMap["floor_id"] = vvFloorID
	d.SetId(joinResourceID(resourceMap))
	return resourceSiteDesignFloormapRead(ctx, d, m)
}

func resourceSiteDesignFloormapRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFloorID := resourceMap["floor_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: ListSpecifiedFloormaps")
		vvFloorID := vFloorID

		response1, err := client.SiteDesign.ListSpecifiedFloormaps(vvFloorID)

		if err != nil || response1 == nil {
			// diags = append(diags, diagErrorWithAlt(
			// 	"Failure when executing ListSpecifiedFloormaps", err,
			// 	"Failure at ListSpecifiedFloormaps, unexpected response", ""))
			// return diags
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %s", response1.String())

		if err := d.Set("item", response1.String()); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ListSpecifiedFloormaps response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSiteDesignFloormapUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFloorID := resourceMap["floor_id"]
	vvFloorID := vFloorID

	item, err := client.SiteDesign.ListSpecifiedFloormaps(vvFloorID)

	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ListSpecifiedFloormaps", err,
			"Failure at ListSpecifiedFloormaps, unexpected response", ""))
		return diags
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvFloorID)
		request1 := expandRequestSiteDesignFloormapUpdateFloormap(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		restyResp1, err := client.SiteDesign.UpdateFloormap(vvFloorID, request1)
		if err != nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateFloormap", err, restyResp1.String(),
					"Failure at UpdateFloormap, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateFloormap", err,
				"Failure at UpdateFloormap, unexpected response", ""))
			return diags
		}
	}

	return resourceSiteDesignFloormapRead(ctx, d, m)
}

func resourceSiteDesignFloormapDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vFloorID := resourceMap["floor_id"]
	vvFloorID := vFloorID

	item, err := client.SiteDesign.ListSpecifiedFloormaps(vvFloorID)

	if err != nil || item == nil {
		//diags = append(diags, diagErrorWithAlt(
		//	"Failure when executing ListSpecifiedFloormaps", err,
		//	"Failure at ListSpecifiedFloormaps, unexpected response", ""))
		d.SetId("")
		return diags
	}
	restyResp1, err := client.SiteDesign.DeleteFloormap(vvFloorID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteFloormap", err, restyResp1.String(),
				"Failure at DeleteFloormap, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteFloormap", err,
			"Failure at DeleteFloormap, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSiteDesignFloormapCreateFloormap(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignCreateFloormap {
	var request dnacentersdkgo.RequestSiteDesignCreateFloormap
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestSiteDesignFloormapUpdateFloormap(ctx context.Context, key string, d *schema.ResourceData) *dnacentersdkgo.RequestSiteDesignUpdateFloormap {
	var request dnacentersdkgo.RequestSiteDesignUpdateFloormap
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}
