
data "dnacenter_applications_health" "example" {
  provider           = dnacenter
  application_health = "string"
  application_name   = "string"
  device_id          = "string"
  end_time           = "hh:mm"
  limit              = "#"
  mac_address        = "string"
  offset             = "#"
  site_id            = "string"
  start_time         = "hh:mm"
}

output "dnacenter_applications_health_example" {
  value = data.dnacenter_applications_health.example.items
}
