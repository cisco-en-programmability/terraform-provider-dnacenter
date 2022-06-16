
data "dnacenter_device_enrichment_details" "example" {
  provider     = dnacenter
  entity_type  = "string"
  entity_value = "string"
}

output "dnacenter_device_enrichment_details_example" {
  value = data.dnacenter_device_enrichment_details.example.items
}
