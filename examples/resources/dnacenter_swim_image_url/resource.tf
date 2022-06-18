
resource "dnacenter_swim_image_url" "example" {
  provider = dnacenter
  parameters {

    application_type = "string"
    image_family     = "string"
    source_url       = "string"
    third_party      = "false"
    vendor           = "string"
  }
}

output "dnacenter_swim_image_url_example" {
  value = dnacenter_swim_image_url.example
}