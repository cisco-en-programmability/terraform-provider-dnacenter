
resource "dnacenter_security_threats_rogue_allowed_list" "example" {
  provider = dnacenter

  parameters {

    category    = 1
    mac_address = "string"
  }
}

output "dnacenter_security_threats_rogue_allowed_list_example" {
  value = dnacenter_security_threats_rogue_allowed_list.example
}
