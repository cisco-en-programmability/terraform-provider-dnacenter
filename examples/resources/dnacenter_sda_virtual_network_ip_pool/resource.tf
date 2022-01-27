
resource "dnacenter_sda_virtual_network_ip_pool" "example" {
  provider = dnacenter
  parameters {

    authentication_policy_name = "string"
    ip_pool_name               = "string"
    is_l2_flooding_enabled     = "false"
    is_this_critical_pool      = "false"
    is_wireless_pool           = "string"
    pool_type                  = "string"
    scalable_group_name        = "string"
    site_name_hierarchy        = "string"
    traffic_type               = "string"
    virtual_network_name       = "string"
    vlan_name                  = "string"
  }
}

output "dnacenter_sda_virtual_network_ip_pool_example" {
  value = dnacenter_sda_virtual_network_ip_pool.example
}