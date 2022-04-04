
resource "dnacenter_business_sda_wireless_controller_delete" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_ipaddress = "string"
  }
}