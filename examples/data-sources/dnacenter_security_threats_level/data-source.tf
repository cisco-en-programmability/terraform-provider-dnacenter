
data "dnacenter_security_threats_level" "example" {
  provider = dnacenter
}

output "dnacenter_security_threats_level_example" {
  value = data.dnacenter_security_threats_level.example.items
}
