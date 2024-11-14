
data "dnacenter_ise_integration_status" "example" {
  provider = dnacenter
}

output "dnacenter_ise_integration_status_example" {
  value = data.dnacenter_ise_integration_status.example.item
}
