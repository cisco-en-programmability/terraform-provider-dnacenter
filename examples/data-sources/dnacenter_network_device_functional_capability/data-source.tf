
data "dnacenter_network_device_functional_capability" "example" {
  provider      = dnacenter
  device_id     = "string"
  function_name = ["string"]
}

output "dnacenter_network_device_functional_capability_example" {
  value = data.dnacenter_network_device_functional_capability.example.items
}

data "dnacenter_network_device_functional_capability" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_functional_capability_example" {
  value = data.dnacenter_network_device_functional_capability.example.item
}
