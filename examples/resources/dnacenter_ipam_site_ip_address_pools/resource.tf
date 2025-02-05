
resource "dnacenter_ipam_site_ip_address_pools" "example" {
  provider = dnacenter

  parameters {

    ip_v4_address_space {

      assigned_addresses         = "string"
      default_assigned_addresses = "string"
      dhcp_servers               = ["string"]
      dns_servers                = ["string"]
      gateway_ip_address         = "string"
      global_pool_id             = "string"
      prefix_length              = 1.0
      slaac_support              = "false"
      subnet                     = "string"
      total_addresses            = "string"
      unassignable_addresses     = "string"
    }
    ip_v6_address_space {

      assigned_addresses         = "string"
      default_assigned_addresses = "string"
      dhcp_servers               = ["string"]
      dns_servers                = ["string"]
      gateway_ip_address         = "string"
      global_pool_id             = "string"
      prefix_length              = 1.0
      slaac_support              = "false"
      subnet                     = "string"
      total_addresses            = "string"
      unassignable_addresses     = "string"
    }
    name      = "string"
    pool_type = "string"
    site_id   = "string"
    site_name = "string"
  }
}

output "dnacenter_ipam_site_ip_address_pools_example" {
  value = dnacenter_ipam_site_ip_address_pools.example
}
