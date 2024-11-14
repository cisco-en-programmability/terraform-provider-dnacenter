
resource "dnacenter_auth_token_create" "example" {
  provider      = dnacenter
  authorization = "string"
}

output "dnacenter_auth_token_create_example" {
  value = dnacenter_auth_token_create.example
}