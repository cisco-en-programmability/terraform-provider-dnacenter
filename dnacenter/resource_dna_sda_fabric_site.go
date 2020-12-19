package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSDAFabricSite() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricSiteCreate,
		ReadContext:   resourceSDAFabricSiteRead,
		DeleteContext: resourceSDAFabricSiteDelete,
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
			"fabric_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSDAFabricSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	siteNameHierarchy := d.Get("site_name_hierarchy").(string)
	fabricName := d.Get("fabric_name").(string)

	searchResponse, _, err := client.SDA.GetSiteFromSDAFabric(&dnac.GetSiteFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err == nil && searchResponse != nil {
		if "success" == searchResponse.Status {
			// Update resource id
			d.SetId(strings.Join([]string{fabricName, siteNameHierarchy}, "_/_"))
			resourceSDAFabricSiteRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddSiteInSDAFabricRequest
	requests = append(requests, dnac.AddSiteInSDAFabricRequest{
		FabricName:        fabricName,
		SiteNameHierarchy: siteNameHierarchy,
	})
	_, _, err = client.SDA.AddSiteInSDAFabric(&requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(strings.Join([]string{fabricName, siteNameHierarchy}, "_/_"))
	resourceSDAFabricSiteRead(ctx, d, m)
	return diags
}

func resourceSDAFabricSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	fabricName, siteNameHierarchy := resourceIDs[0], resourceIDs[1]

	searchResponse, _, err := client.SDA.GetSiteFromSDAFabric(&dnac.GetSiteFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err != nil || searchResponse == nil {
		d.SetId("")
		return diags
	}

	if "success" != searchResponse.Status {
		d.SetId("")
		return diags
	}

	if err := d.Set("fabric_name", fabricName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("site_name_hierarchy", siteNameHierarchy); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSDAFabricSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	resourceIDs := strings.Split(d.Id(), "_/_")
	siteNameHierarchy := resourceIDs[1]

	searchResponse, _, err := client.SDA.GetSiteFromSDAFabric(&dnac.GetSiteFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err != nil || searchResponse == nil {
		return diags
	}
	if "success" != searchResponse.Status {
		return diags
	}

	// Call function to delete resource
	deleteRequest := []dnac.DeleteSiteFromSDAFabricRequest{}
	_, _, err = client.SDA.DeleteSiteFromSDAFabric(&dnac.DeleteSiteFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	}, &deleteRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.SDA.GetSiteFromSDAFabric(&dnac.GetSiteFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})

	if err == nil && searchResponse != nil {
		// Check if element already exists
		if "success" == searchResponse.Status {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete SDA fabric site",
			})
		}
		return diags
	}

	return diags
}
