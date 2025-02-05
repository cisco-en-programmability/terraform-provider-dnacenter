
data "dnacenter_ipam_site_ip_address_pools_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_ipam_site_ip_address_pools_id_example" {
  value = data.dnacenter_ipam_site_ip_address_pools_id.example.item
}
