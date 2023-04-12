
data "dnacenter_service_provider_v2" "example" {
  provider = dnacenter
}

output "dnacenter_service_provider_v2_example" {
  value = data.dnacenter_service_provider_v2.example.items
}
