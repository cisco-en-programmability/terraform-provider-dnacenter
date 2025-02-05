
data "dnacenter_icap_capture_files_id" "example" {
  provider    = dnacenter
  id          = "string"
  xca_lle_rid = "string"
}

output "dnacenter_icap_capture_files_id_example" {
  value = data.dnacenter_icap_capture_files_id.example.item
}
