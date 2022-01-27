
data "dnacenter_pnp_virtual_accounts" "example" {
  provider = dnacenter
  domain   = "string"
}

output "dnacenter_pnp_virtual_accounts_example" {
  value = data.dnacenter_pnp_virtual_accounts.example.items
}
