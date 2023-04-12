
resource "dnacenter_golden_tag_image" "example" {
  provider = dnacenter
  parameters {

    device_family_identifier = "string"
    device_role              = "string"
    image_id                 = "string"
    site_id                  = "string"
  }
}

output "dnacenter_golden_tag_image_example" {
  value = dnacenter_golden_tag_image.example
}