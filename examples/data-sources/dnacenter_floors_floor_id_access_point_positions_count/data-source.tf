
data "dnacenter_floors_floor_id_access_point_positions_count" "example" {
  provider    = dnacenter
  floor_id    = "string"
  mac_address = "string"
  model       = "string"
  name        = "string"
  type        = "string"
}

output "dnacenter_floors_floor_id_access_point_positions_count_example" {
  value = data.dnacenter_floors_floor_id_access_point_positions_count.example.item
}
