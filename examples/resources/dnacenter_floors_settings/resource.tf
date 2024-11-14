
resource "dnacenter_floors_settings" "example" {
  provider = dnacenter

  parameters {

    units_of_measure = "string"
  }
}

output "dnacenter_floors_settings_example" {
  value = dnacenter_floors_settings.example
}