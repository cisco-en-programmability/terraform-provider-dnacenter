
terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

resource "dna_pnp_workflow" "response" {
  provider = dnacenter
  item {
    name = "Workflow 1"
    tasks {
      name        = "Workflow 1 Task 1"
      task_seq_no = 0
      type        = "Reload"
    }
    tasks {
      name        = "Workflow 1 Task 2"
      task_seq_no = 1
      type        = "Reload"
    }
  }
}
output "dna_pnp_workflow_response" {
  value = dna_pnp_workflow.response
}

