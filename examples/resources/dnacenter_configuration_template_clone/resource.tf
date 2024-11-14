
resource "dnacenter_configuration_template_clone" "example" {
  provider    = dnacenter
  name        = "string"
  project_id  = "string"
  template_id = "string"
  parameters {

  }
}

output "dnacenter_configuration_template_clone_example" {
  value = dnacenter_configuration_template_clone.example
}