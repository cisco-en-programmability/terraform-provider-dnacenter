
resource "dnacenter_wireless_controllers_provision" "example" {
  provider  = dnacenter
  device_id = "string"
  parameters {

    interfaces {

      interface_gateway          = "string"
      interface_ipaddress        = "string"
      interface_name             = "string"
      interface_netmask_in_cid_r = 1
      lag_or_port_number         = 1
      vlan_id                    = 1
    }
    rolling_ap_upgrade {

      ap_reboot_percentage      = 1
      enable_rolling_ap_upgrade = "false"
    }
    skip_ap_provision = "false"
  }
}

output "dnacenter_wireless_controllers_provision_example" {
  value = dnacenter_wireless_controllers_provision.example
}