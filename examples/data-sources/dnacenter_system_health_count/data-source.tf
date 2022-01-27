
data "dnacenter_system_health_count" "example" {
  provider  = dnacenter
  domain    = "string"
  subdomain = "string"
}

output "dnacenter_system_health_count_example" {
  value = data.dnacenter_system_health_count.example.item
}
