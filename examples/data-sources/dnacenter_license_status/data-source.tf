
data "dnacenter_license_status" "example" {
  provider = dnacenter
}

output "dnacenter_license_status_example" {
  value = data.dnacenter_license_status.example.item
}
