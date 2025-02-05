
resource "dnacenter_ipam_global_ip_address_pools" "example" {
  provider = dnacenter

  parameters {

    address_space {

      dhcp_servers       = ["string"]
      dns_servers        = ["string"]
      gateway_ip_address = "string"
      prefix_length      = 1.0
      subnet             = "string"
    }
    name      = "string"
    pool_type = "string"
  }
}

output "dnacenter_ipam_global_ip_address_pools_example" {
  value = dnacenter_ipam_global_ip_address_pools.example
}
