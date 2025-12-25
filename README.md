# Microservices System

A scalable backend built with **Golang**, utilizing a microservices architecture. This project demonstrates how to decouple services using **gRPC** for internal communication and **GraphQL** as the public-facing API gateway.

## 🏗 Architecture

The system consists of the following services:

| Service             | Description                                                              | Tech Stack              | Port (Internal)           |
| :------------------ | :----------------------------------------------------------------------- | :---------------------- | :------------------------ |
| **GraphQL Gateway** | Entry point for client applications. Aggregates data from microservices. | `99designs/gqlgen`      | `8080` (mapped to `8000`) |
| **Account Service** | Manages user registration and authentication.                            | Go, PostgreSQL, gRPC    | `8080`                    |
| **Catalog Service** | Manages products and search functionality.                               | Go, Elasticsearch, gRPC | `8080`                    |
| **Order Service**   | Handles order processing and history.                                    | Go, PostgreSQL, gRPC    | `8080`                    |

### Infrastructure

-   **PostgreSQL**: Primary database for Account and Order services.
-   **Elasticsearch**: Search engine for the Catalog service.
-   **Docker Compose**: Orchestrates all services and databases for local development.

## 🛠 Technology Stack

-   **Language**: [Go (Golang)](https://go.dev/) `1.25.5`
-   **API Specification**: [GraphQL](https://graphql.org/)
-   **RPC Framework**: [gRPC](https://grpc.io/)
-   **Libraries**:
    -   `99designs/gqlgen`: GraphQL server library
    -   `segmentio/ksuid`: K-Sortable Unique IDs
    -   `olivere/elastic`: Elasticsearch client
-   **Databases**: PostgreSQL `17`, Elasticsearch `9.2.2`

## 🚀 Getting Started

### Prerequisites

-   [Docker Desktop](https://www.docker.com/products/docker-desktop)
-   [Go](https://go.dev/dl/) (Optional, if you want to run services individually)

### Running the Application

1. **Clone the repository** (if you haven't already):

    ```bash
    git clone https://github.com/rajan-marasini/ecom-microservice.git
    cd ecom-microservice
    ```

2. **Start all services** using Docker Compose:

    ```bash
    docker-compose up --build
    ```

    _Note: The first build may take a few minutes as it downloads dependencies and docker images._

3. **Access the GraphQL Playground**:
   Open your browser and navigate to:
   [http://localhost:8000/playground](http://localhost:8000/playground)

### Sample Queries

Here are some example queries to test the system in the GraphQL Playground:

**Create Account**

```graphql
mutation {
    createAccount(input: { name: "Alice", password: "password123" }) {
        id
        name
    }
}
```

**Search Products**

```graphql
query {
    products(query: "", skip: 0, take: 10) {
        id
        name
        description
        price
    }
}
```

**Place Order**

```graphql
mutation {
    createOrder(
        input: {
            accountId: "<ACCOUNT_ID_FROM_ABOVE>"
            products: [{ id: "<PRODUCT_ID>", quantity: 1 }]
        }
    ) {
        id
        totalPrice
    }
}
```

## 📂 Project Structure

```
├── account/         # Account microservice (User management)
├── catalog/         # Catalog microservice (Product search & management)
├── order/           # Order microservice (Order processing)
├── graphql/         # GraphQL Gateway (API Layer)
├── docker-compose.yml # Docker orchestration config
├── go.mod           # Go module definition
└── README.md        # Project documentation
```
