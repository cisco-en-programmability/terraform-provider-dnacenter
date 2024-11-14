
resource "dnacenter_service_provider_v2" "example" {
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
  }
}

output "dnacenter_service_provider_v2_example" {
  value = dnacenter_service_provider_v2.example
}