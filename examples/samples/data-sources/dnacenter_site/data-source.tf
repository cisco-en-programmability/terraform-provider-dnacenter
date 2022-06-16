
data "dnacenter_site" "example" {
  provider = dnacenter
  limit    = "string"
  name     = "string"
  offset   = "string"
  site_id  = "string"
  type     = "string"
}

output "dnacenter_site_example" {
  value = data.dnacenter_site.example.items
}
