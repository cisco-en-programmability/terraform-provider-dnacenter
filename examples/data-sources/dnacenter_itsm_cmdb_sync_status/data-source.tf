
data "dnacenter_itsm_cmdb_sync_status" "example" {
  provider = dnacenter
  date     = "string"
  status   = "string"
}

output "dnacenter_itsm_cmdb_sync_status_example" {
  value = data.dnacenter_itsm_cmdb_sync_status.example.items
}
