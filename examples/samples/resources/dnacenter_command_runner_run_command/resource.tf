terraform {
  required_providers {
    dnacenter = {
      version = "1.0.8-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

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