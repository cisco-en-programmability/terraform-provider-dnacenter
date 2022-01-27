
data "dnacenter_nfv_provision" "example" {
  provider = dnacenter
  provisioning {

    device {

      custom_networks {

        ip_address_pool = "string"
        name            = "string"
        port            = "string"
      }
      device_serial_number = "string"
      ip                   = "string"
      service_providers {

        service_provider = "string"
        wan_interface {

          bandwidth      = "string"
          gateway        = "string"
          interface_name = "string"
          ip_address     = "string"
          subnetmask     = "string"
        }
      }
      services {

        admin_password_hash      = "string"
        central_manager_ip       = "string"
        central_registration_key = "string"
        common_key               = "string"
        disk                     = "string"
        mode                     = "string"
        system_ip                = "string"
        type                     = "string"
      }
      sub_pools {

        gateway          = "string"
        ip_subnet        = "string"
        name             = "string"
        parent_pool_name = "string"
        type             = "string"
      }
      tag_name = "string"
      template_param {

        asav {

          var1 = "string"
        }
        nfvis {

          var1 = "string"
        }
      }
      vlan {

        id         = "string"
        interfaces = "string"
        network    = "string"
        type       = "string"
      }
    }
    site {

      area {

        name        = "string"
        parent_name = "string"
      }
      building {

        address     = "string"
        latitude    = 1
        longitude   = 1
        name        = "string"
        parent_name = "string"
      }
      floor {

        height      = 1
        length      = 1
        name        = "string"
        parent_name = "string"
        rf_model    = "string"
        width       = 1
      }
      site_profile_name = "string"
    }
  }
  site_profile {

    device {

      custom_networks {

        connection_type = "string"
        name            = "string"
        network_mode    = "string"
        services_to_connect {

          service = "string"
        }
        vlan = "string"
      }
      custom_services {

        application_type = "string"
        image_name       = "string"
        name             = "string"
        profile          = "string"
        topology {

          assign_ip = "string"
          name      = "string"
          type      = "string"
        }
      }
      custom_template {

        device_type = "string"
        template    = "string"
      }
      device_type = "string"
      dia         = "false"
      service_providers {

        connect          = "false"
        default_gateway  = "false"
        link_type        = "string"
        service_provider = "string"
      }
      services {

        image_name = "string"
        mode       = "string"
        name       = "string"
        profile    = "string"
        topology {

          assign_ip = "string"
          name      = "string"
          type      = "string"
        }
        type = "string"
      }
      tag_name = "string"
      vlan {

        id   = "string"
        type = "string"
      }
    }
    site_profile_name = "string"
  }
}