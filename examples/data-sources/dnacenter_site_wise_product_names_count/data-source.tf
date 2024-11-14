
data "dnacenter_site_wise_product_names_count" "example" {
  provider     = dnacenter
  product_name = "string"
  site_id      = "string"
}

output "dnacenter_site_wise_product_names_count_example" {
  value = data.dnacenter_site_wise_product_names_count.example.item
}
