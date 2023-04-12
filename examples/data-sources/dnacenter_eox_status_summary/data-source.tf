
data "dnacenter_eox_status_summary" "example" {
  provider = dnacenter
}

output "dnacenter_eox_status_summary_example" {
  value = data.dnacenter_eox_status_summary.example.item
}
