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

data "dna_site_count" "response" {
  provider = dnacenter
  site_id  = "1ba55132-0fb8-4986-a31e-674d30e8b8ee"
}
output "dna_site_count_response" {
  value = data.dna_site_count.response
}

data "dna_site" "response" {
  provider = dnacenter
  offset   = "1"
  limit    = "3"
  name     = "Global/USA"
}
output "dna_site_response" {
  value = data.dna_site.response
}


data "dna_site_health" "response" {
  provider = dnacenter
}
output "dna_site_health_response" {
  value = data.dna_site_health.response
}
