
resource "dnacenter_global_credential_update" "example" {
  provider             = dnacenter
  global_credential_id = "string"
  parameters {

    site_uuids = ["string"]
  }
}

output "dnacenter_global_credential_update_example" {
  value = dnacenter_global_credential_update.example
}