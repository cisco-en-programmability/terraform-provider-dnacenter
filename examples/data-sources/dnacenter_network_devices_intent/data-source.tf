
data "dnacenter_network_devices_intent" "example" {
  provider            = dnacenter
  family              = "string"
  id                  = "string"
  limit               = "string"
  management_address  = "string"
  management_state    = "string"
  offset              = "string"
  order               = "string"
  reachability_status = "string"
  role                = "string"
  serial_number       = "string"
  sort_by             = "string"
  stack_device        = "string"
  status              = "string"
  views               = "string"
}

output "dnacenter_network_devices_intent_example" {
  value = data.dnacenter_network_devices_intent.example.items
}
