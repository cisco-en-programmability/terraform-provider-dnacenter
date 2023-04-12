
data "dnacenter_users_external_servers" "example" {
  provider      = dnacenter
  invoke_source = "string"
}

output "dnacenter_users_external_servers_example" {
  value = data.dnacenter_users_external_servers.example.item
}
