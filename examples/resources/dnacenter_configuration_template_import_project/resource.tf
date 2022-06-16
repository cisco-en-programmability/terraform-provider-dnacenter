
resource "dnacenter_configuration_template_import_project" "example" {
  provider = dnacenter
}

output "dnacenter_configuration_template_import_project_example" {
  value = dnacenter_configuration_template_import_project.example
}