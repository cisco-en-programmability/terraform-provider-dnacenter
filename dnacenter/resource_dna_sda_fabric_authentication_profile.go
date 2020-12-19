package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSDAFabricAuthenticationProfile() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSDAFabricAuthenticationProfileCreate,
		ReadContext:   resourceSDAFabricAuthenticationProfileRead,
		UpdateContext: resourceSDAFabricAuthenticationProfileUpdate,
		DeleteContext: resourceSDAFabricAuthenticationProfileDelete,
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
			"authenticate_template_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, //REVIEW: It may be only Optional & Computed
			},
			"authenticate_template_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSDAFabricAuthenticationProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	authenticateTemplateName := d.Get("authenticate_template_name").(string)
	siteNameHierarchy := d.Get("site_name_hierarchy").(string)

	searchResponse, _, err := client.SDA.GetDefaultAuthenticationProfileFromSDAFabric(&dnac.GetDefaultAuthenticationProfileFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err == nil && searchResponse != nil {
		if siteNameHierarchy == searchResponse.SiteNameHierarchy {
			// Update resource id
			d.SetId(siteNameHierarchy)
			resourceSDAFabricAuthenticationProfileRead(ctx, d, m)
			return diags
		}
	}

	var requests []dnac.AddDefaultAuthenticationProfileInSDAFabricRequest
	var request dnac.AddDefaultAuthenticationProfileInSDAFabricRequest
	request.AuthenticateTemplateName = authenticateTemplateName
	request.SiteNameHierarchy = siteNameHierarchy
	requests = append(requests, request)
	_, _, err = client.SDA.AddDefaultAuthenticationProfileInSDAFabric(&requests)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(siteNameHierarchy)
	resourceSDAFabricAuthenticationProfileRead(ctx, d, m)
	return diags
}

func resourceSDAFabricAuthenticationProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	siteNameHierarchy := d.Id()

	searchResponse, _, err := client.SDA.GetDefaultAuthenticationProfileFromSDAFabric(&dnac.GetDefaultAuthenticationProfileFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err != nil || searchResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric authentication profile",
		})
		// REVIEW:.
		return diags
	}
	if siteNameHierarchy == searchResponse.SiteNameHierarchy {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric authentication profile",
		})
		// REVIEW:.
		return diags
	}

	if err := d.Set("authenticate_template_name", searchResponse.AuthenticateTemplateName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("site_name_hierarchy", searchResponse.SiteNameHierarchy); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("authenticate_template_id", searchResponse.AuthenticateTemplateID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSDAFabricAuthenticationProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	siteNameHierarchy := d.Id()
	searchResponse, _, err := client.SDA.GetDefaultAuthenticationProfileFromSDAFabric(&dnac.GetDefaultAuthenticationProfileFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err != nil || searchResponse == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to get SDA fabric authentication profile",
		})
		// REVIEW:.
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChange("authenticate_template_name") {
		authenticateTemplateName := d.Get("authenticate_template_name").(string)

		var requests []dnac.UpdateDefaultAuthenticationProfileInSDAFabricRequest
		var request dnac.UpdateDefaultAuthenticationProfileInSDAFabricRequest
		request.AuthenticateTemplateName = authenticateTemplateName
		request.SiteNameHierarchy = siteNameHierarchy
		requests = append(requests, request)
		response, _, err := client.SDA.UpdateDefaultAuthenticationProfileInSDAFabric(&requests)
		if err != nil {
			return diag.FromErr(err)
		}

		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		// Check if task was completed successfully
		if response.Status != "success" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to update SDA fabric authentication profile",
			})
			return diags
		}

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceSDAFabricAuthenticationProfileRead(ctx, d, m)
}

func resourceSDAFabricAuthenticationProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	siteNameHierarchy := d.Id()

	searchResponse, _, err := client.SDA.GetDefaultAuthenticationProfileFromSDAFabric(&dnac.GetDefaultAuthenticationProfileFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err != nil || searchResponse == nil {
		return diags
	}
	if siteNameHierarchy == searchResponse.SiteNameHierarchy {
		return diags
	}

	// Call function to delete resource
	deleteRequest := []dnac.DeleteDefaultAuthenticationProfileFromSDAFabricRequest{}
	_, _, err = client.SDA.DeleteDefaultAuthenticationProfileFromSDAFabric(&dnac.DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	}, &deleteRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.SDA.GetDefaultAuthenticationProfileFromSDAFabric(&dnac.GetDefaultAuthenticationProfileFromSDAFabricQueryParams{
		SiteNameHierarchy: siteNameHierarchy,
	})
	if err == nil && searchResponse != nil {
		// Check if element already exists
		if siteNameHierarchy == searchResponse.SiteNameHierarchy {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete SDA fabric authentication profile",
			})
		}
		return diags
	}

	return diags
}
