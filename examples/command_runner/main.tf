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

data "dna_command_runner_keywords" "response" {
  provider = dnacenter
}
output "dna_command_runner_keywords_response" {
  value = data.dna_command_runner_keywords.response
}

data "dna_command_runner_run_command" "response" {
  provider     = dnacenter
  commands     = ["pwd"]
  device_uuids = ["a9b86e42-6573-4f5d-a0bf-a743e290f46a"]
}
output "dna_command_runner_run_command_response" {
  value = data.dna_command_runner_run_command.response
}

data "dna_task" "response" {
  provider   = dnacenter
  depends_on = [data.dna_command_runner_run_command.response]
  task_id    = data.dna_command_runner_run_command.response.task_id
}

output "dna_task_response" {
  value = data.dna_task.response
}
