
data "dnacenter_system_performance" "example" {
  provider   = dnacenter
  end_time   = "hh:mm"
  function   = "string"
  kpi        = "string"
  start_time = "hh:mm"
}

output "dnacenter_system_performance_example" {
  value = data.dnacenter_system_performance.example.item
}
