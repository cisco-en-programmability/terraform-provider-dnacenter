
data "dnacenter_projects_details" "example" {
  provider   = dnacenter
  id         = "string"
  limit      = 1
  name       = "string"
  offset     = 1
  sort_order = "string"
}

output "dnacenter_projects_details_example" {
  value = data.dnacenter_projects_details.example.items
}
