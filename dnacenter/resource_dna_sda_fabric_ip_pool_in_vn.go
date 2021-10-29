package dnacenter

import (
	"context"
	"strings"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSDAFabricIPPoolInVN() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricIPPoolInVNCreate,
		ReadContext:   resourceSDAFabricIPPoolInVNRead,
		UpdateContext: resourceSDAFabricIPPoolInVNUpdate,
		DeleteContext: resourceSDAFabricIPPoolInVNDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"virtual_network_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_pool_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"authentication_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"scalable_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"is_l2_flooding_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true, //REVIEW: It may be only Optional & Computed
			},
			"is_this_critical_pool": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true, //REVIEW: It may be only Optional & Computed
			},
			"pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true, //REVIEW: It may be only Optional & Computed
			},
		},
	}
}

func resourceSDAFabricIPPoolInVNCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	ipPoolName := d.Get("ip_pool_name").(string)
	virtualNetworkName := d.Get("virtual_network_name").(string)

	searchResponse, _, err := client.SDA.GetIPPoolFromSDAVirtualNetwork(&dnac.GetIPPoolFromSDAVirtualNetworkQueryParams{
		IPPoolName:         ipPoolName,
		VirtualNetworkName: virtualNetworkName,
	})
	if err == nil && searchResponse != nil {
		if virtualNetworkName == searchResponse.VirtualNetworkName {
			// Update resource id
			d.SetId(strings.Join([]string{ipPoolName, virtualNetworkName}, "_/_"))
			resourceSDAFabricIPPoolInVNRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddIPPoolInSDAVirtualNetworkRequest
	var request dnac.AddIPPoolInSDAVirtualNetworkRequest
	request.IPPoolName = ipPoolName
	request.VirtualNetworkName = virtualNetworkName
	if v, ok := d.GetOk("authentication_policy_name"); ok {
		request.AuthenticationPolicyName = v.(string)
	}
	if v, ok := d.GetOk("is_l2_flooding_enabled"); ok {
		request.IsL2FloodingEnabled = v.(bool)
	}
	if v, ok := d.GetOk("is_this_critical_pool"); ok {
		request.IsThisCriticalPool = v.(bool)
	}
	if v, ok := d.GetOk("pool_type"); ok {
		request.PoolType = v.(string)
	}
	if v, ok := d.GetOk("scalable_group_name"); ok {
		request.ScalableGroupName = v.(string)
	}
	if v, ok := d.GetOk("traffic_type"); ok {
		request.TrafficType = v.(string)
	}

	requests = append(requests, request)
	addResponse, _, err := client.SDA.AddIPPoolInSDAVirtualNetwork(&requests)
	if err != nil {
		return diag.FromErr(err)
	}
	if addResponse != nil {
		if addResponse.Status == "failed" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "May have been unable to create SDA fabric IP pool in VN",
				Detail:   addResponse.Description,
			})
			return diags
		}
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(strings.Join([]string{ipPoolName, virtualNetworkName}, "_/_"))
	resourceSDAFabricIPPoolInVNRead(ctx, d, m)
	return diags
}

func resourceSDAFabricIPPoolInVNRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	ipPoolName, virtualNetworkName := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetIPPoolFromSDAVirtualNetwork(&dnac.GetIPPoolFromSDAVirtualNetworkQueryParams{
		IPPoolName:         ipPoolName,
		VirtualNetworkName: virtualNetworkName,
	})
	if err != nil || searchResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric IP pool in VN",
		})
		// REVIEW:.
		return diags
	}
	if virtualNetworkName == searchResponse.VirtualNetworkName {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric IP pool in VN",
		})
		// REVIEW:.
		return diags
	}

	if err := d.Set("virtual_network_name", searchResponse.VirtualNetworkName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ip_pool_name", searchResponse.IPPoolName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("traffic_type", searchResponse.TrafficType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("authentication_policy_name", searchResponse.AuthenticationPolicyName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scalable_group_name", searchResponse.ScalableGroupName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_l2_flooding_enabled", searchResponse.IsL2FloodingEnabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_this_critical_pool", searchResponse.IsThisCriticalPool); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSDAFabricIPPoolInVNUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceSDAFabricIPPoolInVNRead(ctx, d, m)
}

func resourceSDAFabricIPPoolInVNDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	ipPoolName, virtualNetworkName := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetIPPoolFromSDAVirtualNetwork(&dnac.GetIPPoolFromSDAVirtualNetworkQueryParams{
		IPPoolName:         ipPoolName,
		VirtualNetworkName: virtualNetworkName,
	})
	if err != nil || searchResponse == nil {
		return diags
	}
	if virtualNetworkName == searchResponse.VirtualNetworkName {
		return diags
	}

	// Call function to delete resource
	deleteRequest := []dnac.DeleteIPPoolFromSDAVirtualNetworkRequest{}
	_, _, err = client.SDA.DeleteIPPoolFromSDAVirtualNetwork(&dnac.DeleteIPPoolFromSDAVirtualNetworkQueryParams{
		IPPoolName:         ipPoolName,
		VirtualNetworkName: virtualNetworkName,
	}, &deleteRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.SDA.GetIPPoolFromSDAVirtualNetwork(&dnac.GetIPPoolFromSDAVirtualNetworkQueryParams{
		IPPoolName:         ipPoolName,
		VirtualNetworkName: virtualNetworkName,
	})
	if err == nil && searchResponse != nil {
		// Check if element already exists
		if virtualNetworkName == searchResponse.VirtualNetworkName {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete SDA fabric IP pool in VN",
			})
		}
		return diags
	}

	return diags
}
