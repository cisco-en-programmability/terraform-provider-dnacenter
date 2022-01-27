
data "dnacenter_client_enrichment_details" "example" {
  provider       = dnacenter
  entity_type    = "string"
  entity_value   = "string"
  issue_category = "string"
}

output "dnacenter_client_enrichment_details_example" {
  value = data.dnacenter_client_enrichment_details.example.items
}
