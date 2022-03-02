---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_compliance Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  
---

# dnacenter_compliance (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **categories** (List of String)
- **device_uuids** (List of String)
- **trigger_full** (String)


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **ap_manager_interface_ip** (String)
- **associated_wlc_ip** (String)
- **boot_date_time** (String)
- **collection_interval** (String)
- **collection_status** (String)
- **error_code** (String)
- **error_description** (String)
- **family** (String)
- **hostname** (String)
- **id** (String)
- **instance_tenant_id** (String)
- **instance_uuid** (String)
- **interface_count** (String)
- **inventory_status_detail** (String)
- **last_update_time** (Number)
- **last_updated** (String)
- **line_card_count** (String)
- **line_card_id** (String)
- **location** (String)
- **location_name** (String)
- **mac_address** (String)
- **management_ip_address** (String)
- **memory_size** (String)
- **platform_id** (String)
- **reachability_failure_reason** (String)
- **reachability_status** (String)
- **role** (String)
- **role_source** (String)
- **serial_number** (String)
- **series** (String)
- **snmp_contact** (String)
- **snmp_location** (String)
- **software_type** (String)
- **software_version** (String)
- **tag_count** (String)
- **tunnel_udp_port** (String)
- **type** (String)
- **up_time** (String)
- **waas_device_mode** (String)

