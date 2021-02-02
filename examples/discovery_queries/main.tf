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

data "dna_discovery_range" "response" {
  provider          = dnacenter
  start_index       = 1
  records_to_return = 2
}

output "dna_discovery_range_response" {
  value = data.dna_discovery_range.response.items
}

data "dna_discovery_count" "response" {
  provider   = dnacenter
  depends_on = [data.dna_discovery_range.response]
  id         = data.dna_discovery_range.response.items.0.id
}

output "dna_discovery_count_response" {
  value = data.dna_discovery_count.response.response
}

data "dna_discovery_device" "response" {
  provider   = dnacenter
  depends_on = [data.dna_discovery_range.response]
  id         = data.dna_discovery_range.response.items.0.id
}

output "dna_discovery_device_response" {
  value = data.dna_discovery_device.response
}

data "dna_discovery_device_range" "response" {
  provider          = dnacenter
  depends_on        = [data.dna_discovery_range.response]
  id                = data.dna_discovery_range.response.items.0.id
  start_index       = 1
  records_to_return = 4
}

output "dna_discovery_device_range_response" {
  value = data.dna_discovery_device_range.response
}


data "dna_discovery_snmp_property" "response" {
  provider = dnacenter
}

output "dna_discovery_snmp_property_response" {
  value = data.dna_discovery_snmp_property.response
}


data "dna_discovery_summary" "response" {
  provider   = dnacenter
  depends_on = [data.dna_discovery_range.response]
  id         = data.dna_discovery_range.response.items.0.id
}

output "dna_discovery_summary_response" {
  value = data.dna_discovery_summary.response
}

data "dna_discovery_job" "response" {
  provider   = dnacenter
  depends_on = [data.dna_discovery_range.response]
  id         = data.dna_discovery_range.response.items.0.id
}

output "dna_discovery_job_response" {
  value = data.dna_discovery_job.response
}
