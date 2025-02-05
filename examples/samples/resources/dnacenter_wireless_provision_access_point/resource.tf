terraform {
  required_providers {
    dnacenter = {
      version = "1.3.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_wireless_provision_access_point" "example" {
  provider = dnacenter

  parameters {
    payload {
      custom_ap_group_name   = "string"
      custom_flex_group_name = ["string"]
      device_name            = "string"
      rf_profile             = "string"
      site_id                = "string"
      site_name_hierarchy    = "string"
      type                   = "string"
    }
  }
}
