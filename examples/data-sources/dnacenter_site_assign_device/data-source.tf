
data "dnacenter_site_assign_device" "example" {
  provider = dnacenter
  site_id  = "string"
  device {

    ip = "string"
  }
}