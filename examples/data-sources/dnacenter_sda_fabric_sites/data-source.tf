
data "dnacenter_sda_fabric_sites" "example" {
  provider = dnacenter
  id       = "string"
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_sda_fabric_sites_example" {
  value = data.dnacenter_sda_fabric_sites.example.items
}
