terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}


resource "dnacenter_custom_issue_definitions" "example" {
  provider = dnacenter

  parameters {

    description             = "string"
    id                      = "string"
    is_enabled              = "false"
    is_notification_enabled = "false"
    name                    = "string"
    priority                = "string"
    rules {

      duration_in_minutes = 1
      facility            = "string"
      mnemonic            = "string"
      occurrences         = 1
      pattern             = "string"
      severity            = 1
    }
  }
}

output "dnacenter_custom_issue_definitions_example" {
  value = dnacenter_custom_issue_definitions.example
}
