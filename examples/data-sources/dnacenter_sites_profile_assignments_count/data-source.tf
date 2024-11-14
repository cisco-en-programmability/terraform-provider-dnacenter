
data "dnacenter_sites_profile_assignments_count" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_sites_profile_assignments_count_example" {
  value = data.dnacenter_sites_profile_assignments_count.example.item
}
