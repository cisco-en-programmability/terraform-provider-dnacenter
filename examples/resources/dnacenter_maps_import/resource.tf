
resource "dnacenter_maps_import" "example" {
  provider            = dnacenter
  import_context_uuid = "string"
  parameters {

  }
}

output "dnacenter_maps_import_example" {
  value = dnacenter_maps_import.example
}