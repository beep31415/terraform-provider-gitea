---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitea_org Data Source - gitea"
subcategory: ""
description: |-
  Fetches a Gitea organization.
---

# gitea_org (Data Source)

Fetches a Gitea organization.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the organisation without spaces.

### Read-Only

- `avatar_url` (String) The organization avatar url.
- `description` (String) The organization description.
- `full_name` (String) The full name of the organisation.
- `id` (Number) Identifier attribute.
- `location` (String) The organization location.
- `repo_admin_change_team_access` (Boolean) Flag that indicates whether admin can change organization team access.
- `visibility` (String) The organization visibility.
- `website` (String) The organization website.


