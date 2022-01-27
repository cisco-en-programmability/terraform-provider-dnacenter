terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_compliance_check_run" "create" {
  provider = dnacenter
  trigger_full= true
  categories=["PSIRT"]
  device_uuids=["3eb928b8-2414-4121-ac35-1247e5d666a4"]
}
output "dnacenter_compliance_check_run_create" {
  value = data.dnacenter_compliance_check_run.create
}


