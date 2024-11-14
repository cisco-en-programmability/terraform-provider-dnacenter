
resource "dnacenter_sda_fabric_devices_layer2_handoffs_ip_transits" "example" {
  provider = dnacenter

  parameters {

    external_connectivity_ip_pool_name = "string"
    fabric_id                          = "string"
    id                                 = "string"
    interface_name                     = "string"
    local_ip_address                   = "string"
    local_ipv6_address                 = "string"
    network_device_id                  = "string"
    remote_ip_address                  = "string"
    remote_ipv6_address                = "string"
    tcp_mss_adjustment                 = 1
    transit_network_id                 = "string"
    virtual_network_name               = "string"
    vlan_id                            = 1
  }
}

output "dnacenter_sda_fabric_devices_layer2_handoffs_ip_transits_example" {
  value = dnacenter_sda_fabric_devices_layer2_handoffs_ip_transits.example
}