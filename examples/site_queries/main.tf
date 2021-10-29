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

data "dna_site_count" "response" {
  provider = dnacenter
  site_id  = "2397da83-4e12-4d04-9bd3-a57b2ad91652"
}
output "dna_site_count_response" {
  value = data.dna_site_count.response
}

data "dna_site" "response" {
  provider = dnacenter
  offset   = "1"
  limit    = "3"
  name     = "Global/Peru"
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


data "dna_site_membership" "response" {
  provider = dnacenter
  site_id  = "2397da83-4e12-4d04-9bd3-a57b2ad91652"
}
output "dna_site_membership_response" {
  value = data.dna_site_membership.response
}
