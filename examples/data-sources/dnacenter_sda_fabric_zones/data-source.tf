
data "dnacenter_sda_fabric_zones" "example" {
  provider = dnacenter
  id       = "string"
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_sda_fabric_zones_example" {
  value = data.dnacenter_sda_fabric_zones.example.items
}
