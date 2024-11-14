
data "dnacenter_diagnostic_validation_workflows" "example" {
  provider   = dnacenter
  end_time   = 1609459200
  limit      = 1
  offset     = 1
  run_status = "string"
  start_time = 1609459200
}

output "dnacenter_diagnostic_validation_workflows_example" {
  value = data.dnacenter_diagnostic_validation_workflows.example.items
}

data "dnacenter_diagnostic_validation_workflows" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_diagnostic_validation_workflows_example" {
  value = data.dnacenter_diagnostic_validation_workflows.example.item
}
