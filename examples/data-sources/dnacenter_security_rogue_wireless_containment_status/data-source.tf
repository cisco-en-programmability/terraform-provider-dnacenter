
data "dnacenter_security_rogue_wireless_containment_status" "example" {
  provider    = dnacenter
  mac_address = "string"
}

output "dnacenter_security_rogue_wireless_containment_status_example" {
  value = data.dnacenter_security_rogue_wireless_containment_status.example.items
}
