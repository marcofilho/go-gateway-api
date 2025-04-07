## Payment Gateway

- A complete Payment Gateway using microservices architecture.

- The goal is to demonstrate the construction of a modern distributed system, with separation of concerns, asynchronous communication, and real-time fraud analysis.

## Technologies Used

- **Next.js**: React framework for building the Frontend.
- **Go**: Programming language used for developing the Gateway.
- **Apache Kafka**: Streaming platform for asynchronous communication.
- **Nest.js**: Node.js framework for the Anti-Fraud service.

## How to Run the Project

1. Clone the repository:
    ```bash
    git clone https://github.com/marcofilho/go-gateway-api.git
    cd go-gateway-api
    ```

2. Set up the environment variables for each service:
    - Frontend: `.env.local`
    - Gateway: `.env`
    - Anti-Fraud: `.env`

3. Start the services using Docker Compose:
    ```bash
    docker-compose up --build
    ```

4. Access the services:
    - Frontend: [http://localhost:3000](http://localhost:3000)
    - Gateway API: [http://localhost:8080](http://localhost:8080)

## Repository Structure

```plaintext
/go-gateway-api
├── frontend/       # Frontend code (Next.js)
├── gateway/        # Gateway code (Go)
├── antifraude/     # Anti-Fraud service code (Nest.js)
├── kafka/          # Apache Kafka configuration
└── docker-compose.yml # Docker configuration for running the services
```

## Contribution

Contributions are welcome! Follow the steps below:

1. Fork the repository.
2. Create a branch for your feature or fix:
    ```bash
    git checkout -b my-feature
    ```
3. Commit your changes:
    ```bash
    git commit -m "Description of my feature"
    ```
4. Push to the remote repository:
    ```bash
    git push origin my-feature
    ```
5. Open a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## System Components

### Frontend (Next.js)

- User interface for account management and payment processing.
- Built with Next.js to ensure performance and a great user experience.

### Gateway (Go)

- Main payment processing system.
- Manages accounts, transactions, and coordinates payment flows.
- Publishes transaction events to Kafka for fraud analysis.

### Apache Kafka

- Handles asynchronous communication between the API Gateway and Anti-Fraud service.
- Ensures reliable message exchange between services.
- Specific topics for transactions and analysis results.

### Anti-Fraud (Nest.js)

- Consumes transaction events from Kafka.
- Performs real-time analysis to identify potential fraud.
- Publishes analysis results back to Kafka.

## Communication Flow

1. The Frontend sends requests to the API Gateway via REST.
2. The Gateway processes the requests and publishes transaction events to Kafka.
3. The Anti-Fraud service consumes the events and performs real-time analysis.
4. Analysis results are published back to Kafka.
5. The Gateway consumes the results and finalizes the transaction processing.