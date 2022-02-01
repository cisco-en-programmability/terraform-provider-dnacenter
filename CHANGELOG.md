## 0.1.0-beta.1 (February 01, 2022)

BUG FIXES:
* dnacenter/data_source_system_performance_historical: Remove `t1` computed parameter from schema.

IMPROVEMENTS:
* dnacenter/data_source_pnp_device_claim_to_site: Add parameters `config_info, image_info, hostname` to schema.

## 0.1.0-beta (January 27, 2022)

IMPROVEMENTS:
* Removed comments.

## 0.1.0-alpha (January 27, 2022)

* New Release, supports Cisco DNA Center API 2.2.3.3.

BREAKING CHANGES:
* Drop or renamed previous data sources and resources.

FEATURES:
* **New Data Source:** `dnacenter_reserve_ip_subpool`
* **New Data Source:** `dnacenter_site_assign_device`
* **New Data Source:** `dnacenter_wireless_rf_profile`
* **New Data Source:** `dnacenter_wireless_provision_device_create`
* **New Data Source:** `dnacenter_wireless_provision_device_update`
* **New Data Source:** `dnacenter_wireless_provision_access_point`
* **New Data Source:** `dnacenter_wireless_profile`
* **New Data Source:** `dnacenter_template_preview`
* **New Data Source:** `dnacenter_configuration_template_create`
* **New Data Source:** `dnacenter_configuration_template_project`
* **New Data Source:** `dnacenter_tag_membership`
* **New Data Source:** `dnacenter_tag_member`
* **New Data Source:** `dnacenter_tag`
* **New Data Source:** `dnacenter_service_provider_profile_delete`
* **New Data Source:** `dnacenter_snmp_properties`
* **New Data Source:** `dnacenter_site_update`
* **New Data Source:** `dnacenter_site_delete`
* **New Data Source:** `dnacenter_site_create`
* **New Data Source:** `dnacenter_site`
* **New Data Source:** `dnacenter_service_provider_create`
* **New Data Source:** `dnacenter_service_provider`
* **New Data Source:** `dnacenter_service_provider_update`
* **New Data Source:** `dnacenter_sensor_test_template_duplicate`
* **New Data Source:** `dnacenter_sensor_test_run`
* **New Data Source:** `dnacenter_sensor`
* **New Data Source:** `dnacenter_threat_summary`
* **New Data Source:** `dnacenter_threat_detail_count`
* **New Data Source:** `dnacenter_authentication_import_certificate`
* **New Data Source:** `dnacenter_authentication_import_certificate_p12`
* **New Data Source:** `dnacenter_threat_detail`
* **New Data Source:** `dnacenter_pnp_virtual_account_deregister`
* **New Data Source:** `dnacenter_pnp_server_profile_update`
* **New Data Source:** `dnacenter_pnp_virtual_account_add`
* **New Data Source:** `dnacenter_pnp_global_settings`
* **New Data Source:** `dnacenter_pnp_virtual_account_devices_sync`
* **New Data Source:** `dnacenter_pnp_device_unclaim`
* **New Data Source:** `dnacenter_pnp_device_config_preview`
* **New Data Source:** `dnacenter_pnp_device_claim_to_site`
* **New Data Source:** `dnacenter_pnp_device_reset`
* **New Data Source:** `dnacenter_pnp_device_import`
* **New Data Source:** `dnacenter_pnp_device_claim`
* **New Data Source:** `dnacenter_site_design_floormap`
* **New Data Source:** `dnacenter_nfv_profile`
* **New Data Source:** `dnacenter_nfv_provision_details`
* **New Data Source:** `dnacenter_network_create`
* **New Data Source:** `dnacenter_network_update`
* **New Data Source:** `dnacenter_network`
* **New Data Source:** `dnacenter_network_device_sync`
* **New Data Source:** `dnacenter_network_device_export`
* **New Data Source:** `dnacenter_network_device_update_role`
* **New Data Source:** `dnacenter_command_runner_run_command`
* **New Data Source:** `dnacenter_device_configurations_export`
* **New Data Source:** `dnacenter_network_device_list`
* **New Data Source:** `dnacenter_network_device`
* **New Data Source:** `dnacenter_itsm_integration_events_failed`
* **New Data Source:** `dnacenter_itsm_integration_events_retry`
* **New Data Source:** `dnacenter_swim_import_via_url`
* **New Data Source:** `dnacenter_swim_import_local`
* **New Data Source:** `dnacenter_swim_trigger_distribution`
* **New Data Source:** `dnacenter_swim_trigger_activation`
* **New Data Source:** `dnacenter_global_pool`
* **New Data Source:** `dnacenter_snmpv3_credential_create`
* **New Data Source:** `dnacenter_snmpv3_credential_update`
* **New Data Source:** `dnacenter_snmpv2_write_community_credential_create`
* **New Data Source:** `dnacenter_snmpv2_write_community_credential_update`
* **New Data Source:** `dnacenter_snmpv2_read_community_credential_create`
* **New Data Source:** `dnacenter_snmpv2_read_community_credential_update`
* **New Data Source:** `dnacenter_netconf_credential_create`
* **New Data Source:** `dnacenter_netconf_credential_update`
* **New Data Source:** `dnacenter_http_write_credential_create`
* **New Data Source:** `dnacenter_http_write_credential_update`
* **New Data Source:** `dnacenter_http_read_credential_create`
* **New Data Source:** `dnacenter_http_read_credential_update`
* **New Data Source:** `dnacenter_cli_credential_create`
* **New Data Source:** `dnacenter_cli_credential_update`
* **New Data Source:** `dnacenter_global_credential_update`
* **New Data Source:** `dnacenter_global_credential_delete`
* **New Data Source:** `dnacenter_path_trace`
* **New Data Source:** `dnacenter_path_trace_create`
* **New Data Source:** `dnacenter_path_trace_delete`
* **New Data Source:** `dnacenter_event_subscription_syslog`
* **New Data Source:** `dnacenter_event_subscription_rest`
* **New Data Source:** `dnacenter_event_subscription_email`
* **New Data Source:** `dnacenter_event_subscription`
* **New Data Source:** `dnacenter_wireless_dynamic_interface`
* **New Data Source:** `dnacenter_wireless_psk_override`
* **New Data Source:** `dnacenter_wireless_enterprise_ssid`
* **New Data Source:** `dnacenter_discovery_range_delete`
* **New Data Source:** `dnacenter_discovery`
* **New Data Source:** `dnacenter_device_replacement_deploy`
* **New Data Source:** `dnacenter_device_replacement`
* **New Data Source:** `dnacenter_device_credential_create`
* **New Data Source:** `dnacenter_device_credential_update`
* **New Data Source:** `dnacenter_device_credential`
* **New Data Source:** `dnacenter_device_credential_delete`
* **New Data Source:** `dnacenter_reports`
* **New Data Source:** `dnacenter_site_assign_credential`
* **New Data Source:** `dnacenter_compliance_check_run`
* **New Data Source:** `dnacenter_wireless_provision_ssid_delete_reprovision`
* **New Data Source:** `dnacenter_wireless_provision_ssid_create_provision`
* **New Data Source:** `dnacenter_sda_multicast`
* **New Data Source:** `dnacenter_sda_virtual_network_v2`
* **New Data Source:** `dnacenter_sda_provision_device`
* **New Data Source:** `dnacenter_sda_virtual_network_ip_pool`
* **New Data Source:** `dnacenter_sda_virtual_network`
* **New Data Source:** `dnacenter_sda_port_assignment_for_user_device`
* **New Data Source:** `dnacenter_sda_port_assignment_for_access_point`
* **New Data Source:** `dnacenter_sda_fabric_site`
* **New Data Source:** `dnacenter_sda_fabric`
* **New Data Source:** `dnacenter_sda_fabric_edge_device`
* **New Data Source:** `dnacenter_sda_fabric_control_plane_device`
* **New Data Source:** `dnacenter_sda_fabric_border_device`
* **New Data Source:** `dnacenter_sda_fabric_authentication_profile`
* **New Data Source:** `dnacenter_nfv_provision`
* **New Data Source:** `dnacenter_sensor_test_template_edit`
* **New Data Source:** `dnacenter_applications`
* **New Data Source:** `dnacenter_application_sets`
* **New Data Source:** `dnacenter_pnp_workflow`
* **New Data Source:** `dnacenter_pnp_device`
* **New Data Source:** `dnacenter_event_artifact_count`
* **New Data Source:** `dnacenter_event_artifact`
* **New Data Source:** `dnacenter_user_enrichment_details`
* **New Data Source:** `dnacenter_topology_vlan_details`
* **New Data Source:** `dnacenter_topology_site`
* **New Data Source:** `dnacenter_topology_physical`
* **New Data Source:** `dnacenter_topology_layer_3`
* **New Data Source:** `dnacenter_topology_layer_2`
* **New Data Source:** `dnacenter_configuration_template`
* **New Data Source:** `dnacenter_configuration_template_version`
* **New Data Source:** `dnacenter_configuration_template_version_create`
* **New Data Source:** `dnacenter_configuration_template_deploy`
* **New Data Source:** `dnacenter_configuration_template_deploy_status`
* **New Data Source:** `dnacenter_configuration_template_clone`
* **New Data Source:** `dnacenter_configuration_template_deploy_v2`
* **New Data Source:** `dnacenter_configuration_template_export_project`
* **New Data Source:** `dnacenter_configuration_template_export_template`
* **New Data Source:** `dnacenter_configuration_template_import_project`
* **New Data Source:** `dnacenter_configuration_template_import_template`
* **New Data Source:** `dnacenter_network_device_chassis_details`
* **New Data Source:** `dnacenter_network_device_linecard_details`
* **New Data Source:** `dnacenter_network_device_stack_details`
* **New Data Source:** `dnacenter_network_device_supervisor_card_details`
* **New Data Source:** `dnacenter_network_device_inventory_insight_link_mismatch`
* **New Data Source:** `dnacenter_network_device_with_snmp_v3_des`
* **New Data Source:** `dnacenter_network_device_interface_poe`
* **New Data Source:** `dnacenter_system_health`
* **New Data Source:** `dnacenter_system_health_count`
* **New Data Source:** `dnacenter_system_performance`
* **New Data Source:** `dnacenter_system_performance_historical`
* **New Data Source:** `dnacenter_platform_nodes_configuration_summary`
* **New Data Source:** `dnacenter_platform_release_summary`
* **New Data Source:** `dnacenter_task_operation`
* **New Data Source:** `dnacenter_task_count`
* **New Data Source:** `dnacenter_task_tree`
* **New Data Source:** `dnacenter_task`
* **New Data Source:** `dnacenter_tag_member_type`
* **New Data Source:** `dnacenter_tag_count`
* **New Data Source:** `dnacenter_tag_member_count`
* **New Data Source:** `dnacenter_site_count`
* **New Data Source:** `dnacenter_site_health`
* **New Data Source:** `dnacenter_security_advisories_per_device`
* **New Data Source:** `dnacenter_security_advisories_ids_per_device`
* **New Data Source:** `dnacenter_security_advisories_summary`
* **New Data Source:** `dnacenter_security_advisories_devices`
* **New Data Source:** `dnacenter_security_advisories`
* **New Data Source:** `dnacenter_pnp_workflow_count`
* **New Data Source:** `dnacenter_pnp_virtual_accounts`
* **New Data Source:** `dnacenter_pnp_smart_account_domains`
* **New Data Source:** `dnacenter_pnp_virtual_account_sync_result`
* **New Data Source:** `dnacenter_pnp_device_history`
* **New Data Source:** `dnacenter_pnp_device_count`
* **New Data Source:** `dnacenter_topology_network_health`
* **New Data Source:** `dnacenter_network_device_register_for_wsa`
* **New Data Source:** `dnacenter_network_device_by_serial_number`
* **New Data Source:** `dnacenter_network_device_module_count`
* **New Data Source:** `dnacenter_network_device_module`
* **New Data Source:** `dnacenter_network_device_by_ip`
* **New Data Source:** `dnacenter_network_device_functional_capability`
* **New Data Source:** `dnacenter_network_device_count`
* **New Data Source:** `dnacenter_network_device_config_count`
* **New Data Source:** `dnacenter_network_device_config`
* **New Data Source:** `dnacenter_network_device_global_polling_interval`
* **New Data Source:** `dnacenter_network_device_lexicographically_sorted`
* **New Data Source:** `dnacenter_network_device_range`
* **New Data Source:** `dnacenter_network_device_wireless_lan`
* **New Data Source:** `dnacenter_network_device_vlan`
* **New Data Source:** `dnacenter_network_device_meraki_organization`
* **New Data Source:** `dnacenter_network_device_polling_interval`
* **New Data Source:** `dnacenter_network_device_summary`
* **New Data Source:** `dnacenter_network_device_poe`
* **New Data Source:** `dnacenter_network_device_equipment`
* **New Data Source:** `dnacenter_dna_command_runner_keywords`
* **New Data Source:** `dnacenter_site_membership`
* **New Data Source:** `dnacenter_issues`
* **New Data Source:** `dnacenter_issues_enrichment_details`
* **New Data Source:** `dnacenter_device_interface_ospf`
* **New Data Source:** `dnacenter_interface_network_device_detail`
* **New Data Source:** `dnacenter_interface_network_device`
* **New Data Source:** `dnacenter_interface_network_device_range`
* **New Data Source:** `dnacenter_device_interface_isis`
* **New Data Source:** `dnacenter_device_interface_by_ip`
* **New Data Source:** `dnacenter_device_interface_count`
* **New Data Source:** `dnacenter_device_interface`
* **New Data Source:** `dnacenter_swim_image_details`
* **New Data Source:** `dnacenter_global_credential`
* **New Data Source:** `dnacenter_file_namespaces`
* **New Data Source:** `dnacenter_file_namespace_files`
* **New Data Source:** `dnacenter_file`
* **New Data Source:** `dnacenter_event_count`
* **New Data Source:** `dnacenter_event`
* **New Data Source:** `dnacenter_event_subscription_count`
* **New Data Source:** `dnacenter_event_subscription_details_syslog`
* **New Data Source:** `dnacenter_event_subscription_details_rest`
* **New Data Source:** `dnacenter_event_subscription_details_email`
* **New Data Source:** `dnacenter_event_series_count`
* **New Data Source:** `dnacenter_event_series`
* **New Data Source:** `dnacenter_event_api_status`
* **New Data Source:** `dnacenter_discovery_count`
* **New Data Source:** `dnacenter_discovery_range`
* **New Data Source:** `dnacenter_discovery_summary`
* **New Data Source:** `dnacenter_discovery_device_count`
* **New Data Source:** `dnacenter_discovery_device`
* **New Data Source:** `dnacenter_discovery_device_range`
* **New Data Source:** `dnacenter_discovery_job_by_id`
* **New Data Source:** `dnacenter_discovery_jobs`
* **New Data Source:** `dnacenter_device_replacement_count`
* **New Data Source:** `dnacenter_device_health`
* **New Data Source:** `dnacenter_device_enrichment_details`
* **New Data Source:** `dnacenter_device_details`
* **New Data Source:** `dnacenter_reports_view_group_view`
* **New Data Source:** `dnacenter_reports_view_group`
* **New Data Source:** `dnacenter_reports_executions_download`
* **New Data Source:** `dnacenter_reports_executions`
* **New Data Source:** `dnacenter_compliance_device_status_count`
* **New Data Source:** `dnacenter_compliance_device_details_count`
* **New Data Source:** `dnacenter_compliance_device_details`
* **New Data Source:** `dnacenter_compliance_device_by_id_detail`
* **New Data Source:** `dnacenter_compliance_device`
* **New Data Source:** `dnacenter_compliance_device_by_id`
* **New Data Source:** `dnacenter_itsm_cmdb_sync_status`
* **New Data Source:** `dnacenter_client_proximity`
* **New Data Source:** `dnacenter_client_health`
* **New Data Source:** `dnacenter_client_enrichment_details`
* **New Data Source:** `dnacenter_client_detail`
* **New Data Source:** `dnacenter_sda_count`
* **New Data Source:** `dnacenter_sda_device_role`
* **New Data Source:** `dnacenter_sda_device`
* **New Data Source:** `dnacenter_nfv_provision_detail`
* **New Data Source:** `dnacenter_wireless_sensor_test_results`
* **New Data Source:** `dnacenter_applications_count`
* **New Data Source:** `dnacenter_application_sets_count`
* **New Data Source:** `dnacenter_applications_health`
* **New Data Source:** `dnacenter_event_series_audit_logs`
* **New Data Source:** `dnacenter_event_series_audit_logs_summary`
* **New Data Source:** `dnacenter_event_series_audit_logs_parent_records`
* **New Data Source:** `dnacenter_license_virtual_account_change`
* **New Data Source:** `dnacenter_license_device_count`
* **New Data Source:** `dnacenter_license_device_deregistration`
* **New Data Source:** `dnacenter_license_device_license_details`
* **New Data Source:** `dnacenter_license_device_license_summary`
* **New Data Source:** `dnacenter_license_device_registration`
* **New Data Source:** `dnacenter_license_term_details`
* **New Data Source:** `dnacenter_license_usage_details`
* **New Data Source:** `dnacenter_license_smart_account_details`
* **New Data Source:** `dnacenter_license_virtual_account_details`
* **New Data Source:** `dnacenter_app_policy`
* **New Data Source:** `dnacenter_app_policy_default`
* **New Data Source:** `dnacenter_app_policy_intent_create`
* **New Data Source:** `dnacenter_app_policy_queuing_profile`
* **New Data Source:** `dnacenter_app_policy_queuing_profile_count`
* **New Data Source:** `dnacenter_business_sda_hostonboarding_ssid_ippool`
* **New Data Source:** `dnacenter_business_sda_wireless_controller_create`
* **New Data Source:** `dnacenter_business_sda_wireless_controller_delete`
* **New Data Source:** `dnacenter_disasterrecovery_system_operationstatus`
* **New Data Source:** `dnacenter_disasterrecovery_system_status`
* **New Data Source:** `dnacenter_dnacaap_management_execution_status`
* **New Data Source:** `dnacenter_endpoint_analytics_profiling_rules`
* **New Data Source:** `dnacenter_profiling_rules_in_bulk_create`
* **New Data Source:** `dnacenter_profiling_rules_count`
* **New Data Source:** `dnacenter_device_family_identifiers_details`
* **New Data Source:** `dnacenter_golden_image_create`
* **New Data Source:** `dnacenter_golden_tag_image_details`
* **New Data Source:** `dnacenter_golden_tag_image_delete`
* **New Data Source:** `dnacenter_associate_site_to_network_profile`
* **New Data Source:** `dnacenter_disassociate_site_to_network_profile`
* **New Data Source:** `dnacenter_qos_device_interface`
* **New Data Source:** `dnacenter_qos_device_interface_info_count`
* **New Data Source:** `dnacenter_projects_details`
* **New Data Source:** `dnacenter_templates_details`
* **New Data Source:** `dnacenter_sensor_create`
* **New Data Source:** `dnacenter_sensor_delete`
* **New Data Source:** `dnacenter_tag_member_create`
* **New Data Source:** `dnacenter_tag_member_delete`
* **New Resource:** `dnacenter_reserve_ip_subpool`
* **New Resource:** `dnacenter_wireless_rf_profile`
* **New Resource:** `dnacenter_wireless_profile`
* **New Resource:** `dnacenter_configuration_template_project`
* **New Resource:** `dnacenter_tag`
* **New Resource:** `dnacenter_snmp_properties`
* **New Resource:** `dnacenter_pnp_global_settings`
* **New Resource:** `dnacenter_site_design_floormap`
* **New Resource:** `dnacenter_nfv_profile`
* **New Resource:** `dnacenter_network_device_list`
* **New Resource:** `dnacenter_network_device`
* **New Resource:** `dnacenter_global_pool`
* **New Resource:** `dnacenter_event_subscription_syslog`
* **New Resource:** `dnacenter_event_subscription_rest`
* **New Resource:** `dnacenter_event_subscription_email`
* **New Resource:** `dnacenter_event_subscription`
* **New Resource:** `dnacenter_wireless_dynamic_interface`
* **New Resource:** `dnacenter_wireless_enterprise_ssid`
* **New Resource:** `dnacenter_discovery`
* **New Resource:** `dnacenter_device_replacement`
* **New Resource:** `dnacenter_reports`
* **New Resource:** `dnacenter_sda_multicast`
* **New Resource:** `dnacenter_sda_virtual_network_v2`
* **New Resource:** `dnacenter_sda_provision_device`
* **New Resource:** `dnacenter_sda_virtual_network_ip_pool`
* **New Resource:** `dnacenter_sda_virtual_network`
* **New Resource:** `dnacenter_sda_port_assignment_for_user_device`
* **New Resource:** `dnacenter_sda_port_assignment_for_access_point`
* **New Resource:** `dnacenter_sda_fabric_site`
* **New Resource:** `dnacenter_sda_fabric`
* **New Resource:** `dnacenter_sda_fabric_edge_device`
* **New Resource:** `dnacenter_sda_fabric_control_plane_device`
* **New Resource:** `dnacenter_sda_fabric_border_device`
* **New Resource:** `dnacenter_sda_fabric_authentication_profile`
* **New Resource:** `dnacenter_applications`
* **New Resource:** `dnacenter_application_sets`
* **New Resource:** `dnacenter_pnp_workflow`
* **New Resource:** `dnacenter_pnp_device`
* **New Resource:** `dnacenter_configuration_template`
* **New Resource:** `dnacenter_app_policy_queuing_profile`
* **New Resource:** `dnacenter_business_sda_hostonboarding_ssid_ippool`
* **New Resource:** `dnacenter_endpoint_analytics_profiling_rules`
* **New Resource:** `dnacenter_qos_device_interface`

