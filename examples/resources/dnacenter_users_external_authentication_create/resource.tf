
resource "dnacenter_users_external_authentication_create" "example" {
  provider = dnacenter
  parameters {

    enable = "false"
  }
}

output "dnacenter_users_external_authentication_create_example" {
  value = dnacenter_users_external_authentication_create.example
}