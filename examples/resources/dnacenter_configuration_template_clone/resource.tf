
resource "dnacenter_configuration_template_clone" "example" {
  provider = dnacenter
  parameters {

    name        = "string"
    project_id  = "string"
    template_id = "string"
  }
}

output "dnacenter_configuration_template_clone_example" {
  value = dnacenter_configuration_template_clone.example
}