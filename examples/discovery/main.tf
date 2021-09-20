
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

resource "dna_discovery" "response" {
  provider = dnacenter
  item {
    cdp_level                 = 16
    discovery_type            = "CDP"
    global_credential_id_list = ["90acbab8-03d5-4726-9c19-e1e51a40b3cd", "f979d842-f6fd-456a-8137-2cb5113cd2e8"]
    ip_address_list           = "10.10.22.22"
    name                      = "start_discovery_test2"
    protocol_order            = "ssh"
    # id                        = "67"
    # enable_password_list      = [""]
    # user_name_list            = [""]
    # password_list             = [""]
    # ip_filter_list            = [""]
    # http_read_credential {
    #   port   = 0
    #   secure = false
    # }
    # device_ids               = " "
    # discovery_condition      = "In Progress"
    # discovery_status         = "Active"
    # is_auto_cdp              = true
    # preferred_mgmt_ip_method = "None"
    # retry                    = 3
    # timeout                  = 5
    # http_write_credential {
    #   port   = 0
    #   secure = false
    # }
  }
}
output "dna_discovery_response" {
  value = dna_discovery.response
}

data "dna_discovery_device_count" "amount" {
  provider = dnacenter
}

output "dna_discovery_device_count" {
  depends_on = [dna_discovery.response]
  value      = data.dna_discovery_device_count.amount.response
}
