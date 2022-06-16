
data "dnacenter_configuration_template_version" "example" {
  provider    = dnacenter
  template_id = "string"
}

output "dnacenter_configuration_template_version_example" {
  value = data.dnacenter_configuration_template_version.example.items
}
