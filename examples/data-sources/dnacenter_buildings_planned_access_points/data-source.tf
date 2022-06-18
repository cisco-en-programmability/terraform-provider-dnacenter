
data "dnacenter_buildings_planned_access_points" "example" {
  provider    = dnacenter
  building_id = "string"
  limit       = 1
  offset      = 1
  radios      = "false"
}

output "dnacenter_buildings_planned_access_points_example" {
  value = data.dnacenter_buildings_planned_access_points.example.items
}
