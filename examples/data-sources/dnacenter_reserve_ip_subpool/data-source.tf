
data "dnacenter_reserve_ip_subpool" "example" {
  provider = dnacenter
  limit    = "string"
  offset   = "string"
  site_id  = "string"
}

output "dnacenter_reserve_ip_subpool_example" {
  value = data.dnacenter_reserve_ip_subpool.example.items
}
