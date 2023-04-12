
data "dnacenter_role_permissions" "example" {
  provider = dnacenter
}

output "dnacenter_role_permissions_example" {
  value = data.dnacenter_role_permissions.example.item
}
