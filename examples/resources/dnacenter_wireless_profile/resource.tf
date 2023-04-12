
resource "dnacenter_wireless_profile" "example" {
  provider = dnacenter

  parameters {

    profile_details {

      name  = "string"
      sites = ["string"]
      ssid_details {

        enable_fabric = "false"
        flex_connect {

          enable_flex_connect = "false"
          local_to_vlan       = 1
        }
        interface_name      = "string"
        name                = "string"
        policy_profile_name = "string"
        type                = "string"
        wlan_profile_name   = "string"
      }
    }
    wireless_profile_name = "string"
  }
}

output "dnacenter_wireless_profile_example" {
  value = dnacenter_wireless_profile.example
}