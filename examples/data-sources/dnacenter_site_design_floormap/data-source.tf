
data "dnacenter_site_design_floormap" "example" {
  provider = dnacenter
  floor_id = "string"
}

output "dnacenter_site_design_floormap_example" {
  value = data.dnacenter_site_design_floormap.example.item
}
