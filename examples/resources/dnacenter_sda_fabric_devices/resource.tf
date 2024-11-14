
resource "dnacenter_sda_fabric_devices" "example" {
  provider = dnacenter

  parameters {

    border_device_settings {

      border_types = ["string"]
      layer3_settings {

        border_priority                 = 1
        import_external_routes          = "false"
        is_default_exit                 = "false"
        local_autonomous_system_number  = "string"
        prepend_autonomous_system_count = 1
      }
    }
    device_roles      = ["string"]
    fabric_id         = "string"
    id                = "string"
    network_device_id = "string"
  }
}

output "dnacenter_sda_fabric_devices_example" {
  value = dnacenter_sda_fabric_devices.example
}