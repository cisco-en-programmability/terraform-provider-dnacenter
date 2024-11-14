
resource "dnacenter_lan_automation_delete" "example" {
  provider = dnacenter
  id       = "string"
  parameters {

  }
}

output "dnacenter_lan_automation_delete_example" {
  value = dnacenter_lan_automation_delete.example
}