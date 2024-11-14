
resource "dnacenter_service_provider" "example" {
  provider = dnacenter

  parameters {

    settings {

      qos {

        model            = "string"
        old_profile_name = "string"
        profile_name     = "string"
        wan_provider     = "string"
      }
    }
    sp_profile_name = "string"
  }
}

output "dnacenter_service_provider_example" {
  value = dnacenter_service_provider.example
}