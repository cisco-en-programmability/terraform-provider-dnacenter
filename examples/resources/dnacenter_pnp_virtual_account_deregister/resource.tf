
resource "dnacenter_pnp_virtual_account_deregister" "example" {
  provider = dnacenter
}

output "dnacenter_pnp_virtual_account_deregister_example" {
  value = dnacenter_pnp_virtual_account_deregister.example
}