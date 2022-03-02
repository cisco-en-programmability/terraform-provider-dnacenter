
resource "dnacenter_service_provider" "example" {
  provider = dnacenter
  settings {
    qos {
      model        = "string"
      profile_name = "string"
      wan_provider = "string"
    }
  }
}

output "dnacenter_service_provider_example" {
  value = dnacenter_service_provider.example
}