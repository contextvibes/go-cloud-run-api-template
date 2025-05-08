# Roadmap: Go Starter Templates for Firebase Studio

This document outlines the vision and planned enhancements for the collection of Go Starter Templates designed for use with [Firebase Studio](https://studio.firebase.google.com/). Our goal is to provide a diverse set of high-quality, production-ready starting points for various Go applications.

This roadmap is a living document and will evolve based on community feedback and contributions.

## Guiding Principles for Templates

*   **Best Practices:** Templates should embody common Go best practices for project structure, code quality, testing, and deployment.
*   **Firebase Studio Integration:** Ensure seamless integration with Firebase Studio features, primarily through robust Nix environments (`.idx/dev.nix`) and accurate `idx-template.json` configurations.
*   **Developer Experience:** Templates must be easy to understand, use, and extend for developers starting new projects.
*   **Production-Readiness Hints:** While starters, they should guide towards production-quality patterns (e.g., effective configuration, structured logging, containerization where appropriate).
*   **Clarity and Focus:** Each template should serve a clear purpose and avoid unnecessary complexity.

## Near-Term Goals: New Templates & Enhancements

*   **[ ] New Template: Go CLI Application**
    *   **Description:** A starter for building command-line interface (CLI) applications in Go.
    *   **Target Features:**
        *   Argument parsing (e.g., using `cobra` or standard `flag` package).
        *   Example command structure.
        *   Basic project layout for internal logic and commands.
        *   Nix environment (`.idx/dev.nix`) with Go and potentially `delve` for debugging.
        *   Simple `README.md` explaining build and run instructions.
*   **[ ] New Template: Go gRPC Service**
    *   **Description:** A template for a Go-based gRPC service.
    *   **Target Features:**
        *   Example `.proto` definition for a simple service.
        *   Generated Go stubs for client and server.
        *   Basic server implementation.
        *   Example client to demonstrate usage.
        *   Nix environment with Go, `protobuf` compiler, and Go gRPC plugins (`protoc-gen-go`, `protoc-gen-go-grpc`).
        *   `README.md` with instructions for regenerating stubs and running the service.
*   **[ ] Enhancements to `cloud-run-api` Template:**
    *   **Optional Database Integration Example:**
        *   Research and potentially add a commented-out or optional module for connecting to a simple database (e.g., PostgreSQL or SQLite).
        *   Focus on showing patterns for configuration, connection management, and basic query execution.
        *   Clear documentation on how to enable and configure it.
    *   **Improved Unit/Integration Test Examples:**
        *   Add more comprehensive examples of unit tests, particularly for `internal/config` or any new utility functions.
        *   Expand integration tests to cover more edge cases or authentication/authorization if added.

## Medium-Term Goals: Expanding Template Variety

*   **[ ] New Template: Basic Go Web Application (Server-Side HTML)**
    *   **Description:** A template for a traditional web application serving HTML, primarily using Go's `html/template` package.
    *   **Target Features:**
        *   Routing using `net/http` or a lightweight router.
        *   Examples of HTML templates and passing data to them.
        *   Static asset serving (CSS, JS, images).
        *   Simple form handling example.
*   **[ ] New Template: Go Pub/Sub Event Worker**
    *   **Description:** A template for a Go application designed to process messages from a message queue like Google Cloud Pub/Sub or NATS.
    *   **Target Features:**
        *   Configuration for connecting to the message queue.
        *   Message handling logic with clear error handling and retries.
        *   Graceful shutdown mechanisms.
        *   Example Dockerfile for containerized deployment.
*   **[ ] New Template: Go Microservice with Inter-service Communication (e.g., HTTP or gRPC)**
    *   **Description:** A pair of simple Go microservices that communicate with each other, demonstrating basic patterns.
    *   **Target Features:**
        *   Two distinct Go modules/templates.
        *   Communication via either HTTP/REST or gRPC.
        *   Service discovery considerations (even if simple, like ENV vars).
        *   Dockerfiles for each service.

## Long-Term Vision for Templates

*   **Templates with Frontend Pairings:** Explore simple templates that pair a Go backend API with a very basic frontend (e.g., Go API + a minimal React, Vue, Svelte, or HTMX setup) to show full-stack potential.
*   **Domain-Specific Starters:** Consider templates for specific domains if there's community interest (e.g., "Go IoT Data Ingestor," "Go WebAssembly Plugin Host").
*   **Advanced Cloud Native Patterns:** Templates showcasing more advanced patterns like Observability (OpenTelemetry), advanced authentication (OAuth2/OIDC), or specific cloud service integrations in a more elaborate manner.
*   **Alternative IaC Options:** While many Go projects deploy as containers, explore if a template showing direct integration with an IaC tool (e.g., Go with Pulumi SDK for defining infrastructure) would be valuable.

## Contributing to the Roadmap

We welcome your ideas for new templates or improvements to existing ones!

1.  **Check Existing Issues & Discussions:** Someone might have already proposed something similar.
2.  **Open an Issue:** If not, please open a new issue on GitHub to discuss your idea. Provide details about the proposed template, its use case, and potential features.
3.  **Focus on Go:** While broader ecosystem tools are important, the primary focus of these templates is on the Go application itself.

Let's build a fantastic set of Go starter templates for Firebase Studio together!
