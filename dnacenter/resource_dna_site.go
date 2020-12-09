package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSite() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceSiteCreate,
		ReadContext:   resourceSiteRead,
		UpdateContext: resourceSiteUpdate,
		DeleteContext: resourceSiteDelete,
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
				MaxItems: 1, // Site has only one object
				Required: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validateStringHasValueFunc([]string{"area", "building", "floor"}),
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"parent_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"latitude": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"longitude": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"height": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"length": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"rf_model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"width": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	item := d.Get("item").([]interface{})[0]
	site := item.(map[string]interface{})

	var typeS, name, parentName string
	if v, ok := site["type"]; ok {
		typeS = v.(string)
	}
	if v, ok := site["name"]; ok {
		name = v.(string)
	}
	if v, ok := site["parent_name"]; ok {
		parentName = v.(string)
	}

	siteRequest := dnac.CreateSiteRequest{Type: typeS}

	if typeS == "area" {
		siteRequest.Site.Area.Name = name
		siteRequest.Site.Area.ParentName = parentName
	}

	if typeS == "building" {
		addressInterface, okAddress := site["address"]
		latitudeInterface, okLatitude := site["latitude"]
		longitudeInterface, okLongitude := site["longitude"]
		if !okAddress || !okLatitude || !okLongitude {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create site building",
				Detail:   "Expecting address, latitude and longitude values.",
			})
			return diags
		}
		address := addressInterface.(string)
		latitude := latitudeInterface.(float64)
		longitude := longitudeInterface.(float64)

		siteRequest.Site.Building.Name = name
		siteRequest.Site.Building.ParentName = parentName
		siteRequest.Site.Building.Address = address
		siteRequest.Site.Building.Latitude = latitude
		siteRequest.Site.Building.Longitude = longitude
	}

	if typeS == "floor" {
		heightInterface, okHeight := site["height"]
		lengthInterface, okLength := site["length"]
		rfModelInterface, okRfModel := site["rf_model"]
		widthInterface, okWidth := site["width"]
		if !okHeight || !okLength || !okRfModel || !okWidth {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create site floor",
				Detail:   "Expecting height, length, rf_model and width values.",
			})
			return diags
		}
		height := heightInterface.(float64)
		length := lengthInterface.(float64)
		rfModel := rfModelInterface.(string)
		width := widthInterface.(float64)

		siteRequest.Site.Floor.Name = name
		siteRequest.Site.Floor.ParentName = parentName
		siteRequest.Site.Floor.Height = height
		siteRequest.Site.Floor.Length = length
		siteRequest.Site.Floor.RfModel = rfModel
		siteRequest.Site.Floor.Width = width
	}

	_, _, err := client.Sites.CreateSite(&siteRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	pathName := []string{parentName, name}
	newName := strings.Join(pathName, "/")

	// Wait for execution status to complete
	time.Sleep(10 * time.Second)

	siteQueryParams := &dnac.GetSiteQueryParams{
		Name: newName,
	}
	// Call function to read site.name
	_, _, err = client.Sites.GetSite(siteQueryParams)
	if err != nil {
		// Resource was not created
		return diag.FromErr(err)
	}

	// Update resource id
	d.SetId(newName)
	// Update resource on Terraform
	resourceSiteRead(ctx, d, m)
	return diags
}

func resourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Get resource id (that's also the value of site.name)
	siteName := d.Id()

	siteQueryParams := &dnac.GetSiteQueryParams{
		Name: siteName,
	}

	// Call function to read site.name
	response, _, err := client.Sites.GetSite(siteQueryParams)
	if err != nil {
		// Resource does not exist
		d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
		return diags
	}

	siteRead := flattenSiteReadItem(response)
	if err := d.Set("item", siteRead); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	// Get resource id (that's also the value of site.id)
	siteName := d.Id()
	if d.HasChange("item") {
		item := d.Get("item").([]interface{})[0]
		site := item.(map[string]interface{})

		siteQueryParams := &dnac.GetSiteQueryParams{
			Name: siteName,
		}
		// Call function to read site.name
		response, _, err := client.Sites.GetSite(siteQueryParams)
		if err != nil {
			// Resource does not exist
			d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
			return diags
		}
		var siteID string
		if len(response.Response) > 0 {
			siteID = response.Response[0].ID
		} else {
			// Resource does not exist
			d.SetId("") // Set the ID to an empty string so Terraform "destroys" the resource in state.
			return diags
		}

		var typeS, name, parentName string
		if v, ok := site["type"]; ok {
			typeS = v.(string)
		}
		if v, ok := site["name"]; ok {
			name = v.(string)
		}
		if v, ok := site["parent_name"]; ok {
			parentName = v.(string)
		}

		siteRequest := dnac.UpdateSiteRequest{Type: typeS}

		if typeS == "area" {
			siteRequest.Site.Area.Name = name
			siteRequest.Site.Area.ParentName = parentName
		}

		if typeS == "building" {
			addressInterface, okAddress := site["address"]
			latitudeInterface, okLatitude := site["latitude"]
			longitudeInterface, okLongitude := site["longitude"]
			if !okAddress || !okLatitude || !okLongitude {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to update site building",
					Detail:   "Expecting address, latitude and longitude values.",
				})
				return diags
			}
			address := addressInterface.(string)
			latitude := latitudeInterface.(float64)
			longitude := longitudeInterface.(float64)

			siteRequest.Site.Building.Name = name
			siteRequest.Site.Building.ParentName = parentName
			siteRequest.Site.Building.Address = address
			siteRequest.Site.Building.Latitude = latitude
			siteRequest.Site.Building.Longitude = longitude
		}

		if typeS == "floor" {
			heightInterface, okHeight := site["height"]
			lengthInterface, okLength := site["length"]
			rfModelInterface, okRfModel := site["rf_model"]
			widthInterface, okWidth := site["width"]
			if !okHeight || !okLength || !okRfModel || !okWidth {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to update site floor",
					Detail:   "Expecting height, length, rf_model and width values.",
				})
				return diags
			}
			height := heightInterface.(float64)
			length := lengthInterface.(float64)
			rfModel := rfModelInterface.(string)
			width := widthInterface.(float64)

			siteRequest.Site.Floor.Name = name
			// siteRequest.Site.Floor.ParentName = parentName
			siteRequest.Site.Floor.Height = height
			siteRequest.Site.Floor.Length = length
			siteRequest.Site.Floor.RfModel = rfModel
			siteRequest.Site.Floor.Width = width
		}

		_, _, err = client.Sites.UpdateSite(siteID, &siteRequest)
		if err != nil {
			return diag.FromErr(err)
		}

		newName := strings.Join([]string{parentName, name}, "/")
		nsiteQueryParams := &dnac.GetSiteQueryParams{
			Name: newName,
		}
		// Call function to read site.name
		nresponse, _, err := client.Sites.GetSite(nsiteQueryParams)
		if err == nil && len(nresponse.Response) > 0 {
			// Update resource last_updated
			d.Set("last_updated", time.Now().Format(time.RFC850))
		}
	}

	return resourceSiteRead(ctx, d, m)
}

func resourceSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	siteName := d.Id()
	siteQueryParams := &dnac.GetSiteQueryParams{
		Name: siteName,
	}
	// Call function to read site.name
	response, _, err := client.Sites.GetSite(siteQueryParams)
	if err != nil {
		return diags
	}
	var siteID string
	if len(response.Response) > 0 {
		siteID = response.Response[0].ID
		_, _, err = client.Sites.DeleteSite(siteID)
		if err != nil {
			return diag.FromErr(err)
		}

		checkResponse, _, err := client.Sites.GetSite(siteQueryParams)
		if err != nil || len(checkResponse.Response) == 0 {
			return diags
		}
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "May have been unable to delete site",
			Detail:   "Check if site exists",
		})
		return diags

	}
	return diags
}
