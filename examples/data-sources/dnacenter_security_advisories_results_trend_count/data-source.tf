
data "dnacenter_security_advisories_results_trend_count" "example" {
  provider  = dnacenter
  scan_time = 1.0
}

output "dnacenter_security_advisories_results_trend_count_example" {
  value = data.dnacenter_security_advisories_results_trend_count.example.item
}
