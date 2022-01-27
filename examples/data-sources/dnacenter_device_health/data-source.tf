
data "dnacenter_device_health" "example" {
  provider    = dnacenter
  device_role = "string"
  end_time    = "hh:mm"
  health      = "string"
  limit       = "#"
  offset      = "#"
  site_id     = "string"
  start_time  = "hh:mm"
}

output "dnacenter_device_health_example" {
  value = data.dnacenter_device_health.example.items
}
