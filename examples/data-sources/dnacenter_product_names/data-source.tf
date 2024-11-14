
data "dnacenter_product_names" "example" {
  provider     = dnacenter
  limit        = 1
  offset       = 1
  product_id   = "string"
  product_name = "string"
}

output "dnacenter_product_names_example" {
  value = data.dnacenter_product_names.example.items
}

data "dnacenter_product_names" "example" {
  provider             = dnacenter
  product_name_ordinal = 1.0
}

output "dnacenter_product_names_example" {
  value = data.dnacenter_product_names.example.item
}
