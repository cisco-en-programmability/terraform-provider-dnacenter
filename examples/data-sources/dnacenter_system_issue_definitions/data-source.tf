
data "dnacenter_system_issue_definitions" "example" {
  provider      = dnacenter
  attribute     = "string"
  device_type   = "string"
  id            = "string"
  issue_enabled = "false"
  limit         = 1
  name          = "string"
  offset        = 1
  order         = "string"
  priority      = "string"
  profile_id    = "string"
  sort_by       = "string"
  xca_lle_rid   = "string"
}

output "dnacenter_system_issue_definitions_example" {
  value = data.dnacenter_system_issue_definitions.example.items
}

data "dnacenter_system_issue_definitions" "example" {
  provider    = dnacenter
  id          = "string"
  xca_lle_rid = "string"
}

output "dnacenter_system_issue_definitions_example" {
  value = data.dnacenter_system_issue_definitions.example.item
}
