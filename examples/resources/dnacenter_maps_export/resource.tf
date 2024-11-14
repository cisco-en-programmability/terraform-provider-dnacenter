
resource "dnacenter_maps_export" "example" {
  provider            = dnacenter
  site_hierarchy_uuid = "string"
  parameters {

  }
}

output "dnacenter_maps_export_example" {
  value = dnacenter_maps_export.example
}