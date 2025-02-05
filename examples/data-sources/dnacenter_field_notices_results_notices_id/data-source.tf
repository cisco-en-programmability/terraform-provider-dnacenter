
data "dnacenter_field_notices_results_notices_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_field_notices_results_notices_id_example" {
  value = data.dnacenter_field_notices_results_notices_id.example.item
}
