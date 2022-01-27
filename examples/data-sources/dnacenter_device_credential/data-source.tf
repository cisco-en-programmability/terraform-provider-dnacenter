
data "dnacenter_device_credential" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_device_credential_example" {
  value = data.dnacenter_device_credential.example.item
}
