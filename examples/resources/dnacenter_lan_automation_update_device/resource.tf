
resource "dnacenter_lan_automation_update_device" "example" {
  provider = dnacenter
  feature  = "string"
  parameters {

    hostname_update_devices {

      device_management_ipaddress = "string"
      new_host_name               = "string"
    }
    link_update {

      destination_device_interface_name       = "string"
      destination_device_management_ipaddress = "string"
      ip_pool_name                            = "string"
      source_device_interface_name            = "string"
      source_device_management_ipaddress      = "string"
    }
    loopback_update_device_list {

      device_management_ipaddress = "string"
      new_loopback0_ipaddress     = "string"
    }
  }
}

output "dnacenter_lan_automation_update_device_example" {
  value = dnacenter_lan_automation_update_device.example
}