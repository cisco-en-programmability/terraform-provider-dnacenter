
resource "dnacenter_configuration_template_project" "example" {
  provider = dnacenter

  parameters {

    create_time      = 1
    description      = "string"
    id               = "string"
    last_update_time = 1
    name             = "string"
    project_id       = "string"
    tags {

      id   = "string"
      name = "string"
    }
    templates = "string"
  }
}

output "dnacenter_configuration_template_project_example" {
  value = dnacenter_configuration_template_project.example
}