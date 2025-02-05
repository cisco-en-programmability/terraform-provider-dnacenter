
data "dnacenter_projects_project_id" "example" {
  provider   = dnacenter
  project_id = "string"
}

output "dnacenter_projects_project_id_example" {
  value = data.dnacenter_projects_project_id.example.item
}
