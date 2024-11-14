
data "dnacenter_provisioning_settings" "example" {
  provider = dnacenter
}

output "dnacenter_provisioning_settings_example" {
  value = data.dnacenter_provisioning_settings.example.item
}
