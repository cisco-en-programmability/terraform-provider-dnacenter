
data "dnacenter_images_site_wise_product_names_count" "example" {
  provider     = dnacenter
  assigned     = "string"
  image_id     = "string"
  product_id   = "string"
  product_name = "string"
  recommended  = "string"
}

output "dnacenter_images_site_wise_product_names_count_example" {
  value = data.dnacenter_images_site_wise_product_names_count.example.item
}
