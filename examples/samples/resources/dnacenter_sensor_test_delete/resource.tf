terraform {
  required_providers {
    dnacenter = {
      version = "1.1.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_sensor_test_delete" "example" {
  provider = dnacenter

  parameters {
    template_name = "string"
  }
}