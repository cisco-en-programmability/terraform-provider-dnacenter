
resource "dnacenter_configuration_template_import_project" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    do_version = "false"
  }
}