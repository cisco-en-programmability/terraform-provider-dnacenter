
data "dnacenter_application_visibility_network_devices" "example" {
  provider                         = dnacenter
  app_telemetry_deployment_status  = "string"
  app_telemetry_readiness_status   = "string"
  application_registry_sync_status = "string"
  cbar_deployment_status           = "string"
  cbar_readiness_status            = "string"
  hostname                         = "string"
  ids                              = "string"
  limit                            = "string"
  management_address               = "string"
  offset                           = "string"
  order                            = "string"
  protocol_pack_status             = "string"
  protocol_pack_update_status      = "string"
  site_id                          = "string"
  sort_by                          = "string"
}

output "dnacenter_application_visibility_network_devices_example" {
  value = data.dnacenter_application_visibility_network_devices.example.items
}
