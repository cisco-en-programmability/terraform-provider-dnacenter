
data "dnacenter_sites_profile_assignments" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_sites_profile_assignments_example" {
  value = data.dnacenter_sites_profile_assignments.example.items
}
