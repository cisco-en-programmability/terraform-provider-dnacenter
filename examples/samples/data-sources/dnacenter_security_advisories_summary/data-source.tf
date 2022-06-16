
data "dnacenter_security_advisories_summary" "example" {
  provider = dnacenter
}

output "dnacenter_security_advisories_summary_example" {
  value = data.dnacenter_security_advisories_summary.example.item
}
