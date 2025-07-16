# FutureVolatileSignal: Real-time Cryptocurrency Volatility Surge Detection

This project provides a robust and efficient system for detecting real-time volatility surges in cryptocurrency markets. It leverages websocket-streamed order book data, Kalman filter-based anomaly scoring, and a modular Go architecture for performance and scalability. FutureVolatileSignal aims to empower traders and analysts with timely insights into potential market disruptions.

The core objective of FutureVolatileSignal is to identify periods of heightened price instability before they manifest as large price swings. This is achieved by continuously analyzing the order book depth and dynamics from cryptocurrency exchanges via websocket feeds. By monitoring bid-ask spreads, order size distributions, and order book velocity, the system calculates a volatility score. This score is then fed into a Kalman filter, which smooths the data and flags significant deviations from the expected volatility regime as anomalies, effectively signaling potential volatility surges. This approach allows for proactive risk management and informed decision-making in the volatile cryptocurrency market.

Unlike simple moving average-based volatility indicators, FutureVolatileSignal employs a Kalman filter to adapt to changing market conditions, reducing false positives and improving the accuracy of anomaly detection. The modular architecture of the Go codebase ensures that individual components, such as the websocket client, order book processor, and Kalman filter implementation, can be easily updated and maintained. The system is designed to be easily configurable, allowing users to tailor the anomaly detection sensitivity and Kalman filter parameters to their specific risk tolerance and trading strategy.

## Key Features

*   **Real-time Order Book Streaming:** Establishes and maintains persistent websocket connections to cryptocurrency exchanges to receive live order book updates. The implementation utilizes Go's concurrency primitives (goroutines and channels) for efficient data ingestion.
*   **Order Book Aggregation and Processing:** Parses and aggregates raw order book data from the websocket feed into a manageable representation. Calculates key order book metrics, such as weighted average price, bid-ask spread, and order book depth at multiple price levels.
*   **Kalman Filter-Based Anomaly Detection:** Implements a Kalman filter to smooth the volatility score time series and detect significant deviations from the expected volatility regime. The filter parameters (process noise, measurement noise) are configurable to adjust the sensitivity of the anomaly detection.
*   **Configurable Thresholding:** Provides customizable threshold parameters for flagging volatility surges based on the Kalman filter's state estimate and error covariance. Allows users to fine-tune the system's sensitivity to anomalies.
*   **Modular Architecture:** The Go codebase is structured into distinct modules, each responsible for a specific task (e.g., websocket communication, order book processing, Kalman filtering). This promotes code reusability, maintainability, and testability.
*   **Data Persistence (Optional):** The system can optionally persist processed order book data and anomaly scores to a database (e.g., PostgreSQL, InfluxDB) for historical analysis and backtesting.
*   **Alerting System (Optional):** Integrates with alerting services (e.g., Slack, email) to notify users of detected volatility surges in real-time.

## Technology Stack

*   **Go:** The primary programming language, chosen for its concurrency support, performance, and static typing.
*   **Gorilla Websocket:** A Go library for handling websocket connections, providing efficient and reliable communication with cryptocurrency exchanges.
*   **gonum/matrix:** A Go library for linear algebra operations, used in the Kalman filter implementation.
*   **cobra:** A Go library to build powerful command-line applications.
*   **viper:** A Go library for flexible configuration management.
*   **(Optional) PostgreSQL/InfluxDB:** Databases for persisting processed order book data and anomaly scores.
*   **(Optional) Alerting Services (Slack, email):** For real-time notifications of detected volatility surges.

## Installation

1.  Ensure you have Go installed (version 1.18 or later). Verify by running `go version`.
2.  Clone the repository: `git clone https://github.com/uhsr/FutureVolatileSignal.git`
3.  Navigate to the project directory: `cd FutureVolatileSignal`
4.  Download dependencies: `go mod download`
5.  Build the executable: `go build -o FutureVolatileSignal .`

## Configuration

The application utilizes environment variables for configuration. Create a `.env` file in the project root directory and define the following variables:

*   `EXCHANGE_WEBSOCKET_URL`: The websocket URL of the cryptocurrency exchange (e.g., "wss://stream.binance.com:9443/ws/btcusdt@depth@100ms").
*   `SYMBOL`: The trading pair to monitor (e.g., "BTCUSDT").
*   `KALMAN_PROCESS_NOISE`: The process noise parameter for the Kalman filter (e.g., "0.01").
*   `KALMAN_MEASUREMENT_NOISE`: The measurement noise parameter for the Kalman filter (e.g., "0.1").
*   `ANOMALY_THRESHOLD`: The threshold for flagging volatility surges (e.g., "2.0").
*   `(Optional) DATABASE_URL`: The connection string for the database (e.g., "postgres://user:password@host:port/database").
*   `(Optional) SLACK_WEBHOOK_URL`: The webhook URL for Slack notifications.

You can load these environment variables using a library like `godotenv`:
<pre>
import (
    "log"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    // ... rest of your code
}
</pre>

## Usage

Run the executable with the command: `./FutureVolatileSignal`

The application will connect to the specified websocket stream, process order book data, and continuously monitor for volatility surges. Detected anomalies will be printed to the console and optionally sent to the configured alerting service.

Further command line arguments can be added using cobra to specify custom websocket endpoints and configurations at runtime.

## Contributing

We welcome contributions to FutureVolatileSignal! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise code with thorough tests.
4.  Submit a pull request with a detailed description of your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/uhsr/FutureVolatileSignal/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to acknowledge the contributions of the open-source community, particularly the authors and maintainers of the Go libraries used in this project.