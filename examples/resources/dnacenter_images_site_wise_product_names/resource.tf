
resource "dnacenter_images_site_wise_product_names" "example" {
  provider = dnacenter

  parameters {

    image_id             = "string"
    product_name_ordinal = 1.0
    site_ids             = ["string"]
  }
}

output "dnacenter_images_site_wise_product_names_example" {
  value = dnacenter_images_site_wise_product_names.example
}