
data "dnacenter_site" "example" {
  provider = dnacenter
  limit    = 1
  name     = "string"
  offset   = 1
  site_id  = "string"
  type     = "string"
}

output "dnacenter_site_example" {
  value = data.dnacenter_site.example.items
}
