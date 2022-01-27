
data "dnacenter_license_smart_account_details" "example" {
  provider = dnacenter
}

output "dnacenter_license_smart_account_details_example" {
  value = data.dnacenter_license_smart_account_details.example.items
}
