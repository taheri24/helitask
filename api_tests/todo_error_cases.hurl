# Helitask Todo API negative paths

[Options]
variable base_url = http://localhost:8080/api/v0

# Attempt to create todo with missing fields
POST {{base_url}}/todo/
[Headers]
Content-Type: application/json
Accept: application/json

{
  "description": ""
}

HTTP 400
[Asserts]
jsonpath "$.error" contains "Invalid input"

# Request todo with invalid UUID format
GET {{base_url}}/todo/not-a-uuid
[Headers]
Accept: application/json

HTTP 400
[Asserts]
jsonpath "$.error" contains "Invalid UUID"

# Request todo with non-existing UUID
GET {{base_url}}/todo/11111111-1111-1111-1111-111111111111
[Headers]
Accept: application/json

HTTP 404
[Asserts]
jsonpath "$.error" contains "record not found"
