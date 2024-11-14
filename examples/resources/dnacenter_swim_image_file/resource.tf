
resource "dnacenter_swim_image_file" "example" {
  provider                     = dnacenter
  file_name                    = "string"
  file_path                    = "string"
  is_third_party               = "false"
  third_party_application_type = "string"
  third_party_image_family     = "string"
  third_party_vendor           = "string"
}

output "dnacenter_swim_image_file_example" {
  value = dnacenter_swim_image_file.example
}