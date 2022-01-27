
data "dnacenter_system_performance_historical" "example" {
  provider   = dnacenter
  end_time   = "hh:mm"
  kpi        = "string"
  start_time = "hh:mm"
}

output "dnacenter_system_performance_historical_example" {
  value = data.dnacenter_system_performance_historical.example.item
}
