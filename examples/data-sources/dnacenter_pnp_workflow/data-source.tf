
data "dnacenter_pnp_workflow" "example" {
  provider   = dnacenter
  limit      = 1
  name       = ["string"]
  offset     = 1
  sort       = ["string"]
  sort_order = "string"
  type       = ["string"]
}

output "dnacenter_pnp_workflow_example" {
  value = data.dnacenter_pnp_workflow.example.items
}

data "dnacenter_pnp_workflow" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_pnp_workflow_example" {
  value = data.dnacenter_pnp_workflow.example.item
}
