
resource "dnacenter_global_pool" "example" {
  provider = dnacenter
  parameters {

    id = "string"
    settings {

      ippool {

        ip_address_space = "string"
        dhcp_server_ips  = ["string"]
        dns_server_ips   = ["string"]
        gateway          = "string"
        id               = "string"
        ip_pool_cidr     = "string"
        ip_pool_name     = "string"
        type             = "string"
      }
    }
  }
}

output "dnacenter_global_pool_example" {
  value = dnacenter_global_pool.example
}