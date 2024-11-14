
data "dnacenter_sites_device_credentials" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_device_credentials_example" {
  value = data.dnacenter_sites_device_credentials.example.item
}
