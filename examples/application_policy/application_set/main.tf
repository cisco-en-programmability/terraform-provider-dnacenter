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
    name = var.application_set_name
  }
}

data "dna_application_set" "query" {
  provider   = dnacenter
  depends_on = [dna_application_set.response]
  name       = var.application_set_name
}
