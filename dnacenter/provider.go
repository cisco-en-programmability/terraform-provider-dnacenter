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
			// "dna_tag_member":                      resourceTagMember(), //REVIEW: Missing documentation for POST operation
			"dna_applications":                                resourceApplication(),                           //Done. CRUD.
			"dna_application_set":                             resourceApplicationSet(),                        //Done. CRUD.
			"dna_tag":                                         resourceTag(),                                   //Done. CRUD.
			"dna_site":                                        resourceSite(),                                  //Done. CRUD.
			"dna_template":                                    resourceTemplate(),                              //Done. CRUD.
			"dna_template_project":                            resourceTemplateProject(),                       //Done. CRUD. Update did not update anything besides name, and their times.
			"dna_cli_credential":                              resourceCLICredential(),                         //Done. CRUD.
			"dna_http_read_credential":                        resourceHTTPReadCredential(),                    //Done. CRUD.
			"dna_http_write_credential":                       resourceHTTPWriteCredential(),                   //Done. CRUD.
			"dna_snmpv2_read_community_credential":            resourceSNMPReadCommunityCredential(),           //Done. CRUD.
			"dna_snmpv2_write_community_credential":           resourceSNMPWriteCommunityCredential(),          //Done. CRUD.
			"dna_snmpv3_credential":                           resourceSNMPv3Credential(),                      //Done. CRUD.
			"dna_netconf_credential":                          resourceNetconfCredential(),                     //Done. CRUD.
			"dna_discovery":                                   resourceDiscovery(),                             //Done. CRUD.
			"dna_pnp_device":                                  resourcePnPDevice(),                             //Done. CRUD. May require API definition to search PnP device
			"dna_pnp_workflow":                                resourcePnPWorkflow(),                           //Done. CRUD.
			"dna_pnp_global_settings":                         resourcePnPGlobalSettings(),                     //Done. Read Update operations. Create & Delete only for Terraform.
			"dna_network":                                     resourceNetwork(),                               //Done. Create Read Update operations. Delete only for Terraform. May require to change to pointers for all properties & changes in mapping.
			"dna_network_credential_site_assignment":          resourceNetworkCredentialSiteAssignment(),       //Done. Create Read Update operations. Delete only for Terraform.
			"dna_network_global_ip_pool":                      resourceNetworkGlobalIPPool(),                   //Done. CRUD. Could not test well because unable to create one.
			"dna_network_service_provider_profile":            resourceNetworkServiceProviderProfile(),         //Done. CRUD.
			"dna_sda_fabric":                                  resourceSDAFabric(),                             //Done. Only has Create Read Delete operations. Pending tests
			"dna_sda_fabric_virtual_network":                  resourceSDAFabricVirtualNetwork(),               //Done. Only has Create Read Delete operations. Pending tests
			"dna_sda_fabric_edge_device":                      resourceSDAFabricEdgeDevice(),                   //Done. Only has Create Read Delete operations. Pending tests
			"dna_sda_fabric_port_assignment_for_access_point": resourceSDAFabricPortAssignmentForAccessPoint(), //Done. Only has Create Read Delete operations. Pending tests
			"dna_sda_fabric_ip_pool_in_vn":                    resourceSDAFabricIPPoolInVN(),                   //Done. Only has Create Read Delete operations. Pending tests
			"dna_sda_fabric_site":                             resourceSDAFabricSite(),                         //Done. Only has Create Read Delete operations. Pending tests
			"dna_sda_fabric_border_device":                    resourceSDAFabricBorderDevice(),                 //
			"dna_sda_fabric_control_plane_device":             resourceSDAFabricControlPlaneDevice(),           //
			"dna_sda_fabric_port_assignment_for_user_device":  resourceSDAFabricPortAssignmentForUserDevice(),  //
			"dna_sda_fabric_authentication_profile":           resourceSDAFabricAuthenticationProfile(),        //
		},
		DataSourcesMap: map[string]*schema.Resource{
			"dna_tag_count":                        dataSourceTagCount(),
			"dna_tag_member_count":                 dataSourceTagMemberCount(),
			"dna_tag_member_type":                  dataSourceTagMemberType(),
			"dna_tag_member":                       dataSourceTagMemberQuery(),
			"dna_tag":                              dataSourceTagQuery(),
			"dna_applications_count":               dataSourceApplicationsCount(),
			"dna_applications":                     dataSourceApplications(),
			"dna_application_set":                  dataSourceApplicationSet(),
			"dna_site":                             dataSourceSite(),
			"dna_site_count":                       dataSourceSiteCount(),
			"dna_site_health":                      dataSourceSiteHealth(),
			"dna_site_membership":                  dataSourceSiteMembership(),
			"dna_template_project":                 dataSourceTemplateProject(),
			"dna_template":                         dataSourceTemplate(),
			"dna_template_details":                 dataSourceTemplateDetails(),
			"dna_template_version":                 dataSourceTemplateVersion(),
			"dna_template_preview":                 dataSourceTemplatePreview(),
			"dna_template_deploy":                  dataSourceTemplateDeploy(),
			"dna_template_deploy_status":           dataSourceTemplateDeployStatus(),
			"dna_discovery_device_count":           dataSourceDiscoveryDeviceCount(),
			"dna_discovery_count":                  dataSourceDiscoveryCount(),
			"dna_discovery_range_delete":           dataSourceDiscoveryRangeDelete(),
			"dna_discovery_all_delete":             dataSourceDiscoveryAllDelete(),
			"dna_discovery_jobs":                   dataSourceDiscoveryJobs(),
			"dna_global_credentials":               dataSourceDiscoveryGlobalCredentials(),
			"dna_discovery_device":                 dataSourceDiscoveryDevices(),
			"dna_discovery_device_range":           dataSourceDiscoveryDevicesRange(),
			"dna_discovery_job":                    dataSourceDiscoveryJob(),
			"dna_discovery_range":                  dataSourceDiscoveryRange(),
			"dna_discovery_snmp_property":          dataSourceDiscoverySNMPProperty(),
			"dna_discovery_snmp_property_add":      dataSourceDiscoverySNMPPropertyAdd(),
			"dna_discovery_summary":                dataSourceDiscoverySummary(),
			"dna_pnp_device_sync_result_vacct":     dataSourcePnPDeviceSyncResultVacct(),
			"dna_pnp_workflow_count":               dataSourcePnPWorkflowCount(),
			"dna_pnp_workflow":                     dataSourcePnPWorkflow(),
			"dna_pnp_global_settings":              dataSourcePnPGlobalSettings(),
			"dna_pnp_virtual_account":              dataSourcePnPVaact(),
			"dna_pnp_smart_account":                dataSourcePnPSaact(),
			"dna_pnp_device":                       dataSourcePnPDevice(),
			"dna_pnp_device_count":                 dataSourcePnPDeviceCount(),
			"dna_pnp_device_history":               dataSourcePnPDeviceHistory(),
			"dna_pnp_device_unclaim":               dataSourcePnPDeviceUnclaim(),
			"dna_pnp_device_claim_site":            dataSourcePnPDeviceClaimSite(),
			"dna_pnp_device_reset":                 dataSourcePnPDeviceReset(),
			"dna_pnp_device_sync_vaact":            dataSourcePnPDeviceSyncVacct(),
			"dna_pnp_device_config_preview":        dataSourcePnPDeviceConfigPreview(),
			"dna_pnp_device_claim":                 dataSourcePnPDeviceClaim(),
			"dna_network":                          dataSourceNetwork(),
			"dna_network_device_credential":        dataSourceNetworkDeviceCredential(),
			"dna_network_global_ip_pool":           dataSourceNetworkGlobalIPPool(),
			"dna_network_service_provider_profile": dataSourceNetworkServiceProviderProfile(),

			"dna_command_runner_keywords": dataSourceCommandRunnerKeywords(),
			// "dna_command_runner_run_command": dataSourceCommandRunnerRunCommand(), //REVIEW: For full potential requires Task, Network-Device, File data sources
		},
		ConfigureContextFunc: providerConfigure,
	}
}
