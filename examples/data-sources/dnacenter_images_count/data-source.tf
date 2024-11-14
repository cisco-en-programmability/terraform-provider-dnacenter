
data "dnacenter_images_count" "example" {
  provider                        = dnacenter
  golden                          = "string"
  has_addon_images                = "false"
  imported                        = "false"
  integrity                       = "string"
  is_addon_images                 = "false"
  name                            = "string"
  product_name_ordinal            = 1.0
  site_id                         = "string"
  supervisor_product_name_ordinal = 1.0
  version                         = "string"
}

output "dnacenter_images_count_example" {
  value = data.dnacenter_images_count.example.item
}
