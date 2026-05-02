# github.com/h00s/bitbox

Raptor web API project.

## Getting Started

```bash
raptor dev
```

The server starts at http://localhost:3000.

## API Endpoints

### GET /api/v1/hello

Returns a greeting message and a list of all added greetings.

**Response:**
```json
{
  "message": "Hello, World!",
  "greetings": []
}
```

### POST /api/v1/hello

Adds a new greeting to the list.

**Request body:**
```json
{
  "greeting": "Hi there!"
}
```

**Response:** `201 Created`
