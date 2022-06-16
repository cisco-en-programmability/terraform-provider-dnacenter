
data "dnacenter_planned_access_points" "example" {
  provider = dnacenter
  floor_id = "string"
  limit    = 1
  offset   = 1
  radios   = "false"
}

output "dnacenter_planned_access_points_example" {
  value = data.dnacenter_planned_access_points.example.items
}
