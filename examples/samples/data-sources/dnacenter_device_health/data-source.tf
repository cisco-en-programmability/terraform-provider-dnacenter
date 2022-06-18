
data "dnacenter_device_health" "example" {
  provider    = dnacenter
  device_role = "string"
  end_time    = 1609459200
  health      = "string"
  limit       = 1
  offset      = 1
  site_id     = "string"
  start_time  = 1609459200
}

output "dnacenter_device_health_example" {
  value = data.dnacenter_device_health.example.items
}
