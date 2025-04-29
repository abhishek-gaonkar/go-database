# Build Your Own Database in Go (From Scratch)

This repository is an implementation of a simple database engine built from scratch using the Go programming language, based on the principles taught in the book _"Build Your Own Database in Go from Scratch"_ by James Smith. The goal of this project is to understand the inner workings of databases by building a basic key-value store and eventually implementing more advanced features such as indexing, queries, and transactions.

## Project Overview

This project is a Go-based implementation of a custom database engine that allows you to store, retrieve, and manipulate data. It aims to replicate basic database features, such as:

- Key-Value Storage
- Indexing
- Query Execution
- Transactions

The database engine is modular, allowing you to build upon each core component to eventually support more advanced features like multi-table joins, ACID transactions, and more.

## Features

- **Key-Value Store:** A simple in-memory key-value store.
- **B+ Tree Indexing:** Efficiently store and search for keys with B+ tree indexing.
- **Query Execution:** Execute basic SELECT queries.
- **Transactions:** Implement simple transaction support to ensure data consistency.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go 1.x](https://golang.org/dl/) — The Go programming language.
- [Git](https://git-scm.com/) — For version control.

### Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/abhishek-gaonkar/go-database.git
   ```

2. Change to the project directory:

   ```bash
   cd go-database
   ```

3. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

### Running the Project

To run the database engine, simply build and execute the Go code:

```bash
go run main.go
```
