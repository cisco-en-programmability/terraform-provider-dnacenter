
resource "dnacenter_images_download" "example" {
  provider = dnacenter
  id       = "string"
  parameters {

  }
}

output "dnacenter_images_download_example" {
  value = dnacenter_images_download.example
}