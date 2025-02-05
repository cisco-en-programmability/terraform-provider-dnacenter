
data "dnacenter_security_threats_type" "example" {
  provider = dnacenter
}

output "dnacenter_security_threats_type_example" {
  value = data.dnacenter_security_threats_type.example.items
}
