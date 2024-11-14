
data "dnacenter_sites_image_distribution_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_image_distribution_settings_example" {
  value = data.dnacenter_sites_image_distribution_settings.example.item
}
