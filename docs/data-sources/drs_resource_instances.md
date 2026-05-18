---
subcategory: "Data Replication Service (DRS)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_drs_resource_instances"
description: |-
  Use this data source to get the list of DRS resource instances.
---

# huaweicloud_drs_resource_instances

Use this data source to get the list of DRS resource instances, supporting filtering by resource type, tags and resource name.

## Example Usage

```hcl
data "huaweicloud_drs_resource_instances" "test" {
  resource_type = "migration"
}
```
