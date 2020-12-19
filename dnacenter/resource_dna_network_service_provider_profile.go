package dnacenter

import (
	"context"
	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkServiceProviderProfile() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceNetworkServiceProviderProfileCreate,
		ReadContext:   resourceNetworkServiceProviderProfileRead,
		UpdateContext: resourceNetworkServiceProviderProfileUpdate,
		DeleteContext: resourceNetworkServiceProviderProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{

			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"model": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringHasValueFunc([]string{"4-class-model", "5-class-model", "6-class-model", "8-class-model"}),
			},
			"wan_provider": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// hasServiceProvider checks if it has the Service Provider and if it has the same values
func hasServiceProvider(response *dnac.GetServiceProviderDetailsResponse, profileName, model, wanProvider string) (bool, bool, *dnac.GetServiceProviderDetailsResponseResponseValue) {
	if response != nil {
		var foundValue *dnac.GetServiceProviderDetailsResponseResponseValue
		for _, item := range response.Response {
			for _, value := range item.Value {
				if value.SpProfileName == profileName {
					foundValue = &value
					break
				}
			}
		}
		if foundValue != nil {
			return false, (foundValue.SLAProfileName != model || foundValue.WanProvider != wanProvider), foundValue
		}
		return true, false, nil
	}
	return true, false, nil
}

func resourceNetworkServiceProviderProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)
	var diags diag.Diagnostics

	profileName := d.Get("profile_name").(string)
	model := d.Get("model").(string)
	wanProvider := d.Get("wan_provider").(string)

	searchResponse, _, err := client.NetworkSettings.GetServiceProviderDetails()
	if err == nil && searchResponse != nil {
		// Check if element already exists
		_, performUpdate, _ := hasServiceProvider(searchResponse, profileName, model, wanProvider)

		if performUpdate {

			updateRequest := dnac.UpdateSPProfileRequest{
				Settings: dnac.UpdateSPProfileRequestSettings{
					Qos: []dnac.UpdateSPProfileRequestSettingsQos{
						{Model: model, ProfileName: profileName, OldProfileName: profileName, WanProvider: wanProvider},
					},
				},
			}

			_, _, err := client.NetworkSettings.UpdateSPProfile(&updateRequest)
			if err != nil {
				return diag.FromErr(err)
			}

			// Wait for execution status to complete
			time.Sleep(5 * time.Second)

			// Update resource id
			d.SetId(profileName)
			resourceNetworkServiceProviderProfileRead(ctx, d, m)
			return diags
		}
	}

	// Construct payload from resource schema (item)
	createRequest := dnac.CreateSPProfileRequest{
		Settings: dnac.CreateSPProfileRequestSettings{
			Qos: []dnac.CreateSPProfileRequestSettingsQos{
				{Model: model, ProfileName: profileName, WanProvider: wanProvider},
			},
		},
	}
	_, _, err = client.NetworkSettings.CreateSPProfile(&createRequest)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	// Update resource id
	d.SetId(profileName)
	resourceNetworkServiceProviderProfileRead(ctx, d, m)
	return diags
}

func resourceNetworkServiceProviderProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	profileName := d.Id()

	searchResponse, _, err := client.NetworkSettings.GetServiceProviderDetails()
	if err != nil {
		return diag.FromErr(err)
	}
	_, _, foundValue := hasServiceProvider(searchResponse, profileName, "", "")
	if foundValue == nil {
		d.SetId("")
		return diags
	}

	if err := d.Set("profile_name", foundValue.SpProfileName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("model", foundValue.SLAProfileName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("wan_provider", foundValue.WanProvider); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNetworkServiceProviderProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	profileName := d.Id()
	model := d.Get("model").(string)
	wanProvider := d.Get("wan_provider").(string)

	searchResponse, _, err := client.NetworkSettings.GetServiceProviderDetails()
	if err != nil {
		return diag.FromErr(err)
	}
	_, _, foundValue := hasServiceProvider(searchResponse, profileName, "", "")
	if foundValue == nil {
		d.SetId("")
		return diags
	}

	// Check if properties inside resource has changes
	if d.HasChanges("model", "wan_provider") {
		updateRequest := dnac.UpdateSPProfileRequest{
			Settings: dnac.UpdateSPProfileRequestSettings{
				Qos: []dnac.UpdateSPProfileRequestSettingsQos{
					{Model: model, ProfileName: profileName, OldProfileName: profileName, WanProvider: wanProvider},
				},
			},
		}

		_, _, err := client.NetworkSettings.UpdateSPProfile(&updateRequest)
		if err != nil {
			return diag.FromErr(err)
		}

		// Wait for execution status to complete
		time.Sleep(5 * time.Second)

		// Update resource last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Update resource on Terraform by using read function
	return resourceNetworkServiceProviderProfileRead(ctx, d, m)
}

func resourceNetworkServiceProviderProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dnac.Client)

	var diags diag.Diagnostics

	profileName := d.Id()

	searchResponse, _, err := client.NetworkSettings.GetServiceProviderDetails()
	if err == nil && searchResponse != nil {
		// Check if element already exists
		_, _, foundValue := hasServiceProvider(searchResponse, profileName, "", "")
		if foundValue == nil {
			return diags
		}
	}

	// Call function to delete resource
	_, _, err = client.NetworkSettings.DeleteSPProfile(profileName)
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for execution status to complete
	time.Sleep(5 * time.Second)

	searchResponse, _, err = client.NetworkSettings.GetServiceProviderDetails()
	if err == nil && searchResponse != nil {
		// Check if element already exists
		_, _, foundValue := hasServiceProvider(searchResponse, profileName, "", "")
		if foundValue != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to delete Service Provider",
			})
			return diags
		}
	}

	return diags
}
