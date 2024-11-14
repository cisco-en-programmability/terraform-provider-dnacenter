
data "dnacenter_images_addon_images_count" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_images_addon_images_count_example" {
  value = data.dnacenter_images_addon_images_count.example.item
}
