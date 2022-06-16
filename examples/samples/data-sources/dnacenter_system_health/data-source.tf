
data "dnacenter_system_health" "example" {
  provider  = dnacenter
  domain    = "string"
  limit     = 1
  offset    = 1
  subdomain = "string"
  summary   = "false"
}

output "dnacenter_system_health_example" {
  value = data.dnacenter_system_health.example.item
}
