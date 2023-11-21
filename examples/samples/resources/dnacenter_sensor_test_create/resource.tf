
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.28-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}


resource "dnacenter_sensor_test_create" "example" {
  provider = dnacenter

  parameters {
    ap_coverage {

      bands                 = "string"
      number_of_aps_to_test = "string"
      rssi_threshold        = "string"
    }
    r_connection  = "string"
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
