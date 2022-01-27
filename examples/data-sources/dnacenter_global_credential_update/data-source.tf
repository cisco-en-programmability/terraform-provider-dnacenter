
data "dnacenter_global_credential_update" "example" {
  provider             = dnacenter
  global_credential_id = "string"
  site_uuids           = ["string"]
}