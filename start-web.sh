#!/bin/bash
# start-web.sh

echo "Waiting for the database service to be ready..."
sleep 5  # Wait for 30 seconds

echo "Starting web service..."
# Command to start your web application, e.g.,:
exec ./main  # Replace with your web app's start command
