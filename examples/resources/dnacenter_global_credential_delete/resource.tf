
resource "dnacenter_global_credential_delete" "example" {
  provider = dnacenter
  parameters {

    global_credential_id = "string"
  }
}

output "dnacenter_global_credential_delete_example" {
  value = dnacenter_global_credential_delete.example
}