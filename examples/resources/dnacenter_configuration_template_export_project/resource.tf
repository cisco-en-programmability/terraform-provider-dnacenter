
resource "dnacenter_configuration_template_export_project" "example" {
  provider   = dnacenter
  parameters = ["string"]
}

output "dnacenter_configuration_template_export_project_example" {
  value = dnacenter_configuration_template_export_project.example
}