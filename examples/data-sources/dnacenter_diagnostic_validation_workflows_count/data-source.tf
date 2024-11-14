
data "dnacenter_diagnostic_validation_workflows_count" "example" {
  provider   = dnacenter
  end_time   = 1609459200
  run_status = "string"
  start_time = 1609459200
}

output "dnacenter_diagnostic_validation_workflows_count_example" {
  value = data.dnacenter_diagnostic_validation_workflows_count.example.item
}
