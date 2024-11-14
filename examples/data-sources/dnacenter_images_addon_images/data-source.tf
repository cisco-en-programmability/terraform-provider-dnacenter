
data "dnacenter_images_addon_images" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_images_addon_images_example" {
  value = data.dnacenter_images_addon_images.example.items
}
