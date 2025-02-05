
data "dnacenter_network_devices_intent_count" "example" {
  provider            = dnacenter
  family              = "string"
  id                  = "string"
  management_address  = "string"
  management_state    = "string"
  reachability_status = "string"
  role                = "string"
  serial_number       = "string"
  stack_device        = "string"
  status              = "string"
}

output "dnacenter_network_devices_intent_count_example" {
  value = data.dnacenter_network_devices_intent_count.example.item
}
