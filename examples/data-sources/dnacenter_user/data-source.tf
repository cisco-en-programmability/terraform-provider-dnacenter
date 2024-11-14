
data "dnacenter_user" "example" {
  provider      = dnacenter
  auth_source   = "string"
  invoke_source = "string"
}

output "dnacenter_user_example" {
  value = data.dnacenter_user.example.item
}
