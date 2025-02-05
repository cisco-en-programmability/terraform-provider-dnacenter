
data "dnacenter_security_advisories_results_advisories_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_security_advisories_results_advisories_id_example" {
  value = data.dnacenter_security_advisories_results_advisories_id.example.item
}
