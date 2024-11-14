
data "dnacenter_sda_authentication_profiles" "example" {
  provider                    = dnacenter
  authentication_profile_name = "string"
  fabric_id                   = "string"
  limit                       = 1
  offset                      = 1
}

output "dnacenter_sda_authentication_profiles_example" {
  value = data.dnacenter_sda_authentication_profiles.example.items
}
