
data "dnacenter_assurance_tasks" "example" {
  provider    = dnacenter
  limit       = 1
  offset      = 1
  order       = "string"
  sort_by     = "string"
  status      = "string"
  xca_lle_rid = "string"
}

output "dnacenter_assurance_tasks_example" {
  value = data.dnacenter_assurance_tasks.example.items
}
