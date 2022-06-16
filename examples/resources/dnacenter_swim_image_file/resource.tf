
resource "dnacenter_swim_image_file" "example" {
  provider = dnacenter
}

output "dnacenter_swim_image_file_example" {
  value = dnacenter_swim_image_file.example
}