
data "dnacenter_ipam_global_ip_address_pools" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  order    = "string"
  sort_by  = "string"
}

output "dnacenter_ipam_global_ip_address_pools_example" {
  value = data.dnacenter_ipam_global_ip_address_pools.example.items
}
