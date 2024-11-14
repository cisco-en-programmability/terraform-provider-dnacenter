
data "dnacenter_maps_import_status" "example" {
  provider            = dnacenter
  import_context_uuid = "string"
}

output "dnacenter_maps_import_status_example" {
  value = data.dnacenter_maps_import_status.example.item
}
