
data "dnacenter_sites_device_credentials_status" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_sites_device_credentials_status_example" {
  value = data.dnacenter_sites_device_credentials_status.example.item
}
