
data "dnacenter_global_pool" "example" {
  provider = dnacenter
  limit    = "string"
  offset   = "string"
}

output "dnacenter_global_pool_example" {
  value = data.dnacenter_global_pool.example.items
}
