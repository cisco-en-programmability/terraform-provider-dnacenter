
data "dnacenter_ipam_global_ip_address_pools_count" "example" {
  provider = dnacenter
}

output "dnacenter_ipam_global_ip_address_pools_count_example" {
  value = data.dnacenter_ipam_global_ip_address_pools_count.example.item
}
