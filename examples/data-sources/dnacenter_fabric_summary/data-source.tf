
data "dnacenter_fabric_summary" "example" {
  provider    = dnacenter
  end_time    = 1609459200
  start_time  = 1609459200
  xca_lle_rid = "string"
}

output "dnacenter_fabric_summary_example" {
  value = data.dnacenter_fabric_summary.example.item
}
