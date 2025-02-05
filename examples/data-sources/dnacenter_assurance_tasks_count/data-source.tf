
data "dnacenter_assurance_tasks_count" "example" {
  provider    = dnacenter
  status      = "string"
  xca_lle_rid = "string"
}

output "dnacenter_assurance_tasks_count_example" {
  value = data.dnacenter_assurance_tasks_count.example.item
}
