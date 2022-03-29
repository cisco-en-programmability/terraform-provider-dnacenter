
resource "dnacenter_configuration_template_clone" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    name        = "string"
    project_id  = "string"
    template_id = "string"
  }
}