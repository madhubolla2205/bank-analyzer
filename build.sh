#!/bin/bash

echo "🔵 Tidying Go modules..."
go mod tidy

echo "🔵 Running Go server..."
go run cmd/main.go &

# Capture the server PID to stop later if needed
SERVER_PID=$!

# Wait a second for the server to start
sleep 2

echo "🔵 Opening browser at http://localhost:8080 ..."
# For Linux or Mac
if which xdg-open > /dev/null
then
    xdg-open http://localhost:8080
elif which open > /dev/null
then
    open http://localhost:8080
else
    echo "Please manually open: http://localhost:8080"
fi

# Wait until user presses [ENTER] to stop server
echo "🔴 Press [ENTER] to stop the server."
read

# Kill the Go server
kill $SERVER_PID
echo "🛑 Server stopped. Goodbye!"
