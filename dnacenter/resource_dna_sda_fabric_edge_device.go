package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSDAFabricEdgeDevice() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricEdgeDeviceCreate,
		ReadContext:   resourceSDAFabricEdgeDeviceRead,
		DeleteContext: resourceSDAFabricEdgeDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"site_name_hierarchy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"device_management_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
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

func resourceSDAFabricEdgeDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	deviceManagementIPAddress := d.Get("device_management_ip_address").(string)
	siteNameHierarchy := d.Get("site_name_hierarchy").(string)

	searchResponse, _, err := client.SDA.GetEdgeDeviceFromSDAFabric(&dnac.GetEdgeDeviceFromSDAFabricQueryParams{
		DeviceIPAddress: deviceManagementIPAddress,
	})
	if err == nil && searchResponse != nil {
		if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
			// Update resource id
			d.SetId(deviceManagementIPAddress)
			resourceSDAFabricEdgeDeviceRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddEdgeDeviceInSDAFabricRequest
	requests = append(requests, dnac.AddEdgeDeviceInSDAFabricRequest{
		SiteNameHierarchy:         siteNameHierarchy,
		DeviceManagementIPAddress: deviceManagementIPAddress,
	})
	_, _, err = client.SDA.AddEdgeDeviceInSDAFabric(&requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(deviceManagementIPAddress)
	resourceSDAFabricEdgeDeviceRead(ctx, d, m)
	return diags
}

func resourceSDAFabricEdgeDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	deviceManagementIPAddress := d.Id()

	searchResponse, _, err := client.SDA.GetEdgeDeviceFromSDAFabric(&dnac.GetEdgeDeviceFromSDAFabricQueryParams{
		DeviceIPAddress: deviceManagementIPAddress,
	})
	if err != nil || searchResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric edge device",
		})
		// REVIEW:.
		return diags
	}
	if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric edge device",
		})
		// REVIEW:.
		return diags
	}

	if err := d.Set("name", searchResponse.Name); err != nil {
		return diag.FromErr(err)
	}
	// REVIEW: Is SiteHierarchy the value of site_name_hierarchy ?
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

func resourceSDAFabricEdgeDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	deviceManagementIPAddress := d.Id()

	searchResponse, _, err := client.SDA.GetEdgeDeviceFromSDAFabric(&dnac.GetEdgeDeviceFromSDAFabricQueryParams{
		DeviceIPAddress: deviceManagementIPAddress,
	})
	if err != nil || searchResponse == nil {
		return diags
	}
	if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
		return diags
	}

	// Call function to delete resource
	_, _, err = client.SDA.DeleteEdgeDeviceFromSDAFabric(&dnac.DeleteEdgeDeviceFromSDAFabricQueryParams{
		DeviceIPAddress: deviceManagementIPAddress,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.SDA.GetEdgeDeviceFromSDAFabric(&dnac.GetEdgeDeviceFromSDAFabricQueryParams{
		DeviceIPAddress: deviceManagementIPAddress,
	})
	if err == nil && searchResponse != nil {
		// Check if element already exists
		if deviceManagementIPAddress == searchResponse.DeviceManagementIPAddress {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete SDA fabric edge device",
			})
		}
		return diags
	}

	return diags
}
