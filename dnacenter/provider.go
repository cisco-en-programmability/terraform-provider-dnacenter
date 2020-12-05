package dnacenter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider definition of schema(configuration), resources(CRUD) operations and dataSources(query)
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DNAC_BASE_URL", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DNAC_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DNAC_PASSWORD", nil),
			},
			"debug": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("DNAC_DEBUG", "false"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
			},
			"ssl_verify": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("DNAC_SSL_VERIFY", "true"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dna_applications":                      resourceApplication(),
			"dna_tag":                               resourceTag(),
			"dna_site":                              resourceSite(),
			"dna_template":                          resourceTemplate(),
			"dna_template_project":                  resourceTemplateProject(),
			"dna_cli_credential":                    resourceCLICredential(),
			"dna_http_read_credential":              resourceHTTPReadCredential(),
			"dna_http_write_credential":             resourceHTTPWriteCredential(),
			"dna_snmpv2_read_community_credential":  resourceSNMPReadCommunityCredential(),
			"dna_snmpv2_write_community_credential": resourceSNMPWriteCommunityCredential(),
			"dna_snmpv3_credential":                 resourceSNMPv3Credential(),
			"dna_netconf_credential":                resourceNetconfCredential(),
			"dna_discovery":                         resourceDiscovery(),
			"dna_pnp_device":                        resourcePnPDevice(),
			"dna_pnp_workflow":                      resourcePnPWorkflow(),
			"dna_pnp_global_settings":               resourcePnPGlobalSettings(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"dna_tag_count":                    dataSourceTagCount(),
			"dna_tag_member_count":             dataSourceTagMemberCount(),
			"dna_tag_member_type":              dataSourceTagMemberType(),
			"dna_tag":                          dataSourceTagQuery(),
			"dna_applications_count":           dataSourceApplicationsCount(),
			"dna_applications":                 dataSourceApplications(),
			"dna_site":                         dataSourceSite(),
			"dna_site_count":                   dataSourceSiteCount(),
			"dna_site_health":                  dataSourceSiteHealth(),
			"dna_template_project":             dataSourceTemplateProject(),
			"dna_template":                     dataSourceTemplate(),
			"dna_template_details":             dataSourceTemplateDetails(),
			"dna_template_version":             dataSourceTemplateVersion(),
			"dna_template_preview":             dataSourceTemplatePreview(),
			"dna_template_deploy":              dataSourceTemplateDeploy(),
			"dna_template_deploy_status":       dataSourceTemplateDeployStatus(),
			"dna_pnp_device_sync_result_vacct": dataSourcePnPDeviceSyncResultVacct(),
			"dna_pnp_workflow_count":           dataSourcePnPWorkflowCount(),
			"dna_pnp_workflow":                 dataSourcePnPWorkflow(),
			"dna_pnp_global_settings":          dataSourcePnPGlobalSettings(),
			"dna_pnp_virtual_account":          dataSourcePnPVaact(),
			"dna_pnp_smart_account":            dataSourcePnPSaact(),
			"dna_pnp_device":                   dataSourcePnPDevice(),
			"dna_pnp_device_count":             dataSourcePnPDeviceCount(),
			"dna_pnp_device_history":           dataSourcePnPDeviceHistory(),
			"dna_pnp_device_unclaim":           dataSourcePnPDeviceUnclaim(),
			"dna_pnp_device_claim_site":        dataSourcePnPDeviceClaimSite(),
			"dna_pnp_device_reset":             dataSourcePnPDeviceReset(),
			"dna_pnp_device_sync_vaact":        dataSourcePnPDeviceSyncVacct(),
			"dna_pnp_device_config_preview":    dataSourcePnPDeviceConfigPreview(),
			"dna_pnp_device_claim":             dataSourcePnPDeviceClaim(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
