
data "dnacenter_threat_detail" "example" {
  provider      = dnacenter
  end_time      = 1
  is_new_threat = "false"
  limit         = 1
  offset        = 1
  site_id       = ["string"]
  start_time    = 1
  threat_level  = ["string"]
  threat_type   = ["string"]
}