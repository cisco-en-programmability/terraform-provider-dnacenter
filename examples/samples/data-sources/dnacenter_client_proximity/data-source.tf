
data "dnacenter_client_proximity" "example" {
  provider        = dnacenter
  number_days     = 1
  time_resolution = 1.0
  username        = "string"
}

output "dnacenter_client_proximity_example" {
  value = data.dnacenter_client_proximity.example.item
}
