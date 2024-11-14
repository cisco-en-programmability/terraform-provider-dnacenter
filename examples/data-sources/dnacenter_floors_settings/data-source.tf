
data "dnacenter_floors_settings" "example" {
  provider = dnacenter
}

output "dnacenter_floors_settings_example" {
  value = data.dnacenter_floors_settings.example.item
}
