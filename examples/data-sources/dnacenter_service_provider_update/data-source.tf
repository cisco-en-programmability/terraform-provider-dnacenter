
data "dnacenter_service_provider_update" "example" {
  provider = dnacenter
  settings {

    qos {

      model            = "string"
      old_profile_name = "string"
      profile_name     = "string"
      wan_provider     = "string"
    }
  }
}