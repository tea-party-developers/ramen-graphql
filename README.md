# ramen-graphql

Graphql API for Ramen SNS

## 1. Table of Contents

- [1. Table of Contents](#1-table-of-contents)
- [2. About This Repository](#2-about-this-repository)
- [3. Usage](#3-usage)
  - [a. Create a `.env` file](#a-create-a-env-file)
  - [b. Run servers](#b-run-servers)
  - [c. Access GraphQL Playground](#c-access-graphql-playground)

## 2. About This Repository

This is a repository for a GraphQL API that returns Ramen SNS data,
created with [gqlgen](https://github.com/99designs/gqlgen) and [Bun](https://github.com/uptrace/bun).

## 3. Usage

### a. Create a `.env` file

Create a `.env` file based on `.env.tmp`.
※ You can use `.env.tmp` by copying it and renaming it to `.env`.

### b. Run servers

```sh
docker compose up --build
```

### c. Access GraphQL Playground

Go to Playground: [http://localhost:8080](http://localhost:8080)
