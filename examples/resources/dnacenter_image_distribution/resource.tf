
resource "dnacenter_image_distribution" "example" {
  provider = dnacenter
  parameters {
    payload {
      device_uuid = "string"
      image_uuid  = "string"
    }
  }
}

output "dnacenter_image_distribution_example" {
  value = dnacenter_image_distribution.example
}