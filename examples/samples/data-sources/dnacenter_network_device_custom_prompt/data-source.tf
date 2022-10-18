terraform {
  required_providers {
    dnacenter = {
      version = "1.0.9-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_network_device_custom_prompt" "example" {
  provider = dnacenter
}

output "dnacenter_network_device_custom_prompt_example" {
  value = data.dnacenter_network_device_custom_prompt.example.item
}
