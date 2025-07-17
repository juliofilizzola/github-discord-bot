table "git_hub_events" {
  schema = schema.public
  column "id" {
    null = false
    type = text
  }
  column "id_pull_request" {
    null = true
    type = text
  }
  column "pull_request" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}
table "pull_requests" {
  schema = schema.public
  column "id" {
    null = false
    type = text
  }
  column "git_hub_id" {
    null = true
    type = bigint
  }
  column "number" {
    null = true
    type = bigint
  }
  column "title" {
    null = true
    type = text
  }
  column "state" {
    null = true
    type = text
  }
  column "html_url" {
    null = true
    type = text
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  column "merged_at" {
    null = true
    type = timestamptz
  }
  column "closed_at" {
    null = true
    type = timestamptz
  }
  column "head_ref" {
    null = true
    type = text
  }
  column "head_sha" {
    null = true
    type = text
  }
  column "base_ref" {
    null = true
    type = text
  }
  column "commits" {
    null = true
    type = bigint
  }
  column "additions" {
    null = true
    type = bigint
  }
  column "deletions" {
    null = true
    type = bigint
  }
  column "changed_files" {
    null = true
    type = bigint
  }
  column "repository_id" {
    null = true
    type = character_varying(64)
  }
  column "author_id" {
    null = true
    type = character_varying(64)
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "fk_pull_requests_author" {
    columns     = [column.author_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fk_pull_requests_repository" {
    columns     = [column.repository_id]
    ref_columns = [table.repositories.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  index "idx_pull_requests_git_hub_id" {
    unique  = true
    columns = [column.git_hub_id]
  }
}
table "repositories" {
  schema = schema.public
  column "id" {
    null = false
    type = text
  }
  column "git_hub_id" {
    null = true
    type = bigint
  }
  column "name" {
    null = true
    type = text
  }
  column "full_name" {
    null = true
    type = text
  }
  column "html_url" {
    null = true
    type = text
  }
  column "description" {
    null = true
    type = text
  }
  column "language" {
    null = true
    type = text
  }
  column "private" {
    null = true
    type = boolean
  }
  column "default_branch" {
    null = true
    type = text
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_repositories_git_hub_id" {
    unique  = true
    columns = [column.git_hub_id]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = text
  }
  column "git_hub_id" {
    null = true
    type = bigint
  }
  column "login" {
    null = true
    type = text
  }
  column "html_url" {
    null = true
    type = text
  }
  column "avatar_url" {
    null = true
    type = text
  }
  column "created_at" {
    null = true
    type = timestamptz
  }
  column "updated_at" {
    null = true
    type = timestamptz
  }
  primary_key {
    columns = [column.id]
  }
  index "idx_users_git_hub_id" {
    unique  = true
    columns = [column.git_hub_id]
  }
}
schema "public" {
  comment = "standard public schema"
}
