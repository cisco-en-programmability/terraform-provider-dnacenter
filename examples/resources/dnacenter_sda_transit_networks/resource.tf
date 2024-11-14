
resource "dnacenter_sda_transit_networks" "example" {
  provider = dnacenter

  parameters {

    id = "string"
    ip_transit_settings {

      autonomous_system_number = "string"
      routing_protocol_name    = "string"
    }
    name = "string"
    sda_transit_settings {

      control_plane_network_device_ids  = ["string"]
      is_multicast_over_transit_enabled = "false"
    }
    type = "string"
  }
}

output "dnacenter_sda_transit_networks_example" {
  value = dnacenter_sda_transit_networks.example
}