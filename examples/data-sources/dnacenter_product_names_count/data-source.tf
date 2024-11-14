
data "dnacenter_product_names_count" "example" {
  provider     = dnacenter
  product_id   = "string"
  product_name = "string"
}

output "dnacenter_product_names_count_example" {
  value = data.dnacenter_product_names_count.example.item
}
