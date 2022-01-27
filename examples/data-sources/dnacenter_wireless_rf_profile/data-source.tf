
data "dnacenter_wireless_rf_profile" "example" {
  provider        = dnacenter
  rf_profile_name = "string"
}

output "dnacenter_wireless_rf_profile_example" {
  value = data.dnacenter_wireless_rf_profile.example.items
}
