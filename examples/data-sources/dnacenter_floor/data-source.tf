
data "dnacenter_floor" "example" {
  provider = dnacenter
  limit    = 1
  name     = "string"
  offset   = 1
  site_id  = "string"
  type     = "string"
}

output "dnacenter_floor_example" {
  value = data.dnacenter_floor.example.items
}
