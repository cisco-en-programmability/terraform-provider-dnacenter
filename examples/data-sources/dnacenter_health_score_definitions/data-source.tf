
data "dnacenter_health_score_definitions" "example" {
  provider                   = dnacenter
  attribute                  = "string"
  device_type                = "string"
  id                         = "string"
  include_for_overall_health = "false"
  limit                      = 1
  offset                     = 1
  xca_lle_rid                = "string"
}

output "dnacenter_health_score_definitions_example" {
  value = data.dnacenter_health_score_definitions.example.items
}
