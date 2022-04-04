
resource "dnacenter_configuration_template_export_project" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload = ["string"]
  }
}