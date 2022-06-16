
data "dnacenter_system_performance_historical" "example" {
  provider   = dnacenter
  end_time   = 1609459200
  kpi        = "string"
  start_time = 1609459200
}

output "dnacenter_system_performance_historical_example" {
  value = data.dnacenter_system_performance_historical.example.item
}
