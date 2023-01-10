
data "dnacenter_nfv_profile" "example" {
  provider = dnacenter
  id       = "string"
  limit    = 1
  name     = "string"
  offset   = 1
}

output "dnacenter_nfv_profile_example" {
  value = data.dnacenter_nfv_profile.example.items
}
