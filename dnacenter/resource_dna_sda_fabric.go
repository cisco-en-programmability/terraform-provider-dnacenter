package dnacenter

import (
	"context"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSDAFabric() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricCreate,
		ReadContext:   resourceSDAFabricRead,
		DeleteContext: resourceSDAFabricDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"fabric_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fabric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabric_domain_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSDAFabricCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	fabricName := d.Get("fabric_name").(string)

	searchResponse, _, err := client.SDA.GetSDAFabricInfo(&dnac.GetSDAFabricInfoQueryParams{
		FabricName: fabricName,
	})
	if err == nil && searchResponse != nil {
		if fabricName == searchResponse.FabricName {
			// Update resource id
			d.SetId(fabricName)
			resourceSDAFabricRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddFabricRequest
	requests = append(requests, dnac.AddFabricRequest{
		FabricName: fabricName,
	})
	_, _, err = client.SDA.AddFabric(&requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(fabricName)
	resourceSDAFabricRead(ctx, d, m)
	return diags
}

func resourceSDAFabricRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	fabricName := d.Id()

	searchResponse, _, err := client.SDA.GetSDAFabricInfo(&dnac.GetSDAFabricInfoQueryParams{
		FabricName: fabricName,
	})
	if err != nil || searchResponse == nil {
		d.SetId("")
		return diags
	}

	if fabricName != searchResponse.FabricName {
		d.SetId("")
		return diags
	}

	if err := d.Set("fabric_name", searchResponse.FabricName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fabric_type", searchResponse.FabricType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fabric_domain_type", searchResponse.FabricDomainType); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSDAFabricDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	fabricName := d.Id()

	searchResponse, _, err := client.SDA.GetSDAFabricInfo(&dnac.GetSDAFabricInfoQueryParams{
		FabricName: fabricName,
	})
	if err != nil || searchResponse == nil {
		return diags
	}
	if fabricName != searchResponse.FabricName {
		return diags
	}

	// Call function to delete resource
	deleteRequest := []dnac.DeleteSDAFabricRequest{}
	_, _, err = client.SDA.DeleteSDAFabric(&dnac.DeleteSDAFabricQueryParams{
		FabricName: fabricName,
	}, &deleteRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.SDA.GetSDAFabricInfo(&dnac.GetSDAFabricInfoQueryParams{
		FabricName: fabricName,
	})

	if err == nil && searchResponse != nil {
		// Check if element already exists
		if fabricName == searchResponse.FabricName {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete SDA fabric",
			})
		}
		return diags
	}

	return diags
}
