provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_golden_image" "example" {
  provider = dnacenter
  parameters {
    image_id                 = "string"
    site_id                  = "string"
    device_role              = "string"
    device_family_identifier = "string"
  }
}

output "dnacenter_golden_image_example" {
  value = dnacenter_golden_image.example
}