
resource "dnacenter_execute_suggested_actions_commands" "example" {
  provider = dnacenter
  parameters {

    entity_type  = "string"
    entity_value = "string"
  }
}

output "dnacenter_execute_suggested_actions_commands_example" {
  value = dnacenter_execute_suggested_actions_commands.example
}