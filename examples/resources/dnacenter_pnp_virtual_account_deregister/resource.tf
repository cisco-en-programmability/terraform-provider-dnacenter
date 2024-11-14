
resource "dnacenter_pnp_virtual_account_deregister" "example" {
  provider = dnacenter
  domain   = "string"
  name     = "string"
}

output "dnacenter_pnp_virtual_account_deregister_example" {
  value = dnacenter_pnp_virtual_account_deregister.example
}