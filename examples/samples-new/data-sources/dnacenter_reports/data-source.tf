terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}
/*
data "dnacenter_reports" "example" {
  provider      = dnacenter
  view_group_id = "string"
  view_id       = "string"
}

output "dnacenter_reports_example" {
  value = data.dnacenter_reports.example.items
}
*/
data "dnacenter_reports" "example" {
  provider  = dnacenter
  report_id = "string"
}

output "dnacenter_reports_example" {
  value = data.dnacenter_reports.example.item
}
