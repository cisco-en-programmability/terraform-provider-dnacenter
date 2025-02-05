terraform {
  required_providers {
    dnacenter = {
      version = "1.3.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}


data "dnacenter_authentication_policy_servers" "example" {
  provider = dnacenter
  # is_ise_enabled = "false"
  # role           = "string"
  # state          = "string"
}

output "dnacenter_authentication_policy_servers_example" {
  value = data.dnacenter_authentication_policy_servers.example.items
}
