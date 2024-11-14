
resource "dnacenter_configuration_template_import_project" "example" {
  provider   = dnacenter
  do_version = "false"
}

output "dnacenter_configuration_template_import_project_example" {
  value = dnacenter_configuration_template_import_project.example
}