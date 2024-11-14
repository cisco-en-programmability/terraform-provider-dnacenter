
resource "dnacenter_license_setting" "example" {
  provider = dnacenter

  parameters {

    auto_registration_virtual_account_id = "string"
    default_smart_account_id             = "string"
  }
}

output "dnacenter_license_setting_example" {
  value = dnacenter_license_setting.example
}