# Get GitHub Issues (`gghi`)
## Description
CLI tool get issues from a repository (more or less a subset of the `issues` endpoint of the GitHub cli `gh`)

## Features
- Auth 
  - Basic Auth (stored in ~/)
  - OAuth (Web flow?)
- Issues
  - Get x issues of a repo with filters y
  - Get issues of all repositories of a user (Number of issues and filters per repo or general)
    - Has to be done with concurrency otherwise might be too slow
- Caching of requests (Limited to 5000 per hour)

## CLI
- auth
  - status
  - login
  - logout
- issues
  - get
    - -filters

## License
- MIT
