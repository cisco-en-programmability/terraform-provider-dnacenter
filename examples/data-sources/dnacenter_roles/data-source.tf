
data "dnacenter_roles" "example" {
  provider      = dnacenter
  invoke_source = "string"
}

output "dnacenter_roles_example" {
  value = data.dnacenter_roles.example.item
}
