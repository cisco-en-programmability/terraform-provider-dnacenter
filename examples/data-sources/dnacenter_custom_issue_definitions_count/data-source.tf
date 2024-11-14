
data "dnacenter_custom_issue_definitions_count" "example" {
  provider    = dnacenter
  facility    = "string"
  id          = "string"
  is_enabled  = "false"
  mnemonic    = "string"
  name        = "string"
  priority    = "string"
  profile_id  = "string"
  severity    = 1.0
  xca_lle_rid = "string"
}

output "dnacenter_custom_issue_definitions_count_example" {
  value = data.dnacenter_custom_issue_definitions_count.example.item
}
