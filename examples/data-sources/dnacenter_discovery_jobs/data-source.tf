
data "dnacenter_discovery_jobs" "example" {
  provider   = dnacenter
  ip_address = "string"
  limit      = 1
  name       = "string"
  offset     = 1
}

output "dnacenter_discovery_jobs_example" {
  value = data.dnacenter_discovery_jobs.example.items
}
