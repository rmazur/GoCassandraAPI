#!/bin/bash

echo "Minimum Go container is running..."
echo "CTRL-C to stop it."

docker run -p 6060:8080 --name GoCassandraPAI --rm gocassandraapi
