
data "dnacenter_security_advisories" "example" {
  provider = dnacenter
}

output "dnacenter_security_advisories_example" {
  value = data.dnacenter_security_advisories.example.items
}
