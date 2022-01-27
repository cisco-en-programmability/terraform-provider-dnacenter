
data "dnacenter_nfv_profile" "example" {
  provider = dnacenter
  id       = "string"
  limit    = "string"
  name     = "string"
  offset   = "string"
}

output "dnacenter_nfv_profile_example" {
  value = data.dnacenter_nfv_profile.example.items
}
