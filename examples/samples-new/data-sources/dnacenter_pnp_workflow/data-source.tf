terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}
/*
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
*/
data "dnacenter_pnp_workflow" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_pnp_workflow_example" {
  value = data.dnacenter_pnp_workflow.example.item
}
