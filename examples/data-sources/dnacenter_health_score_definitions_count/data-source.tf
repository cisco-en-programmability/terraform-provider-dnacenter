
data "dnacenter_health_score_definitions_count" "example" {
  provider                   = dnacenter
  device_type                = "string"
  id                         = "string"
  include_for_overall_health = "false"
  xca_lle_rid                = "string"
}

output "dnacenter_health_score_definitions_count_example" {
  value = data.dnacenter_health_score_definitions_count.example.item
}
