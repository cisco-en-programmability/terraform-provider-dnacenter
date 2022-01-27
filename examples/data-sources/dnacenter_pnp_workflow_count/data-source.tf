
data "dnacenter_pnp_workflow_count" "example" {
  provider = dnacenter
  name     = ["string"]
}

output "dnacenter_pnp_workflow_count_example" {
  value = data.dnacenter_pnp_workflow_count.example.item
}
