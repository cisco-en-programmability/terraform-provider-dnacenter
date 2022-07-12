
resource "dnacenter_sda_fabric_border_device" "example" {
  provider = dnacenter
  parameters {
    payload {
      border_session_type                = "string"
      border_with_external_connectivity  = "false"
      connected_to_internet              = "false"
      device_management_ip_address       = "string"
      device_role                        = ["string"]
      external_connectivity_ip_pool_name = "string"
      external_connectivity_settings {

        external_autonomou_system_number = "string"
        interface_description            = "string"
        interface_name                   = "string"
        l2_handoff {

          virtual_network_name = "string"
          vlan_name            = "string"
        }
        l3_handoff {

          virtual_network {

            virtual_network_name = "string"
            vlan_id              = "string"
          }
        }
      }
      external_domain_routing_protocol_name = "string"
      internal_autonomou_system_number      = "string"
      sda_transit_network_name              = "string"
      site_name_hierarchy                   = "string"
    }
  }
}

output "dnacenter_sda_fabric_border_device_example" {
  value = dnacenter_sda_fabric_border_device.example
}