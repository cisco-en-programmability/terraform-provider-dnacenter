
data "dnacenter_service_provider" "example" {
  provider = dnacenter
}

output "dnacenter_service_provider_example" {
  value = data.dnacenter_service_provider.example.items
}
