
data "dnacenter_assurance_tasks_id" "example" {
  provider    = dnacenter
  id          = "string"
  xca_lle_rid = "string"
}

output "dnacenter_assurance_tasks_id_example" {
  value = data.dnacenter_assurance_tasks_id.example.item
}
