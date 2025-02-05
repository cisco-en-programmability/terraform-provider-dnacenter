
data "dnacenter_projects" "example" {
  provider = dnacenter
  limit    = 1
  name     = "string"
  offset   = 1
}

output "dnacenter_projects_example" {
  value = data.dnacenter_projects.example.items
}
