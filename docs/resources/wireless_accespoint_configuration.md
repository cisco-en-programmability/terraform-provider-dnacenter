---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_wireless_accespoint_configuration Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
  User can configure multiple access points with required options using this intent API
---

# dnacenter_wireless_accespoint_configuration (Resource)

It performs create operation on Wireless.

- User can configure multiple access points with required options using this intent API

~>**Warning:**
This resource does not represent a real-world entity in Cisco DNA Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco DNA Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "dnacenter_wireless_accespoint_configuration" "example" {
  provider = dnacenter
  parameters {

    admin_status = "false"
    ap_height    = 1.0
    ap_list {

      ap_name     = "string"
      ap_name_new = "string"
      mac_address = "string"
    }
    ap_mode                        = 1
    configure_admin_status         = "false"
    configure_ap_height            = "false"
    configure_ap_mode              = "false"
    configure_failover_priority    = "false"
    configure_hacontroller         = "false"
    configure_led_brightness_level = "false"
    configure_led_status           = "false"
    configure_location             = "false"
    failover_priority              = 1
    led_brightness_level           = 1
    led_status                     = "false"
    location                       = "string"
    primary_controller_name        = "string"
    primary_ip_address {

      address = "string"
    }
    radio_configurations {

      admin_status                    = "false"
      antenna_cable_name              = "string"
      antenna_degree                  = 1
      antenna_elev_angle_degree       = 1
      antenna_elev_angle_sign         = 1
      antenna_gain                    = 1
      antenna_pattern_name            = "string"
      cable_loss                      = 1.0
      channel_assignment_mode         = 1
      channel_number                  = 1
      channel_width                   = 1
      clean_air_si                    = 1
      configure_admin_status          = "false"
      configure_antenna_cable         = "false"
      configure_antenna_degree        = "false"
      configure_antenna_pattern_name  = "false"
      configure_channel               = "false"
      configure_channel_width         = "false"
      configure_clean_air_si          = "false"
      configure_elev_angle_degree     = "false"
      configure_power                 = "false"
      configure_radio_role_assignment = "false"
      power_assignment_mode           = 1
      powerlevel                      = 1
      radio_band                      = "string"
      radio_role_assignment           = "string"
      radio_type                      = 1
    }
    secondary_controller_name = "string"
    secondary_ip_address {

      address = "string"
    }
    tertiary_controller_name = "string"
    tertiary_ip_address {

      address = "string"
    }
  }
}

