
resource "dnacenter_network_device_config_write_memory" "example" {
  provider = dnacenter
  parameters {

    device_id = ["string"]
  }
}

output "dnacenter_network_device_config_write_memory_example" {
  value = dnacenter_network_device_config_write_memory.example
}