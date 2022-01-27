
data "dnacenter_command_runner_run_command" "example" {
  provider     = dnacenter
  commands     = ["string"]
  description  = "string"
  device_uuids = ["string"]
  name         = "string"
  timeout      = 1
}