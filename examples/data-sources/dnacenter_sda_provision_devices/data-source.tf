
data "dnacenter_sda_provision_devices" "example" {
  provider          = dnacenter
  id                = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
  site_id           = "string"
}

output "dnacenter_sda_provision_devices_example" {
  value = data.dnacenter_sda_provision_devices.example.items
}
