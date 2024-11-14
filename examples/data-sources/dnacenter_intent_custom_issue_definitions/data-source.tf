
data "dnacenter_intent_custom_issue_definitions" "example" {
  provider    = dnacenter
  id          = "string"
  xca_lle_rid = "string"
}

output "dnacenter_intent_custom_issue_definitions_example" {
  value = data.dnacenter_intent_custom_issue_definitions.example.item
}
