
resource "dnacenter_sda_virtual_network_ip_pool" "example" {
  provider = dnacenter

  parameters {

    auto_generate_vlan_name  = "false"
    ip_pool_name             = "string"
    is_bridge_mode_vm        = "false"
    is_common_pool           = "false"
    is_ip_directed_broadcast = "false"
    is_l2_flooding_enabled   = "false"
    is_layer2_only           = "false"
    is_this_critical_pool    = "false"
    is_wireless_pool         = "false"
    pool_type                = "string"
    scalable_group_name      = "string"
    site_name_hierarchy      = "string"
    traffic_type             = "string"
    virtual_network_name     = "string"
    vlan_id                  = "string"
    vlan_name                = "string"
  }
}

output "dnacenter_sda_virtual_network_ip_pool_example" {
  value = dnacenter_sda_virtual_network_ip_pool.example
}