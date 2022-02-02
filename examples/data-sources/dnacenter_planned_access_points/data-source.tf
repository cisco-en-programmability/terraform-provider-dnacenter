
data "dnacenter_planned_access_points" "example" {
  provider = dnacenter
  floor_id = "string"
  limit    = 0.0
  offset   = 0.0
  radios   = "false"
}

output "dnacenter_planned_access_points_example" {
  value = data.dnacenter_planned_access_points.example.items
}
