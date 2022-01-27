
data "dnacenter_wireless_profile" "example" {
  provider     = dnacenter
  profile_name = "string"
}

output "dnacenter_wireless_profile_example" {
  value = data.dnacenter_wireless_profile.example.items
}
