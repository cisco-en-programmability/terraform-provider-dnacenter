
data "dnacenter_system_health" "example" {
  provider  = dnacenter
  domain    = "string"
  limit     = "#"
  offset    = "#"
  subdomain = "string"
  summary   = "false"
}

output "dnacenter_system_health_example" {
  value = data.dnacenter_system_health.example.item
}
