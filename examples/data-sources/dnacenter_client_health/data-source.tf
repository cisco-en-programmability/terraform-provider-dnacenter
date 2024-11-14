
data "dnacenter_client_health" "example" {
  provider  = dnacenter
  timestamp = 1.0
}

output "dnacenter_client_health_example" {
  value = data.dnacenter_client_health.example.items
}
