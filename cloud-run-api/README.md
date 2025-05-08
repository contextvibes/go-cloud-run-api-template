# Go Hello World API Template (Cloud Run)

This Go application template provides a starting point for a "Hello World" API, suitable for deployment on Google Cloud Run or similar containerized environments. It features a basic project structure, configuration loading from environment variables, structured JSON logging with `log/slog` (via `dui-go`), HTTP routing with middleware, and example tests.

This project uses `your-module-name` as its initial Go module path. **You must change this to your own desired module path.**

The development environment within Firebase Studio comes pre-configured with the `contextvibes` CLI, available at `./bin/contextvibes`. This CLI can help streamline certain development tasks (refer to its documentation for available commands).

## Quick Start & Setup

1.  **Initialize Your Go Module:**
    *   Decide on your Go module path (e.g., `github.com/your-username/my-cool-api`, or `my-internal-project/service-x`).
    *   In the `go.mod` file, change `module your-module-name` to your chosen module path.
    *   Using your IDE's search and replace, or a tool like `sd` or `rg`, replace all occurrences of `your-module-name/internal` with `your-chosen-module-path/internal` across all `.go` files in this project.
        *Example using `sd` (if installed): `sd 'your-module-name/internal' 'your-chosen-module-path/internal' $(git ls-files '*.go')`*
    *   Run `go mod tidy` to update dependencies.

2.  **Configuration:**
    *   Copy `.env.example` to `.env`.
    *   Update `.env` with your settings, especially `GOOGLE_CLOUD_PROJECT`.
    *   The application loads configuration from environment variables (see `internal/config/config.go`).

3.  **Install Dependencies:**
    If not automatically handled by the initial Firebase Studio setup:
    ```bash
    go mod download
    go mod tidy
    ```

## Features

*   **HTTP Endpoints:**
    *   `GET /hello`: Returns a JSON response with a "Hello, World!" message, service name, and timestamp.
    *   `POST /echo`: Expects a JSON payload `{"text_to_echo": "your message"}` and returns a JSON response echoing the message.
    *   `GET /healthz`: A simple health check endpoint returning "ok".
    *   `GET /`: A root endpoint with a welcome message.
*   **Structured Logging:** Uses `log/slog` via `dui-go/logging/cloudlogging` for GCP-compatible JSON logs, including trace correlation from the `X-Cloud-Trace-Context` header. Log level is configurable.
*   **Configuration:** Loads settings via environment variables using `dui-go/env`.
*   **Middleware:** Includes middleware for Cloud Trace context propagation provided by `dui-go/logging/cloudlogging`.
*   **Testing:** Demonstrates unit tests (using `testify`) and basic integration tests.
*   **Containerized:** Includes a `Dockerfile` and `.dockerignore` for building a minimal distroless image.
*   **Development Environment:** Configured for Firebase Studio using `.idx/dev.nix`, which provides Go tools, `golangci-lint`, the `contextvibes` CLI, and other utilities.

## Prerequisites

1.  **Go:** Version **1.24** or higher (see `go.mod`).
2.  **Docker:** Required for building and running the container image locally.
3.  **GCP Project:** (Required for full Cloud Logging trace correlation) Access to a GCP project. The `GOOGLE_CLOUD_PROJECT` environment variable must be set.
4.  **Authentication (for local GCP logging):** If running locally and wanting logs to correlate in GCP Console:
    ```bash
    gcloud auth application-default login
    ```

## Configuration Variables

*   `GOOGLE_CLOUD_PROJECT`: (Required) Your GCP Project ID.
*   `PORT`: (Optional) Defaults to `8080`.
*   `LOG_LEVEL`: (Optional) `DEBUG`, `INFO`, `WARN`, `ERROR`, etc. Defaults to `INFO`. Case-insensitive.
*   `API_SERVICE_NAME`: (Optional) Defaults to `go-hello-world-api`.

## Input/Output Payloads

### `GET /hello`
*   **Request:** No payload.
*   **Response (200 OK):**
    ```json
    {
      "message": "Hello, World from go-hello-world-api!",
      "timestamp": "2023-10-27T10:00:00.123456789Z"
    }
    ```

### `POST /echo`
*   **Request Body:**
    ```json
    {
      "text_to_echo": "Your message here"
    }
    ```
*   **Response (200 OK):**
    ```json
    {
      "received_text": "Your message here",
      "reply": "Service 'go-hello-world-api' received your message: 'Your message here'",
      "timestamp": "2023-10-27T10:01:00.123456789Z"
    }
    ```
*   **Response (400 Bad Request):** If `text_to_echo` is missing or empty, or if JSON is malformed.


