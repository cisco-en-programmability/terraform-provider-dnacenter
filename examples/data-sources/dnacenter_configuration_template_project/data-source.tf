
data "dnacenter_configuration_template_project" "example" {
  provider   = dnacenter
  name       = "string"
  sort_order = "string"
}

output "dnacenter_configuration_template_project_example" {
  value = data.dnacenter_configuration_template_project.example.items
}

data "dnacenter_configuration_template_project" "example" {
  provider   = dnacenter
  project_id = "string"
}

output "dnacenter_configuration_template_project_example" {
  value = data.dnacenter_configuration_template_project.example.item
}
