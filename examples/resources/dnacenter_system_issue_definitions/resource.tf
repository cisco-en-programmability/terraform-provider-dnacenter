
resource "dnacenter_system_issue_definitions" "example" {
  provider = dnacenter

  parameters {

    id                              = "string"
    issue_enabled                   = "false"
    priority                        = "string"
    synchronize_to_health_threshold = "false"
    threshold_value                 = 1.0
  }
}

output "dnacenter_system_issue_definitions_example" {
  value = dnacenter_system_issue_definitions.example
}