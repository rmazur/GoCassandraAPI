#!/bin/bash
#   Use this script to initialise Cassandra DB

/init/scripts/wait-for-it.sh -t 0 cassandraapi:9042 -- echo "CASSANDRA Node started"

cqlsh -f /init/scripts/cassandra_keyspace_init.cql cassandraapi

echo "### CASSANDRA INITIALISED! ###"