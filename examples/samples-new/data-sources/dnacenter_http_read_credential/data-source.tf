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

data "dnacdnacenter_http_read_credential" "example" {
  provider = dnacenter
  payload {

    comments           = "string"
    credential_type    = "string"
    description        = "string"
    id                 = "string"
    instance_tenant_id = "string"
    instance_uuid      = "string"
    password           = "******"
    port               = 1
    secure             = "false"
    username           = "string"
  }
}