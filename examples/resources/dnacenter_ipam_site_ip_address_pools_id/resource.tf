
resource "dnacenter_ipam_site_ip_address_pools_id" "example" {
  provider = dnacenter

  parameters {

    id = "string"
    ip_v4_address_space {

      dhcp_servers       = ["string"]
      dns_servers        = ["string"]
      gateway_ip_address = "string"
      global_pool_id     = "string"
      prefix_length      = 1.0
      slaac_support      = "false"
      subnet             = "string"
    }
    ip_v6_address_space {

      dhcp_servers       = ["string"]
      dns_servers        = ["string"]
      gateway_ip_address = "string"
      global_pool_id     = "string"
      prefix_length      = 1.0
      slaac_support      = "false"
      subnet             = "string"
    }
    name      = "string"
    pool_type = "string"
    site_id   = "string"
  }
}

output "dnacenter_ipam_site_ip_address_pools_id_example" {
  value = dnacenter_ipam_site_ip_address_pools_id.example
}
