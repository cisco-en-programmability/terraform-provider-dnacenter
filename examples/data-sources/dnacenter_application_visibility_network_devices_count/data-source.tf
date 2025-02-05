
data "dnacenter_application_visibility_network_devices_count" "example" {
  provider                         = dnacenter
  app_telemetry_deployment_status  = "string"
  app_telemetry_readiness_status   = "string"
  application_registry_sync_status = "string"
  cbar_deployment_status           = "string"
  cbar_readiness_status            = "string"
  hostname                         = "string"
  ids                              = "string"
  management_address               = "string"
  protocol_pack_status             = "string"
  protocol_pack_update_status      = "string"
  site_id                          = "string"
}

output "dnacenter_application_visibility_network_devices_count_example" {
  value = data.dnacenter_application_visibility_network_devices_count.example.item
}
