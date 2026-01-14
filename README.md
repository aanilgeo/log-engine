# Distributed Log Engine with AI Insights

A high-performance distributed logging system built in **Go**. This engine ingests logs via **gRPC**, stores them in high-efficiency **segmented files**, and utilizes **Gemini 2.5 Flash** for automated root-cause analysis and system health summaries.

## Key Features
- **gRPC Streaming:** Low-latency log ingestion using Protocol Buffers.
- **Custom Storage Engine:** Implements append-only segment files with manual rotation logic.
- **AI-Driven Observability:** Integrated RAG (Retrieval-Augmented Generation) pipeline that provides natural language summaries of system errors.
- **Containerized Build:** Fully Dockerized for reproducible deployments across any environment.

## Tech Stack
- **Language:** Go 1.21+
- **Communication:** gRPC / Protocol Buffers
- **AI:** Google Gemini 2.5 Flash API
- **DevOps:** Docker, Makefile

## Quick Start

### 1. Prerequisites
- Docker and Docker Compose
- Gemini API Key

### 2. Run with Docker
```bash
export GEMINI_API_KEY="your_key_here"
docker-compose up --build
```
### 3. Run a Test Client
Once the server is running, open a second terminal window to simulate log ingestion and trigger AI summarization.
```bash
# Navigate to your project root
cd log-engine

# Run the client to ingest logs and fetch the AI summary
go run cmd/client/main.go
```