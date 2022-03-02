
resource "dnacenter_swim_image_file" "example" {
  provider = dnacenter
  parameters {
    file_path = "string"
    file_name = "string"
  }
}

output "dnacenter_swim_image_file_example" {
  value = dnacenter_swim_image_file.example
}