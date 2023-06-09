---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitea Provider"
subcategory: ""
description: |-
  Interact with Gitea hosting server.
---

# gitea Provider

Interact with Gitea hosting server.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `base_url` (String) Gitea server base URL. Can also be set from GITEA_BASE_URL env.
- `cacert_file` (String, Sensitive) Gitea TLS trusted CA cert.
- `debug` (Boolean) Indicates debug mode activation.
- `debug_log_path` (String) Path of file to write the debug logs to.
- `insecure` (Boolean) Disable SSL verification of API calls.
- `password` (String, Sensitive) Gitea password. Can also be set from GITEA_PASSWORD env.
- `token` (String, Sensitive) Gitea authentication token. Can also be set from GITEA_TOKEN env.
- `username` (String) Gitea user name. Can also be set from GITEA_USERNAME env.
