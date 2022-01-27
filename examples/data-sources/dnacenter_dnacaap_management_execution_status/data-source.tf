
data "dnacenter_dnacaap_management_execution_status" "example" {
  provider     = dnacenter
  execution_id = "string"
}

output "dnacenter_dnacaap_management_execution_status_example" {
  value = data.dnacenter_dnacaap_management_execution_status.example.item
}
