
resource "dnacenter_system_issue_definitions_count" "example" {
  provider = dnacenter

}

output "dnacenter_system_issue_definitions_count_example" {
  value = dnacenter_system_issue_definitions_count.example
}