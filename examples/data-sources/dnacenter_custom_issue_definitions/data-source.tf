
data "dnacenter_custom_issue_definitions" "example" {
  provider   = dnacenter
  facility   = "string"
  id         = "string"
  is_enabled = "false"
  limit      = 1
  mnemonic   = "string"
  name       = "string"
  offset     = 1
  order      = "string"
  priority   = "string"
  profile_id = "string"
  severity   = 1.0
  sort_by    = "string"
}

output "dnacenter_custom_issue_definitions_example" {
  value = data.dnacenter_custom_issue_definitions.example.items
}
