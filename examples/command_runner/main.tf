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

data "dna_command_runner_keywords" "response" {
  provider = dnacenter
}
output "dna_command_runner_keywords_response" {
  value = data.dna_command_runner_keywords.response
}
