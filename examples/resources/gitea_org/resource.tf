resource "gitea_org" "org" {
  rule_name   = "branch-protection-rule-name"
  owner       = "owner-name"
  repo        = "repo-name"
  branch_name = "branch-name"
}