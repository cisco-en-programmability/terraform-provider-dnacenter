---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_disasterrecovery_system_status Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Disaster Recovery.
  Detailed and Summarized status of DR components (Active, Standby and Witness system's health).
---

# dnacenter_disasterrecovery_system_status (Data Source)

It performs read operation on Disaster Recovery.

- Detailed and Summarized status of DR components (Active, Standby and Witness system's health).

## Example Usage

```terraform
data "dnacenter_disasterrecovery_system_status" "example" {
  provider = dnacenter
}

output "dnacenter_disasterrecovery_system_status_example" {
  value = data.dnacenter_disasterrecovery_system_status.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **ipconfig** (List of Object) (see [below for nested schema](#nestedobjatt--item--ipconfig))
- **ipsec_tunnel** (List of Object) (see [below for nested schema](#nestedobjatt--item--ipsec_tunnel))
- **main** (List of Object) (see [below for nested schema](#nestedobjatt--item--main))
- **recovery** (List of Object) (see [below for nested schema](#nestedobjatt--item--recovery))
- **site** (String)
- **state** (String)
- **witness** (List of Object) (see [below for nested schema](#nestedobjatt--item--witness))

<a id="nestedobjatt--item--ipconfig"></a>
### Nested Schema for `item.ipconfig`

Read-Only:

- **interface** (String)
- **ip** (String)
- **vip** (String)


<a id="nestedobjatt--item--ipsec_tunnel"></a>
### Nested Schema for `item.ipsec_tunnel`

Read-Only:

- **side_a** (String)
- **side_b** (String)
- **status** (String)


<a id="nestedobjatt--item--main"></a>
### Nested Schema for `item.main`

Read-Only:

- **ipconfig** (List of Object) (see [below for nested schema](#nestedobjatt--item--main--ipconfig))
- **nodes** (List of Object) (see [below for nested schema](#nestedobjatt--item--main--nodes))
- **state** (String)

<a id="nestedobjatt--item--main--ipconfig"></a>
### Nested Schema for `item.main.ipconfig`

Read-Only:

- **interface** (String)
- **ip** (String)
- **vip** (String)


<a id="nestedobjatt--item--main--nodes"></a>
### Nested Schema for `item.main.nodes`

Read-Only:

- **hostname** (String)
- **ipaddresses** (List of Object) (see [below for nested schema](#nestedobjatt--item--main--nodes--ipaddresses))
- **state** (String)

<a id="nestedobjatt--item--main--nodes--ipaddresses"></a>
### Nested Schema for `item.main.nodes.state`

Read-Only:

- **interface** (String)
- **ip** (String)
- **vip** (String)




<a id="nestedobjatt--item--recovery"></a>
### Nested Schema for `item.recovery`

Read-Only:

- **ipconfig** (List of Object) (see [below for nested schema](#nestedobjatt--item--recovery--ipconfig))
- **nodes** (List of Object) (see [below for nested schema](#nestedobjatt--item--recovery--nodes))
- **state** (String)

<a id="nestedobjatt--item--recovery--ipconfig"></a>
### Nested Schema for `item.recovery.ipconfig`

Read-Only:

- **interface** (String)
- **ip** (String)
- **vip** (String)


<a id="nestedobjatt--item--recovery--nodes"></a>
### Nested Schema for `item.recovery.nodes`

Read-Only:

- **hostname** (String)
- **ipconfig** (List of Object) (see [below for nested schema](#nestedobjatt--item--recovery--nodes--ipconfig))
- **state** (String)

<a id="nestedobjatt--item--recovery--nodes--ipconfig"></a>
### Nested Schema for `item.recovery.nodes.state`

Read-Only:

- **interface** (String)
- **ip** (String)
- **vip** (String)




<a id="nestedobjatt--item--witness"></a>
### Nested Schema for `item.witness`

Read-Only:

- **ipconfig** (List of Object) (see [below for nested schema](#nestedobjatt--item--witness--ipconfig))
- **nodes** (List of Object) (see [below for nested schema](#nestedobjatt--item--witness--nodes))
- **state** (String)

<a id="nestedobjatt--item--witness--ipconfig"></a>
### Nested Schema for `item.witness.ipconfig`

Read-Only:

- **interface** (String)
- **ip** (String)
- **vip** (String)


<a id="nestedobjatt--item--witness--nodes"></a>
### Nested Schema for `item.witness.nodes`

Read-Only:

- **hostname** (String)
- **ipconfig** (List of Object) (see [below for nested schema](#nestedobjatt--item--witness--nodes--ipconfig))
- **state** (String)

<a id="nestedobjatt--item--witness--nodes--ipconfig"></a>
### Nested Schema for `item.witness.nodes.state`

Read-Only:

- **interface** (String)
- **ip** (String)
- **vip** (String)

