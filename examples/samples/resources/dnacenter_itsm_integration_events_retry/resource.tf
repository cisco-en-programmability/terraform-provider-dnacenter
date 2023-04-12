
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_itsm_integration_events_retry" "example" {
  provider = dnacenter

  parameters {
    payload = ["string2"]
  }
}

output "dnacenter_itsm_integration_events_retry_example" {
  value = dnacenter_itsm_integration_events_retry.example
}


data "dnacenter_dnacaap_management_execution_status" "example" {
  depends_on   = [dnacenter_itsm_integration_events_retry.example]
  provider     = dnacenter
  execution_id = dnacenter_itsm_integration_events_retry.example.item.0.execution_id
}

output "dnacenter_dnacaap_management_execution_status_example" {
  value = data.dnacenter_dnacaap_management_execution_status.example.item
}
