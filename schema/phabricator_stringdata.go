// Code generated by stringdata. DO NOT EDIT.

package schema

// PhabricatorSchemaJSON is the content of the file "phabricator.schema.json".
const PhabricatorSchemaJSON = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "phabricator.schema.json#",
  "title": "Phabricator Connection",
  "description": "Configuration settings for Phabricator",
  "type": "object",
  "additionalProperties": false,
  "properties": {
    "url": {
      "description": "URL of a Phabricator instance, such as https://phabricator.example.com",
      "type": "string"
    },
    "token": {
      "description": "API token for the Phabricator instance.",
      "type": "string"
    },
    "repos": {
      "description": "The list of repositories available on Phabricator.",
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["path", "callsign"],
        "properties": {
          "path": {
            "description": "Display path for the url e.g. gitolite/my/repo",
            "type": "string"
          },
          "callsign": {
            "description": "The unique Phabricator identifier for the repository, like 'MUX'.",
            "type": "string"
          }
        }
      }
    }
  }
}
`
