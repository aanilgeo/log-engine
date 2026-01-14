# AI-Powered Log Engine (Go + gRPC + Gemini)

A high-performance distributed logging system built in **Go**. This engine ingests logs via **gRPC**, stores them in high-efficiency **segmented files**, and utilizes **Gemini 2.5 Flash** for automated root-cause analysis and system health summaries.

## 🚀 Key Features
- **gRPC Streaming:** Low-latency log ingestion using Protocol Buffers.
- **Custom Storage Engine:** Implements append-only segment files with manual rotation logic—mimicking the architecture of high-throughput systems like Apache Kafka.
- **AI-Driven Observability:** Integrated RAG (Retrieval-Augmented Generation) pipeline that provides natural language summaries of system errors.
- **Containerized Build:** Fully Dockerized for reproducible deployments across any environment.



## 🛠 Tech Stack
- **Language:** Go 1.21+
- **Communication:** gRPC / Protocol Buffers
- **AI:** Google Gemini 2.5 Flash API
- **DevOps:** Docker, Makefile

## 🚦 Quick Start

### 1. Prerequisites
- Docker & Docker Compose
- Gemini API Key ([Get one here](https://aistudio.google.com/))

### 2. Run with Docker
```bash
export GEMINI_API_KEY="your_key_here"
docker-compose up --build
```

### 3. Run a Test Client
Once your server is running (either via Docker or **go run**), open a second terminal window on your Mac. This client script simulates service logs being sent to the engine and then triggers the AI summarization.
```bash
# Navigate to your project root
cd log-engine

# Run the client to ingest logs and fetch the AI summary
go run cmd/client/main.go
```