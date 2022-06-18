
data "dnacenter_client_health" "example" {
  provider  = dnacenter
  timestamp = "string"
}

output "dnacenter_client_health_example" {
  value = data.dnacenter_client_health.example.items
}
