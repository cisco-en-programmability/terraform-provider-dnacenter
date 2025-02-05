
resource "dnacenter_wireless_settings_power_profiles_id" "example" {
  provider = dnacenter

  parameters {

    description  = "string"
    id           = "string"
    profile_name = "string"
    rules {

      interface_id    = "string"
      interface_type  = "string"
      parameter_type  = "string"
      parameter_value = "string"
    }
  }
}

output "dnacenter_wireless_settings_power_profiles_id_example" {
  value = dnacenter_wireless_settings_power_profiles_id.example
}
