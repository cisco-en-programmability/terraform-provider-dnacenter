
data "dnacenter_pnp_smart_account_domains" "example" {
  provider = dnacenter
}

output "dnacenter_pnp_smart_account_domains_example" {
  value = data.dnacenter_pnp_smart_account_domains.example.items
}