## Development Workflow (using Go & Docker tools)

The `contextvibes` CLI is installed at `./bin/contextvibes` in your Firebase Studio workspace and may offer additional helpful commands (check its documentation). The following are standard Go and Docker commands:

*   **Format Code:**
    ```bash
    go fmt ./...
    ```
*   **Lint Code:**
    The Firebase Studio environment includes `golangci-lint`.
    ```bash
    # Basic checks:
    go vet ./...
    # Comprehensive linting:
    golangci-lint run ./...
    # Consider adding a .golangci.yml for custom rules
    ```
*   **Run Tests:**
    ```bash
    go test ./...
    ```
*   **Build Binary:**
    ```bash
    go build -o ./bin/app ./cmd/main.go
    ```
*   **Run Locally:**
    Ensure your `.env` file is configured or export necessary environment variables (especially `GOOGLE_CLOUD_PROJECT`).
    ```bash
    # Example:
    # export GOOGLE_CLOUD_PROJECT="your-gcp-project-id"
    go run ./cmd/main.go
    ```
*   **Build Docker Image:**
    ```bash
    docker build -t your-api-image-name .
    ```
*   **Run Docker Container Locally:**
    ```bash
    # Ensure required env vars like GOOGLE_CLOUD_PROJECT are set in your .env file
    docker run -p 8080:8080 --env-file .env your-api-image-name
    ```
    (Note: Adjust port mapping `-p` if needed.)


## Example Usage (Curl)

Assuming the service is running locally on port `8080`:

**1. GET /hello:**
```bash
curl -i http://localhost:8080/hello
```

**2. POST /echo:**
```bash
curl -X POST http://localhost:8080/echo \
-H "Content-Type: application/json" \
-i \
-d '{
  "text_to_echo": "Greetings from curl!"
}'
```

**3. GET /healthz:**
```bash
curl -i http://localhost:8080/healthz
```

**4. GET /:**
```bash
curl -i http://localhost:8080/
```

## Deployment (Example: Cloud Run)

1.  **Build & Push Image:**
    Follow standard procedures to build the Docker image (e.g., `docker build -t your-image-name .`) and push it to a registry like Google Artifact Registry.
    ```bash
    # Example using gcloud and Artifact Registry
    export PROJECT_ID="your-gcp-project-id"
    export REGION="your-gcp-region"
    export REPO_NAME="your-artifact-registry-repo"
    export IMAGE_NAME="go-hello-world-api" # Or your chosen name
    export IMAGE_TAG="latest" # Or a specific version like "v1.0.0"

    gcloud auth configure-docker ${REGION}-docker.pkg.dev
    docker build -t ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO_NAME}/${IMAGE_NAME}:${IMAGE_TAG} .
    docker push ${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO_NAME}/${IMAGE_NAME}:${IMAGE_TAG}
    ```

2.  **Deploy to Cloud Run:**
    ```bash
    gcloud run deploy ${IMAGE_NAME} \
      --image=${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPO_NAME}/${IMAGE_NAME}:${IMAGE_TAG} \
      --platform=managed \
      --region=${REGION} \
      --allow-unauthenticated \ # OR --no-allow-unauthenticated and configure IAM invoker
      --service-account=your-runtime-sa@${PROJECT_ID}.iam.gserviceaccount.com \
      --set-env-vars=GOOGLE_CLOUD_PROJECT=${PROJECT_ID},LOG_LEVEL=INFO \
      # --set-env-vars=API_SERVICE_NAME=my-prod-hello-world \ # Optional override
      --port=8080 # Port exposed in Dockerfile
    ```

3.  **IAM Permissions:**
    The runtime service account (`your-runtime-sa@...`) needs:
    *   `roles/logging.logWriter` (for Cloud Logging, if GOOGLE_CLOUD_PROJECT is set)
    *   (If `--no-allow-unauthenticated`) Invokers need `roles/run.invoker`.

## Error Handling

*   **HTTP Method Not Allowed (405):** Returned if an incorrect HTTP method is used for an endpoint (e.g., POST to `/hello`).
*   **Bad Request (400):**
    *   Malformed JSON payload for `/echo`.
    *   Missing or empty `text_to_echo` field in `/echo` request.
*   **Not Found (404):** For undefined paths.

## TODO / Future Work

*   Implement graceful shutdown in `cmd/main.go` for the HTTP server.
*   Add more example endpoints showcasing different patterns (see root `ROADMAP.md`).
*   Enhance integration tests with more scenarios.
*   Ensure `contextvibes` CLI commands are well-documented if specific workflow integrations are added to this README in the future.