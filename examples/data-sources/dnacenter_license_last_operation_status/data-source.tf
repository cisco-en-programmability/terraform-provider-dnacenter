
data "dnacenter_license_last_operation_status" "example" {
  provider = dnacenter
}

output "dnacenter_license_last_operation_status_example" {
  value = data.dnacenter_license_last_operation_status.example.item
}
