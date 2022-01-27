
data "dnacenter_golden_tag_image_details" "example" {
  provider                 = dnacenter
  device_family_identifier = "string"
  device_role              = "string"
  image_id                 = "string"
  site_id                  = "string"
}

output "dnacenter_golden_tag_image_details_example" {
  value = data.dnacenter_golden_tag_image_details.example.item
}
