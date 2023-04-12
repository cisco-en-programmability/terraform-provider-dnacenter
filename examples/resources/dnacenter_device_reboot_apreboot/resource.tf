
resource "dnacenter_device_reboot_apreboot" "example" {
  provider = dnacenter

  parameters {

    ap_mac_addresses = ["string"]
  }
}

output "dnacenter_device_reboot_apreboot_example" {
  value = dnacenter_device_reboot_apreboot.example
}