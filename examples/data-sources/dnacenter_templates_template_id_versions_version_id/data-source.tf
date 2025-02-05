
data "dnacenter_templates_template_id_versions_version_id" "example" {
  provider    = dnacenter
  template_id = "string"
  version_id  = "string"
}

output "dnacenter_templates_template_id_versions_version_id_example" {
  value = data.dnacenter_templates_template_id_versions_version_id.example.item
}
