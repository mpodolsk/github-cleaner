# github-cleaner
Delete all unprotected branches on a repo.
Uses a github token for authentication.
Token needs repo read/write permissions.

Usage:
```bash
 GH_CLEANER_TOKEN="xxxxxx" GH_CLEANER_ORG="mpodolsk" GH_CLEANER_TARGET="some_repo" go run main.go
```
