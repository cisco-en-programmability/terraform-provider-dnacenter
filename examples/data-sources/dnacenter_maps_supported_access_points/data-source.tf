
data "dnacenter_maps_supported_access_points" "example" {
  provider = dnacenter
}

output "dnacenter_maps_supported_access_points_example" {
  value = data.dnacenter_maps_supported_access_points.example.items
}
