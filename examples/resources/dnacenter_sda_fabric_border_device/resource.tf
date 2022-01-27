
resource "dnacenter_sda_fabric_border_device" "example" {
  provider = dnacenter
  parameters {

    border_session_type                   = "string"
    connected_to_internet                 = "false"
    device_management_ip_address          = "string"
    external_autonomou_system_number      = "string"
    external_connectivity_ip_pool_name    = "string"
    external_connectivity_settings        = ["string"]
    external_domain_routing_protocol_name = "string"
    interface_name                        = "string"
    internal_autonomou_system_number      = "string"
    l3_handoff                            = ["string"]
    site_name_hierarchy                   = "string"
    virtual_network                       = ["string"]
    virtual_network_name                  = "string"
    vlan_id                               = "string"
  }
}

output "dnacenter_sda_fabric_border_device_example" {
  value = dnacenter_sda_fabric_border_device.example
}