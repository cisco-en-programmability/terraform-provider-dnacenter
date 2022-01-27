
data "dnacenter_endpoint_analytics_profiling_rules" "example" {
  provider        = dnacenter
  include_deleted = "false"
  limit           = "#"
  offset          = "#"
  order           = "string"
  rule_type       = "string"
  sort_by         = "string"
}

output "dnacenter_endpoint_analytics_profiling_rules_example" {
  value = data.dnacenter_endpoint_analytics_profiling_rules.example.items
}

data "dnacenter_endpoint_analytics_profiling_rules" "example" {
  provider = dnacenter
  rule_id  = "string"
}

output "dnacenter_endpoint_analytics_profiling_rules_example" {
  value = data.dnacenter_endpoint_analytics_profiling_rules.example.item
}
