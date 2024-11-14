
resource "dnacenter_lan_automation_v2" "example" {
  provider = dnacenter
  parameters {

    discovered_device_site_name_hierarchy = "string"
    discovery_devices {

      device_host_name            = "string"
      device_management_ipaddress = "string"
      device_serial_number        = "string"
      device_site_name_hierarchy  = "string"
    }
    discovery_level   = 1
    discovery_timeout = 1
    host_name_file_id = "string"
    host_name_prefix  = "string"
    ip_pools {

      ip_pool_name = "string"
      ip_pool_role = "string"
    }
    isis_domain_pwd                    = "string"
    multicast_enabled                  = "false"
    peer_device_managment_ipaddress    = "string"
    primary_device_interface_names     = ["string"]
    primary_device_managment_ipaddress = "string"
    redistribute_isis_to_bgp           = "false"
  }
}

output "dnacenter_lan_automation_v2_example" {
  value = dnacenter_lan_automation_v2.example
}