output "dnacenter_wireless_accespoint_configuration_example" {
  value = dnacenter_wireless_accespoint_configuration.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `admin_status` (String) Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.
- `ap_height` (Number) Configure the height of the access point by setting a value between 3 and height of the floor.
- `ap_list` (Block List) (see [below for nested schema](#nestedblock--parameters--ap_list))
- `ap_mode` (Number) Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".
- `configure_admin_status` (String) To change the access point's admin status, set this parameter's value to "true".
- `configure_ap_height` (String) To change the access point's height, set this parameter's value to "true".
- `configure_ap_mode` (String) To change the access point's mode, set this parameter's value to "true".
- `configure_failover_priority` (String) To change the access point's failover priority, set this parameter's value to "true".
- `configure_hacontroller` (String) To change the access point's HA controller, set this parameter's value to "true".
- `configure_led_brightness_level` (String) To change the access point's LED brightness level, set this parameter's value to "true".
- `configure_led_status` (String) To change the access point's LED status, set this parameter's value to "true".
- `configure_location` (String) To change the access point's location, set this parameter's value to "true".
- `failover_priority` (Number) Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".
- `led_brightness_level` (Number) Configure the access point's LED brightness level by setting a value between 1 and 8.
- `led_status` (String) Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.
- `location` (String) Configure the access point's location.
- `primary_controller_name` (String) Configure the hostname for an access point's primary controller.
- `primary_ip_address` (Block List) (see [below for nested schema](#nestedblock--parameters--primary_ip_address))
- `radio_configurations` (Block List) (see [below for nested schema](#nestedblock--parameters--radio_configurations))
- `secondary_controller_name` (String) Configure the hostname for an access point's secondary controller.
- `secondary_ip_address` (Block List) (see [below for nested schema](#nestedblock--parameters--secondary_ip_address))
- `tertiary_controller_name` (String) Configure the hostname for an access point's tertiary controller.
- `tertiary_ip_address` (Block List) (see [below for nested schema](#nestedblock--parameters--tertiary_ip_address))

<a id="nestedblock--parameters--ap_list"></a>
### Nested Schema for `parameters.ap_list`

Optional:

- `ap_name` (String) The current host name of the access point.
- `ap_name_new` (String) The modified hostname of the access point.
- `mac_address` (String) The ethernet MAC address of the access point.


<a id="nestedblock--parameters--primary_ip_address"></a>
### Nested Schema for `parameters.primary_ip_address`

Optional:

- `address` (String) Configure the IP address for an access point's primary controller.


<a id="nestedblock--parameters--radio_configurations"></a>
### Nested Schema for `parameters.radio_configurations`

Optional:

- `admin_status` (String) Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.
- `antenna_cable_name` (String) Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".
- `antenna_degree` (Number) Configure the antenna degree on the specified radio for an access point.
- `antenna_elev_angle_degree` (Number) Configure the antenna elevation angle on the specified radio for an access point.
- `antenna_elev_angle_sign` (Number) Configure the antenna elevation angle direction on the specified radio for an access point: for up, set "1"; for down, set "-1".
- `antenna_gain` (Number) Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi).
- `antenna_pattern_name` (String) Configure the antenna pattern name on the specified radio for an access point. If antenna gain needs to be configured, set this parameter's value to "other".
- `cable_loss` (Number) Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).
- `channel_assignment_mode` (Number) Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
- `channel_number` (Number) Configure the channel number on the specified radio for an access point.
- `channel_width` (Number) Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; and for 160 MHz, set "6".
- `clean_air_si` (Number) Configure CleanAir or Spectrum Intelligence on the specified radio for an access point. Set this parameter's value to "0" to disable the feature or "1" to enable it.
- `configure_admin_status` (String) To change the admin status on the specified radio for an access point, set this parameter's value to "true".
- `configure_antenna_cable` (String) To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".
- `configure_antenna_degree` (String) To change the antenna degree on the specified radio for an access point, set this parameter's value to "true".
- `configure_antenna_pattern_name` (String) To change the antenna pattern name on the specified radio for an access point, set the value for this parameter to "true".
- `configure_channel` (String) To change the channel on the specified radio for an access point, set this parameter's value to "true".
- `configure_channel_width` (String) To change the channel width on the specified radio for an access point, set this parameter's value to "true".
- `configure_clean_air_si` (String) To enable or disable either CleanAir or Spectrum Intelligence on the specified radio for an access point, set this parameter's value to "true".
- `configure_elev_angle_degree` (String) To change the elevation angle degree on the specified radio for an access point, set this parameter's value to "true".
- `configure_power` (String) To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".
- `configure_radio_role_assignment` (String) To change the radio role on the specified radio for an access point, set this parameter's value to "true".
- `power_assignment_mode` (Number) Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
- `powerlevel` (Number) Configure the power level on the specified radio for an access point by setting a value between 1 and 8.
- `radio_band` (String) Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5".
- `radio_role_assignment` (String) Configure one of the following roles on the specified radio for an access point: "auto", "serving", or "monitor".
- `radio_type` (Number) Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".


<a id="nestedblock--parameters--secondary_ip_address"></a>
### Nested Schema for `parameters.secondary_ip_address`

Optional:

- `address` (String) Configure the IP address for an access point's secondary controller.


<a id="nestedblock--parameters--tertiary_ip_address"></a>
### Nested Schema for `parameters.tertiary_ip_address`

Optional:

- `address` (String) Configure the IP address for an access point's tertiary controller.



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)

