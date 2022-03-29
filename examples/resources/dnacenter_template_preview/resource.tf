
resource "dnacenter_template_preview" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_id       = "string"
    params          = ["string"]
    resource_params = ["string"]
    template_id     = "string"
  }
}