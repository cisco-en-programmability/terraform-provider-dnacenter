
resource "dnacenter_global_credential_update" "example" {
  provider = dnacenter
  parameters {

    global_credential_id = "string"
    site_uuids           = ["string"]
  }
}

output "dnacenter_global_credential_update_example" {
  value = dnacenter_global_credential_update.example
}