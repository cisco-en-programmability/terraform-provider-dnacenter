
data "dnacenter_license_setting" "example" {
  provider = dnacenter
}

output "dnacenter_license_setting_example" {
  value = data.dnacenter_license_setting.example.item
}
