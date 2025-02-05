
data "dnacenter_security_threats_rogue_allowed_list" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_security_threats_rogue_allowed_list_example" {
  value = data.dnacenter_security_threats_rogue_allowed_list.example.items
}
