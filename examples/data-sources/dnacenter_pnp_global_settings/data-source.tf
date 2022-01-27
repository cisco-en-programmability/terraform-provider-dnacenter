
data "dnacenter_pnp_global_settings" "example" {
  provider = dnacenter
}

output "dnacenter_pnp_global_settings_example" {
  value = data.dnacenter_pnp_global_settings.example.item
}
