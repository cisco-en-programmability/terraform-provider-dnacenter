
data "dnacenter_sensor" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_sensor_example" {
  value = data.dnacenter_sensor.example.items
}
