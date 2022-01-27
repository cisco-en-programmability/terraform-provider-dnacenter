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

data "dnacenter_pnp_workflow_count" "example" {
  provider = dnacenter
  name     = ["string"]
}

output "dnacenter_pnp_workflow_count_example" {
  value = data.dnacenter_pnp_workflow_count.example.item
}
