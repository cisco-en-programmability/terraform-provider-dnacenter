
data "dnacenter_site_health" "example" {
  provider  = dnacenter
  limit     = 1
  offset    = 1
  site_type = "string"
  timestamp = "string"
}

output "dnacenter_site_health_example" {
  value = data.dnacenter_site_health.example.items
}
