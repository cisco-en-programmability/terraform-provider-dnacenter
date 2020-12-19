package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSDAFabricVirtualNetwork() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricVirtualNetworkCreate,
		ReadContext:   resourceSDAFabricVirtualNetworkRead,
		DeleteContext: resourceSDAFabricVirtualNetworkDelete,
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
			"site_name_hierarchy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"device_management_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"roles": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSDAFabricVirtualNetworkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	virtualNetworkName := d.Get("virtual_network_name").(string)
	siteNameHierarchy := d.Get("site_name_hierarchy").(string)

	searchResponse, _, err := client.SDA.GetVNFromSDAFabric(&dnac.GetVNFromSDAFabricQueryParams{
		VirtualNetworkName: virtualNetworkName,
		SiteNameHierarchy:  siteNameHierarchy,
	})
	if err == nil && searchResponse != nil {
		// REVIEW: Comparison options : searchResponse.Status == "success"
		if virtualNetworkName == searchResponse.Name {
			// Update resource id
			d.SetId(strings.Join([]string{virtualNetworkName, siteNameHierarchy}, "_/_"))
			resourceSDAFabricVirtualNetworkRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddVNInSDAFabricRequest
	requests = append(requests, dnac.AddVNInSDAFabricRequest{
		SiteNameHierarchy:  siteNameHierarchy,
		VirtualNetworkName: virtualNetworkName,
	})
	_, _, err = client.SDA.AddVNInSDAFabric(&requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(strings.Join([]string{virtualNetworkName, siteNameHierarchy}, "_/_"))
	resourceSDAFabricVirtualNetworkRead(ctx, d, m)
	return diags
}

func resourceSDAFabricVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	virtualNetworkName, siteNameHierarchy := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetVNFromSDAFabric(&dnac.GetVNFromSDAFabricQueryParams{
		VirtualNetworkName: virtualNetworkName,
		SiteNameHierarchy:  siteNameHierarchy,
	})
	if err != nil || searchResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA Fabric VN",
		})
		// REVIEW:.
		return diags
	}
	if virtualNetworkName == searchResponse.Name {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA Fabric VN",
		})
		// REVIEW:.
		return diags
	}

	if err := d.Set("virtual_network_name", searchResponse.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("site_name_hierarchy", searchResponse.SiteHierarchy); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("device_management_ip_address", searchResponse.DeviceManagementIPAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("roles", searchResponse.Roles); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSDAFabricVirtualNetworkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	virtualNetworkName, siteNameHierarchy := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetVNFromSDAFabric(&dnac.GetVNFromSDAFabricQueryParams{
		VirtualNetworkName: virtualNetworkName,
		SiteNameHierarchy:  siteNameHierarchy,
	})
	if err != nil || searchResponse == nil {
		return diags
	}
	if virtualNetworkName != searchResponse.Name {
		return diags
	}

	// Call function to delete resource
	_, _, err = client.SDA.DeleteVNFromSDAFabric(&dnac.DeleteVNFromSDAFabricQueryParams{
		VirtualNetworkName: virtualNetworkName,
		SiteNameHierarchy:  siteNameHierarchy,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.SDA.GetVNFromSDAFabric(&dnac.GetVNFromSDAFabricQueryParams{
		VirtualNetworkName: virtualNetworkName,
		SiteNameHierarchy:  siteNameHierarchy,
	})

	if err == nil && searchResponse != nil {
		// Check if element already exists
		if virtualNetworkName == searchResponse.Name {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete SDA fabric VN",
			})
		}
		return diags
	}

	return diags
}
