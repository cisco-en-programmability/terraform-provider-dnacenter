
resource "dnacenter_maps_import_start" "example" {
  provider = dnacenter
}

output "dnacenter_maps_import_start_example" {
  value = dnacenter_maps_import_start.example
}