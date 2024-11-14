
resource "dnacenter_sites_device_credentials_apply" "example" {
  provider = dnacenter
  parameters {

    device_credential_id = "string"
    site_id              = "string"
  }
}

output "dnacenter_sites_device_credentials_apply_example" {
  value = dnacenter_sites_device_credentials_apply.example
}