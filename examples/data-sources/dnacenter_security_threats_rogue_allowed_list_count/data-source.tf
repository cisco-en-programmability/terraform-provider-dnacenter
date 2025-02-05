
data "dnacenter_security_threats_rogue_allowed_list_count" "example" {
  provider = dnacenter
}

output "dnacenter_security_threats_rogue_allowed_list_count_example" {
  value = data.dnacenter_security_threats_rogue_allowed_list_count.example.item
}
