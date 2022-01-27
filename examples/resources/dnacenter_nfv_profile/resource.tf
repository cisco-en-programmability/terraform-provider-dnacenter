
resource "dnacenter_nfv_profile" "example" {
  provider = dnacenter
  parameters {

    device {

      current_device_tag = "string"
      custom_networks {

        connection_type = "string"
        network_name    = "string"
        services_to_connect {

          service_name = "string"
        }
        vlan_id   = 1.0
        vlan_mode = "string"
      }
      custom_template {

        device_type   = "string"
        template      = "string"
        template_type = "string"
      }
      device_tag                          = "string"
      device_type                         = "string"
      direct_internet_access_for_firewall = "false"
      service_provider_profile {

        connect                        = "false"
        connect_default_gateway_on_wan = "false"
        link_type                      = "string"
        service_provider               = "string"
      }
      services {

        firewall_mode = "string"
        image_name    = "string"
        profile_type  = "string"
        service_name  = "string"
        service_type  = "string"
        v_nic_mapping {

          assign_ip_address_to_network = "string"
          network_type                 = "string"
        }
      }
      vlan_for_l2 {

        vlan_description = "string"
        vlan_id          = 1.0
        vlan_type        = "string"
      }
    }
    id           = "string"
    profile_name = "string"
  }
}

output "dnacenter_nfv_profile_example" {
  value = dnacenter_nfv_profile.example
}