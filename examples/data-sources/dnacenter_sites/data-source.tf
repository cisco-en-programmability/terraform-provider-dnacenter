
data "dnacenter_sites" "example" {
  provider         = dnacenter
  limit            = 1
  name             = "string"
  name_hierarchy   = "string"
  offset           = 1
  type             = "string"
  units_of_measure = "string"
}

output "dnacenter_sites_example" {
  value = data.dnacenter_sites.example.items
}
