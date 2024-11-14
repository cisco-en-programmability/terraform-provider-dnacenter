
resource "dnacenter_assign_device_to_site" "example" {
  provider = dnacenter
  site_id  = "string"
  parameters {

    device {

      ip = "string"
    }
  }
}

output "dnacenter_assign_device_to_site_example" {
  value = dnacenter_assign_device_to_site.example
}