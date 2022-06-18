
data "dnacenter_pnp_virtual_account_sync_result" "example" {
  provider = dnacenter
  domain   = "string"
  name     = "string"
}

output "dnacenter_pnp_virtual_account_sync_result_example" {
  value = data.dnacenter_pnp_virtual_account_sync_result.example.item
}
