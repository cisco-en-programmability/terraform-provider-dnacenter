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

resource "dna_application_set" "response" {
  provider = dnacenter
  item {
    name = var.application_set_name
  }
}

data "dna_application_set" "query" {
  provider   = dnacenter
  depends_on = [dna_application_set.response]
  name       = dna_application_set.response.item[0].name
}
