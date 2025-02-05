
resource "dnacenter_wireless_settings_anchor_groups_id" "example" {
  provider = dnacenter

  parameters {

    anchor_group_name = "string"
    id                = "string"
    mobility_anchors {

      anchor_priority     = "string"
      device_name         = "string"
      ip_address          = "string"
      mac_address         = "string"
      managed_anchor_wlc  = "false"
      mobility_group_name = "string"
      peer_device_type    = "string"
      private_ip          = "string"
    }
  }
}

output "dnacenter_wireless_settings_anchor_groups_id_example" {
  value = dnacenter_wireless_settings_anchor_groups_id.example
}