## 0.0.4 (October 29, 2021)

BUG FIXES:

* dnacenter/data_source_dna_discovery_snmp_property_add: Change comparisson to verify if is present.
* dnacenter/data_source_dna_pnp_device_claim: Change comparisson to verify if is present.
* dnacenter/data_source_dna_pnp_device_reset: Change comparisson to verify if is present.
* dnacenter/data_source_dna_pnp_device_sync_vaact: Change comparisson to verify if is present.
* dnacenter/data_source_dna_template_deploy: Change comparisson to verify if is present.
* dnacenter/resource_dna_cli_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_discovery: Change comparisson to verify if is present.
* dnacenter/resource_dna_http_read_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_http_write_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_netconf_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_network: Change comparisson to verify if is present.
* dnacenter/resource_dna_network_global_ip_pool: Change comparisson to verify if is present.
* dnacenter/resource_dna_pnp_device: Change comparisson to verify if is present.
* dnacenter/resource_dna_pnp_global_settings: Change comparisson to verify if is present.
* dnacenter/resource_dna_pnp_workflow: Change comparisson to verify if is present.
* dnacenter/resource_dna_sda_fabric_border_device: Change comparisson to verify if is present.
* dnacenter/resource_dna_sda_fabric_ip_pool_in_vn: Change comparisson to verify if is present.
* dnacenter/resource_dna_sda_fabric_ip_pool_in_vn: Change comparisson to verify if is present.
* dnacenter/resource_dna_site: Change comparisson to verify if is present.
* dnacenter/resource_dna_snmpv2_read_community_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_snmpv2_write_community_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_snmpv3_credential: Change comparisson to verify if is present.
* dnacenter/resource_dna_tag: Change comparisson to verify if is present.
* dnacenter/resource_dna_template: Change comparisson to verify if is present.

