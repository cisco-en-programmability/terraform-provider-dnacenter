
data "dnacenter_users_external_authentication" "example" {
  provider = dnacenter
}

output "dnacenter_users_external_authentication_example" {
  value = data.dnacenter_users_external_authentication.example.item
}
