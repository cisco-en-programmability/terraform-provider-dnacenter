
resource "dnacenter_wireless_settings_ap_authorization_lists" "example" {
  provider = dnacenter

  parameters {

    ap_authorization_list_name = "string"
    local_authorization {

      ap_mac_entries           = ["string"]
      ap_serial_number_entries = ["string"]
    }
    remote_authorization {

      aaa_servers                     = ["string"]
      authorize_ap_with_mac           = "false"
      authorize_ap_with_serial_number = "false"
    }
  }
}

output "dnacenter_wireless_settings_ap_authorization_lists_example" {
  value = dnacenter_wireless_settings_ap_authorization_lists.example
}
