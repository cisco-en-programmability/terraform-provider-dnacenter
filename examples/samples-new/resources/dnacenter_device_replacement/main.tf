terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_device_replacement" "example" {
    provider = dnacenter
    parameters {
      
      creation_time = 1
      family = "string"
      faulty_device_id = "string"
      faulty_device_name = "string"
      faulty_device_platform = "string"
      faulty_device_serial_number = "string"
      id = "string"
      neighbour_device_id = "string"
      network_readiness_task_id = "string"
      replacement_device_platform = "string"
      replacement_device_serial_number = "string"
      replacement_status = "string"
      replacement_time = 1
      workflow_id = "string"
    }
}

output "dnacenter_device_replacement_example" {
    value = dnacenter_device_replacement.example
}