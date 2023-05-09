
resource "dnacenter_configuration_template_version" "example" {
  provider = dnacenter
  parameters {
    # comments = "string"
    template_id = "string"
  }
}

output "dnacenter_configuration_template_version_example" {
  value = dnacenter_configuration_template_version.example
}