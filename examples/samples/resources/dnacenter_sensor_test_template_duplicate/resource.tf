terraform {
  required_providers {
    dnacenter = {
      version = "0.3.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}


resource "dnacenter_sensor_test_template_duplicate" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    new_template_name = "string"
    template_name     = "string"
  }
}