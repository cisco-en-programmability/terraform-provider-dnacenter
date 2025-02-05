
data "dnacenter_icap_capture_files" "example" {
  provider    = dnacenter
  ap_mac      = "string"
  client_mac  = "string"
  end_time    = 1609459200
  limit       = 1
  offset      = 1
  order       = "string"
  sort_by     = "string"
  start_time  = 1609459200
  type        = "string"
  xca_lle_rid = "string"
}

output "dnacenter_icap_capture_files_example" {
  value = data.dnacenter_icap_capture_files.example.items
}
