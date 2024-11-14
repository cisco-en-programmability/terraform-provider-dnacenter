
data "dnacenter_images_site_wise_product_names" "example" {
  provider     = dnacenter
  assigned     = "string"
  image_id     = "string"
  limit        = 1
  offset       = 1
  product_id   = "string"
  product_name = "string"
  recommended  = "string"
}

output "dnacenter_images_site_wise_product_names_example" {
  value = data.dnacenter_images_site_wise_product_names.example.items
}
