
resource "dnacenter_template_preview" "example" {
  provider = dnacenter
  parameters {

    device_id       = "string"
    params          = "string"
    resource_params = "string"
    template_id     = "string"
  }
}

output "dnacenter_template_preview_example" {
  value = dnacenter_template_preview.example
}