
data "dnacenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools_count" "example" {
  provider                  = dnacenter
  global_ip_address_pool_id = "string"
}

output "dnacenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools_count_example" {
  value = data.dnacenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools_count.example.item
}
