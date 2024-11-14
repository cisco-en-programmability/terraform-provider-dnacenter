
resource "dnacenter_sda_anycast_gateways" "example" {
  provider = dnacenter

  parameters {

    auto_generate_vlan_name                      = "false"
    fabric_id                                    = "string"
    id                                           = "string"
    ip_pool_name                                 = "string"
    is_critical_pool                             = "false"
    is_group_based_policy_enforcement_enabled    = "false"
    is_intra_subnet_routing_enabled              = "false"
    is_ip_directed_broadcast                     = "false"
    is_layer2_flooding_enabled                   = "false"
    is_multiple_ip_to_mac_addresses              = "false"
    is_supplicant_based_extended_node_onboarding = "false"
    is_wireless_pool                             = "false"
    pool_type                                    = "string"
    security_group_name                          = "string"
    tcp_mss_adjustment                           = 1
    traffic_type                                 = "string"
    virtual_network_name                         = "string"
    vlan_id                                      = 1
    vlan_name                                    = "string"
  }
}

output "dnacenter_sda_anycast_gateways_example" {
  value = dnacenter_sda_anycast_gateways.example
}