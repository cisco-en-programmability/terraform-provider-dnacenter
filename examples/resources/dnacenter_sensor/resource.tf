
resource "dnacenter_sensor" "example" {
  provider = dnacenter
  parameters {

    ap_coverage {

      bands                 = "string"
      number_of_aps_to_test = "string"
      rssi_threshold        = "string"
    }
    connection    = "string"
    model_version = 1
    name          = "string"
    ssids {

      auth_type    = "string"
      categories   = ["string"]
      profile_name = "string"
      psk          = "string"
      qos_policy   = "string"
      ssid         = "string"
      tests {

        config = ["string"]
        name   = "string"
      }
      third_party {

        selected = "false"
      }
    }
  }
}

output "dnacenter_sensor_example" {
  value = dnacenter_sensor.example
}