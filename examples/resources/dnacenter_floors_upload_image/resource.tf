
resource "dnacenter_floors_upload_image" "example" {
  provider = dnacenter
  id       = "string"
  parameters {

  }
}

output "dnacenter_floors_upload_image_example" {
  value = dnacenter_floors_upload_image.example
}