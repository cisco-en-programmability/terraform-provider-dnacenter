
data "dnacenter_reserve_ip_subpool" "example" {
  provider                = dnacenter
  group_name              = "string"
  ignore_inherited_groups = "string"
  limit                   = 1
  offset                  = 1
  pool_usage              = "string"
  site_id                 = "string"
}

output "dnacenter_reserve_ip_subpool_example" {
  value = data.dnacenter_reserve_ip_subpool.example.items
}
