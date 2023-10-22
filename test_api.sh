#!/usr/bin/env bash
# This is a simple script to test the API.
# We will test your API by running a similar script.

# Your implementation should be running on localhost:8080.
BASE_URL="http://localhost:8080"

# 1. Test ping
curl -X POST "${BASE_URL}/ping" -H "accept: application/json" \
    -H "Content-Type: application/json" -d "{\"ping\":\"pong\"}"
