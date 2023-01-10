
data "dnacenter_network_device_inventory_insight_link_mismatch" "example" {
  provider = dnacenter
  category = "string"
  limit    = 1
  offset   = 1
  order    = "string"
  site_id  = "string"
  sort_by  = "string"
}

output "dnacenter_network_device_inventory_insight_link_mismatch_example" {
  value = data.dnacenter_network_device_inventory_insight_link_mismatch.example.items
}
