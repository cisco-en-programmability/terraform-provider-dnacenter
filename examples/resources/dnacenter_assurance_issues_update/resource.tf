
resource "dnacenter_assurance_issues_update" "example" {
  provider        = dnacenter
  accept_language = "string"
  id              = "string"
  xca_lle_rid     = "string"
  parameters {

    notes = "string"
  }
}

output "dnacenter_assurance_issues_update_example" {
  value = dnacenter_assurance_issues_update.example
}