
resource "dnacenter_sda_provision_devices" "example" {
  provider = dnacenter

  parameters {

    id                = "string"
    network_device_id = "string"
    site_id           = "string"
  }
}

output "dnacenter_sda_provision_devices_example" {
  value = dnacenter_sda_provision_devices.example
}