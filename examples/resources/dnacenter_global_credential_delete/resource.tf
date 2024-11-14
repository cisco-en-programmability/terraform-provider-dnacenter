
resource "dnacenter_global_credential_delete" "example" {
  provider             = dnacenter
  global_credential_id = "string"
  parameters {

  }
}

output "dnacenter_global_credential_delete_example" {
  value = dnacenter_global_credential_delete.example
}