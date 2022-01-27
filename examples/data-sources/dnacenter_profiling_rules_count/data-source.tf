
data "dnacenter_profiling_rules_count" "example" {
  provider        = dnacenter
  include_deleted = "false"
  rule_type       = "string"
}

output "dnacenter_profiling_rules_count_example" {
  value = data.dnacenter_profiling_rules_count.example.item
}
