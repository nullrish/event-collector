# Event Collector

Event Collector is **Go-based analytics event ingestion service** inspired by systems
like PostHog, Segment and Mixpanel.

The goal of this project is to explore how high-throughput analytics pipelines
are designed — focusing on **concurrency, buffering, and efficient data storage**
rather than just basic CRUD APIs.

This repository is a **work in progress** and will gradually evolve into a fine grade
event collection and querying system.

## Project Goals

- Build a fast and non-blocking event ingestion API
- Use Go's concurrency primitives (goroutines, channels) effectively
- Design a buffered pipeline for handling high write throughput
- Store semi-structured events efficiently for later analysis
- Practice real-world backend engineering patterns

### Status

**🚧 Under active Development**

## Licensing

MIT - Do what ever you want with this project.
