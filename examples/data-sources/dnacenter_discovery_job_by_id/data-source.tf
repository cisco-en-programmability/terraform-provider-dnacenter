
data "dnacenter_discovery_job_by_id" "example" {
  provider   = dnacenter
  id         = "string"
  ip_address = "string"
  limit      = 1
  offset     = 1
}

output "dnacenter_discovery_job_by_id_example" {
  value = data.dnacenter_discovery_job_by_id.example.items
}
