
resource "dnacenter_assurance_issues_resolve" "example" {
  provider    = dnacenter
  xca_lle_rid = "string"
  parameters {

    issue_ids = ["string"]
  }
}

output "dnacenter_assurance_issues_resolve_example" {
  value = dnacenter_assurance_issues_resolve.example
}