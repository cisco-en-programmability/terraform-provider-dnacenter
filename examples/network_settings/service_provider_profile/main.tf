
terraform {
  required_providers {
    dnacenter = {
      versions = ["0.0.3"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

resource "dna_network_service_provider_profile" "response1" {
  provider     = dnacenter
  profile_name = "Test1"
  model        = "6-class-model"
  wan_provider = "test1-provider"
}
output "dna_network_service_provider_profile_response1" {
  value = dna_network_service_provider_profile.response1
}

data "dna_network_service_provider_profile" "response" {
  provider = dnacenter
}
output "dna_network_service_provider_profile_response" {
  value = data.dna_network_service_provider_profile.response
}
