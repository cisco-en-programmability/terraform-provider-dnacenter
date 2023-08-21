
data "dnacenter_area" "example" {
  provider = dnacenter
  limit    = 1
  name     = "string"
  offset   = 1
  site_id  = "string"
  type     = "string"
}

output "dnacenter_area_example" {
  value = data.dnacenter_area.example.items
}
