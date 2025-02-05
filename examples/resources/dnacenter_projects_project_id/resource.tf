
resource "dnacenter_projects_project_id" "example" {
  provider = dnacenter

  parameters {

    description = "string"
    name        = "string"
    project_id  = "string"
  }
}

output "dnacenter_projects_project_id_example" {
  value = dnacenter_projects_project_id.example
}
