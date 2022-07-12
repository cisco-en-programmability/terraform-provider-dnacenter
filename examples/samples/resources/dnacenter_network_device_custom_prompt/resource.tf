terraform {
  required_providers {
    dnacenter = {
      version = "1.0.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_network_device_custom_prompt" "example" {
  provider = dnacenter
  parameters {

    password_prompt = "******"
    username_prompt = "string2"
  }
}

output "dnacenter_network_device_custom_prompt_example" {
  value     = dnacenter_network_device_custom_prompt.example
  sensitive = true
}