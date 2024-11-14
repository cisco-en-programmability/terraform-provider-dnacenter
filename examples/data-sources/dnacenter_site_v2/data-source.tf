
data "dnacenter_site_v2" "example" {
  provider             = dnacenter
  group_name_hierarchy = "string"
  id                   = "string"
  limit                = "string"
  offset               = "string"
  type                 = "string"
}

output "dnacenter_site_v2_example" {
  value = data.dnacenter_site_v2.example.items
}
