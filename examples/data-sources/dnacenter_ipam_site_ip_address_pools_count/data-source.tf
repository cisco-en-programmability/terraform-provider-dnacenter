
data "dnacenter_ipam_site_ip_address_pools_count" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_ipam_site_ip_address_pools_count_example" {
  value = data.dnacenter_ipam_site_ip_address_pools_count.example.item
}
