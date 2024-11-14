
resource "dnacenter_maps_import_perform" "example" {
  provider            = dnacenter
  import_context_uuid = "string"
  parameters {

  }
}

output "dnacenter_maps_import_perform_example" {
  value = dnacenter_maps_import_perform.example
}