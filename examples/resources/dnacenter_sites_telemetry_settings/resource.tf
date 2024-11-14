
resource "dnacenter_sites_telemetry_settings" "example" {
  provider = dnacenter

  parameters {

    application_visibility {

      collector {

        address        = "string"
        collector_type = "string"
        port           = 1
      }
      enable_on_wired_access_devices = "false"
    }
    id = "string"
    snmp_traps {

      external_trap_servers   = ["string"]
      use_builtin_trap_server = "false"
    }
    syslogs {

      external_syslog_servers   = ["string"]
      use_builtin_syslog_server = "false"
    }
    wired_data_collection {

      enable_wired_data_collectio = "false"
    }
    wireless_telemetry {

      enable_wireless_telemetry = "false"
    }
  }
}

output "dnacenter_sites_telemetry_settings_example" {
  value = dnacenter_sites_telemetry_settings.example
}