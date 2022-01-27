
data "dnacenter_service_provider_create" "example" {
  provider = dnacenter
  settings {

    qos {

      model        = "string"
      profile_name = "string"
      wan_provider = "string"
    }
  }
}