
data "dnacenter_system_performance" "example" {
  provider   = dnacenter
  end_time   = 1609459200
  function   = "string"
  kpi        = "string"
  start_time = 1609459200
}

output "dnacenter_system_performance_example" {
  value = data.dnacenter_system_performance.example.item
}
