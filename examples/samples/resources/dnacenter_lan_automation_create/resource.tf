
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.13-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_lan_automation_create" "example" {
  provider = dnacenter
  parameters {

    payload {
      discovered_device_site_name_hierarchy = "string"
      host_name_file_id                     = "string"
      host_name_prefix                      = "string"
      ip_pools {

        ip_pool_name = "string"
        ip_pool_role = "string"
      }
      isis_domain_pwd                    = "string"
      mulitcast_enabled                  = "false"
      peer_device_managment_ipaddress    = "string"
      primary_device_interface_names     = ["string"]
      primary_device_managment_ipaddress = "string"
      redistribute_isis_to_bgp           = "false"
    }
  }
}

output "dnacenter_lan_automation_create_example" {
  value = dnacenter_lan_automation_create.example
}