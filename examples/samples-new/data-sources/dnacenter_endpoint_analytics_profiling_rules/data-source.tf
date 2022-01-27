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
/*
data "dnacenter_endpoint_analytics_profiling_rules" "example" {
  provider        = dnacenter
  include_deleted = "false"
  limit           = 1
  offset          = 1
  order           = "string"
  rule_type       = "string"
  sort_by         = "string"
}

output "dnacenter_endpoint_analytics_profiling_rules_example" {
  value = data.dnacenter_endpoint_analytics_profiling_rules.example.items
}

*/
data "dnacenter_endpoint_analytics_profiling_rules" "example" {
  provider = dnacenter
  rule_id  = "string"
}

output "dnacenter_endpoint_analytics_profiling_rules_example" {
  value = data.dnacenter_endpoint_analytics_profiling_rules.example.item
}
