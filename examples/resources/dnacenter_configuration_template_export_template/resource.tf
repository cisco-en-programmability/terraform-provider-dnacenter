
resource "dnacenter_configuration_template_export_template" "example" {
  provider = dnacenter
  parameters {
    payload = ["string"]
  }
}

output "dnacenter_configuration_template_export_template_example" {
  value = dnacenter_configuration_template_export_template.example
}