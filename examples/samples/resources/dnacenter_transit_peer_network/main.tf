terraform {
  required_providers {
    dnacenter = {
      version = "1.0.12-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_transit_peer_network" "example" {
  provider = dnacenter
  parameters {
    ip_transit_settings {
      autonomous_system_number = "11"
      routing_protocol_name    = "BGP"
    }
    /*sda_transit_settings{
        transit_control_plane_settings{
            device_management_ip_address="string"
            site_name_hierarchy="string"
        }
    }*/
    transit_peer_network_name = "string"
    transit_peer_network_type = "ip_transit"
  }
}

output "dnacenter_transit_peer_network_example" {
  value = dnacenter_transit_peer_network.example
}
