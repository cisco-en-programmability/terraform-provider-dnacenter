
data "dnacenter_event_config_connector_types" "example" {
  provider = dnacenter
}

output "dnacenter_event_config_connector_types_example" {
  value = data.dnacenter_event_config_connector_types.example.items
}
