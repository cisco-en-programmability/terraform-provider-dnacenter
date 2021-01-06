
terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

resource "dna_network_credential_site_assignment" "response1" {
  provider = dnacenter
  site_id  = "b85fe0be-d971-4eb7-92a9-3498356ad87f"
  http_read {
    id = "babc42b9-0bdd-49ef-912e-66f533fb5d59"
  }
  cli {
    id = "f979d842-f6fd-456a-8137-2cb5113cd2e8"
  }
}
output "dna_network_credential_site_assignment_response1" {
  value = dna_network_credential_site_assignment.response1
}
