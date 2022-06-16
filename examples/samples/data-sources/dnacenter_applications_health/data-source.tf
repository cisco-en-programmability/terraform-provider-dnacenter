
data "dnacenter_applications_health" "example" {
  provider           = dnacenter
  application_health = "string"
  application_name   = "string"
  device_id          = "string"
  end_time           = 1609459200
  limit              = 1
  mac_address        = "string"
  offset             = 1
  site_id            = "string"
  start_time         = 1609459200
}

output "dnacenter_applications_health_example" {
  value = data.dnacenter_applications_health.example.items
}
