
data "dnacenter_site_wise_product_names" "example" {
  provider     = dnacenter
  limit        = 1
  offset       = 1
  product_name = "string"
  site_id      = "string"
}

output "dnacenter_site_wise_product_names_example" {
  value = data.dnacenter_site_wise_product_names.example.item
}
