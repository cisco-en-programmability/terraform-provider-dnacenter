
resource "dnacenter_command_runner_run_command" "example" {
  provider = dnacenter
  parameters {

    commands     = ["string"]
    description  = "string"
    device_uuids = ["string"]
    name         = "string"
    timeout      = 1
  }
}

output "dnacenter_command_runner_run_command_example" {
  value = dnacenter_command_runner_run_command.example
}