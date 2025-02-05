
resource "dnacenter_projects" "example" {
  provider = dnacenter

  parameters {

    description = "string"
    name        = "string"
  }
}

output "dnacenter_projects_example" {
  value = dnacenter_projects.example
}
