
data "dnacenter_templates_template_id_versions_count" "example" {
  provider       = dnacenter
  latest_version = "false"
  template_id    = "string"
  version_number = 1
}

output "dnacenter_templates_template_id_versions_count_example" {
  value = data.dnacenter_templates_template_id_versions_count.example.item
}
