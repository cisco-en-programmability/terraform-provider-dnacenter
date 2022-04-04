
resource "dnacenter_swim_trigger_distribution" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_uuid = "string"
    image_uuid  = "string"
  }
}