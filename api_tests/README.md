# Helitask API Integration Tests (Hurl)

These scripts exercise the Helitask REST API using the [Hurl](https://hurl.dev) declarative test language. Each `.hurl` file
follows the same request / response layout as the examples below:

```hurl
GET https://example.org/api
HTTP 200
[Asserts]
jsonpath "$.slideshow.title" == "A beautiful ✈!"
```

```hurl
POST https://example.org/api
{
  "jsonKey": 1
}

HTTP 200
[Asserts]
jsonpath "$.title" == "Text"
```

The provided suites focus on the `/api/v0/todo` endpoints:

- `todo_success.hurl` covers the happy path of creating an item and retrieving it with the captured identifier.
- `todo_error_cases.hurl` validates common error responses such as validation failures, malformed identifiers, and missing
  records.

> **Tip:** Run the Helitask server locally (defaults to `http://localhost:8080`) before executing the suites:
>
> ```bash
> APP_ENV=development go run ./main.go
> hurl --test api_tests/todo_success.hurl
> hurl --test api_tests/todo_error_cases.hurl
> ```
