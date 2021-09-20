terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dna_pnp_workflow_count" "response" {
  provider = dnacenter
  name     = ["Workflow 2", "Workflow 1"]
}
output "dna_pnp_workflow_count_response" {
  value = data.dna_pnp_workflow_count.response
}


data "dna_pnp_device_sync_result_vacct" "response" {
  provider = dnacenter
  name     = ""
  domain   = ""
}
output "dna_pnp_device_sync_result_vacct_response" {
  value = data.dna_pnp_device_sync_result_vacct.response
}


data "dna_pnp_workflow" "response" {
  provider = dnacenter
  name     = ["Workflow 1"]
}
output "dna_pnp_workflow_response" {
  value = data.dna_pnp_workflow.response
}


data "dna_pnp_global_settings" "response" {
  provider = dnacenter
}
output "dna_pnp_global_settings_response" {
  value = data.dna_pnp_global_settings.response
}


data "dna_pnp_virtual_account" "response" {
  provider = dnacenter
  domain   = ""
}
output "dna_pnp_virtual_account_response" {
  value = data.dna_pnp_virtual_account.response
}


data "dna_pnp_smart_account" "response" {
  provider = dnacenter
}
output "dna_pnp_smart_account_response" {
  value = data.dna_pnp_smart_account.response
}


data "dna_pnp_device" "response" {
  provider = dnacenter
}
output "dna_pnp_device_response" {
  value = data.dna_pnp_device.response
}

data "dna_pnp_device_count" "response" {
  provider = dnacenter
  # name     = ["FOCTEST1"]
}
output "dna_pnp_device_count_response" {
  value = data.dna_pnp_device_count.response
}


data "dna_pnp_device_history" "response" {
  provider      = dnacenter
  serial_number = "FOCTEST1"
}
output "dna_pnp_device_history_response" {
  value = data.dna_pnp_device_history.response
}

data "dna_pnp_device_unclaim" "response" {
  provider       = dnacenter
  device_id_list = []
}
output "dna_pnp_device_unclaim_response" {
  value = data.dna_pnp_device_unclaim.response
}
