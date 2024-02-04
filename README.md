Followed along with https://www.youtube.com/watch?v=dpXhDzgUSe4

# RSS Aggregator

An RSS Aggregator built with Go, PostgreSQL, and concurrency.

## Overview

This RSS Aggregator allows users to follow different RSS feeds, fetch posts from all the feeds they follow, and add new RSS feeds to the database. The project is implemented in Go, leveraging PostgreSQL for data storage and employing Go routines for concurrency.

## Features

- **Add New Feeds:** Users can add new RSS feeds to the database.
- **Feed Management:** Users can follow and unfollow RSS feeds.
- **Post Retrieval:** Users can fetch posts from all the feeds they follow.

## Prerequisites

- Go installed (along with goose for go)
- PostgreSQL installed and running

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/rss-aggregator.git
   cd rss-aggregator
   ```
2. Setup your databse:
   
   POSTGRES_DATABASE_CONNECTION_STRING example - postgres://postgres:root@localhost:5432/rssagg
   ```
   goose postgres YOUR_POSTGRES_DATABASE_CONNECTION_STRING up
   ```
4. Build and Run:
   ```
   go build; go build; .\rss-aggregator.exe
   ```
