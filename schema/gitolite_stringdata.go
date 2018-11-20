// Code generated by stringdata. DO NOT EDIT.

package schema

// GitoliteSchemaJSON is the content of the file "gitolite.schema.json".
const GitoliteSchemaJSON = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "gitolite.schema.json#",
  "title": "Gitolite Connection",
  "description": "Configuration settings for Gitolite ",
  "type": "object",
  "additionalProperties": false,
  "required": ["prefix", "host"],
  "properties": {
    "prefix": {
      "description":
        "Repository name prefix that will map to this Gitolite host. This should likely end with a trailing slash. E.g., \"gitolite.example.com/\".\n\nIt is important that the Sourcegraph repository name generated with this prefix be unique to this code host. If different code hosts generate repository names that collide, Sourcegraph's behavior is undefined.",
      "type": "string"
    },
    "host": {
      "description": "Gitolite host that stores the repositories (e.g., git@gitolite.example.com).",
      "type": "string"
    },
    "blacklist": {
      "description":
        "Regular expression to filter repositories from auto-discovery, so they will not get cloned automatically.",
      "type": "string"
    },
    "phabricatorMetadataCommand": {
      "description":
        "Bash command that prints out the Phabricator callsign for a Gitolite repository. This will be run with environment variable $REPO set to the name of the repository and used to obtain the Phabricator metadata for a Gitolite repository. (Note: this requires ` + "`" + `bash` + "`" + ` to be installed.)",
      "type": "string"
    }
  }
}
`
