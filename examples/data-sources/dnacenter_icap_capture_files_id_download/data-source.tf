
data "dnacenter_icap_capture_files_id_download" "example" {
  provider    = dnacenter
  id          = "string"
  dirpath     = "string"
  xca_lle_rid = "string"
}

output "dnacenter_icap_capture_files_id_download_example" {
  value = data.dnacenter_icap_capture_files_id_download.example.item
}
