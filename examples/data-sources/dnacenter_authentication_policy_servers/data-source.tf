
data "dnacenter_authentication_policy_servers" "example" {
  provider       = dnacenter
  is_ise_enabled = "false"
  role           = "string"
  state          = "string"
}

output "dnacenter_authentication_policy_servers_example" {
  value = data.dnacenter_authentication_policy_servers.example.items
}
