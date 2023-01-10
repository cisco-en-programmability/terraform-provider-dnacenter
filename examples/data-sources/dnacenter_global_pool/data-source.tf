
data "dnacenter_global_pool" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_global_pool_example" {
  value = data.dnacenter_global_pool.example.items
}
