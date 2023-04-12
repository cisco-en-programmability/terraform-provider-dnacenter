
data "dnacenter_golden_tag_image" "example" {
  provider                 = dnacenter
  device_family_identifier = "string"
  device_role              = "string"
  image_id                 = "string"
  site_id                  = "string"
}

output "dnacenter_golden_tag_image_example" {
  value = data.dnacenter_golden_tag_image.example.item
}
