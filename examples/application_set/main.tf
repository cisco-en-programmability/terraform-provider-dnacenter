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

resource "dna_application_set" "response" {
  provider = dnacenter
  item {
    name = "test-set"
  }
}

output "dna_application_set_response" {
  value = dna_application_set.response
}


data "dna_application_set" "query" {
  provider = dnacenter
  # name     = "local-services"
}

output "dna_application_set_query" {
  value = data.dna_application_set.query
}
