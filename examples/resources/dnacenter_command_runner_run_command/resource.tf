
resource "dnacenter_command_runner_run_command" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    commands     = ["string"]
    description  = "string"
    device_uuids = ["string"]
    name         = "string"
    timeout      = 1
  }
}