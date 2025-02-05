
data "dnacenter_security_advisories_results_advisories_count" "example" {
  provider               = dnacenter
  cvss_base_score        = "string"
  device_count           = 1.0
  id                     = "string"
  security_impact_rating = "string"
}

output "dnacenter_security_advisories_results_advisories_count_example" {
  value = data.dnacenter_security_advisories_results_advisories_count.example.item
}
