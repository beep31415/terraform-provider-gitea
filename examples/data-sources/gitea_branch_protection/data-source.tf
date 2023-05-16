data "gitea_branch_protection" "branch-protection" {
  rule_name = "branch-protection-rule-name"
  owner     = "owner-name"
  repo      = "repo-name"
}