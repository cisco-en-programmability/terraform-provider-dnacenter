
resource "dnacenter_users_external_authentication" "example" {
  provider = dnacenter

}

output "dnacenter_users_external_authentication_example" {
  value = dnacenter_users_external_authentication.example
}