## 0.0.3 (February 01, 2021)

* Initial Release, supports Cisco DNA Center API 2.1.1.

FEATURES:
* **New Data Source:** `dna_application_set`
* **New Data Source:** `dna_applications`
* **New Data Source:** `dna_applications_count`
* **New Data Source:** `dna_command_runner_keywords`
* **New Data Source:** `dna_command_runner_run_command`
* **New Data Source:** `dna_discovery_all_delete`
* **New Data Source:** `dna_discovery_count`
* **New Data Source:** `dna_discovery_device`
* **New Data Source:** `dna_discovery_device_count`
* **New Data Source:** `dna_discovery_device_range`
* **New Data Source:** `dna_discovery_job`
* **New Data Source:** `dna_discovery_jobs`
* **New Data Source:** `dna_discovery_range`
* **New Data Source:** `dna_discovery_range_delete`
* **New Data Source:** `dna_discovery_snmp_property`
* **New Data Source:** `dna_discovery_snmp_property_add`
* **New Data Source:** `dna_discovery_summary`
* **New Data Source:** `dna_global_credentials`
* **New Data Source:** `dna_network`
* **New Data Source:** `dna_network_device_credential`
* **New Data Source:** `dna_network_global_ip_pool`
* **New Data Source:** `dna_network_service_provider_profile`
* **New Data Source:** `dna_pnp_device`
* **New Data Source:** `dna_pnp_device_claim`
* **New Data Source:** `dna_pnp_device_claim_site`
* **New Data Source:** `dna_pnp_device_config_preview`
* **New Data Source:** `dna_pnp_device_count`
* **New Data Source:** `dna_pnp_device_history`
* **New Data Source:** `dna_pnp_device_reset`
* **New Data Source:** `dna_pnp_device_saact`
* **New Data Source:** `dna_pnp_device_sync_result_vacct`
* **New Data Source:** `dna_pnp_device_sync_vaact`
* **New Data Source:** `dna_pnp_device_unclaim`
* **New Data Source:** `dna_pnp_global_settings`
* **New Data Source:** `dna_pnp_vacct`
* **New Data Source:** `dna_pnp_workflow`
* **New Data Source:** `dna_pnp_workflow_count`
* **New Data Source:** `dna_site`
* **New Data Source:** `dna_site_count`
* **New Data Source:** `dna_site_health`
* **New Data Source:** `dna_site_membership`
* **New Data Source:** `dna_tag`
* **New Data Source:** `dna_tag_count`
* **New Data Source:** `dna_tag_member`
* **New Data Source:** `dna_tag_member_count`
* **New Data Source:** `dna_tag_member_type`
* **New Data Source:** `dna_task`
* **New Data Source:** `dna_template`
* **New Data Source:** `dna_template_deploy`
* **New Data Source:** `dna_template_deploy_status`
* **New Data Source:** `dna_template_details`
* **New Data Source:** `dna_template_preview`
* **New Data Source:** `dna_template_project`
* **New Data Source:** `dna_template_version`
* **New Resource:** `dna_application_set`
* **New Resource:** `dna_applications`
* **New Resource:** `dna_cli_credential`
* **New Resource:** `dna_discovery`
* **New Resource:** `dna_http_read_credential`
* **New Resource:** `dna_http_write_credential`
* **New Resource:** `dna_netconf_credential`
* **New Resource:** `dna_network`
* **New Resource:** `dna_network_credential_site_assignment`
* **New Resource:** `dna_network_global_ip_pool`
* **New Resource:** `dna_network_service_provider_profile`
* **New Resource:** `dna_pnp_device`
* **New Resource:** `dna_pnp_global_settings`
* **New Resource:** `dna_pnp_workflow`
* **New Resource:** `dna_sda_fabric`
* **New Resource:** `dna_sda_fabric_authentication_profile`
* **New Resource:** `dna_sda_fabric_border_device`
* **New Resource:** `dna_sda_fabric_control_plane_device`
* **New Resource:** `dna_sda_fabric_edge_device`
* **New Resource:** `dna_sda_fabric_ip_pool_in_vn`
* **New Resource:** `dna_sda_fabric_port_assignment_for_access_point`
* **New Resource:** `dna_sda_fabric_port_assignment_for_user_device`
* **New Resource:** `dna_sda_fabric_site`
* **New Resource:** `dna_sda_fabric_virtual_network`
* **New Resource:** `dna_site`
* **New Resource:** `dna_snmpv2_read_community_credential`
* **New Resource:** `dna_snmpv2_write_community_credential`
* **New Resource:** `dna_snmpv3_credential`
* **New Resource:** `dna_tag`
* **New Resource:** `dna_template`
* **New Resource:** `dna_template_project`
