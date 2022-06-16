
data "dnacenter_network_device_module" "example" {
  provider                    = dnacenter
  device_id                   = "string"
  limit                       = "string"
  name_list                   = ["string"]
  offset                      = "string"
  operational_state_code_list = ["string"]
  part_number_list            = ["string"]
  vendor_equipment_type_list  = ["string"]
}

output "dnacenter_network_device_module_example" {
  value = data.dnacenter_network_device_module.example.items
}

data "dnacenter_network_device_module" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_module_example" {
  value = data.dnacenter_network_device_module.example.item
}
