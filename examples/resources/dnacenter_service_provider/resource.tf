resource "dnacenter_service_provider" "example" {
  provider = dnacenter
  parameters {

    settings {
      qos {
        profile_name = "Test_tf_new"
        model        = "8-class-model"
        wan_provider = "test1-provider"
      }
    }
  }
}

output "dnacenter_service_provider_example" {
  value = dnacenter_service_provider.example
}
