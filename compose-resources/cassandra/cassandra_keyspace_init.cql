CREATE KEYSPACE IF NOT EXISTS microservpoc WITH REPLICATION = {
    'class': 'SimpleStrategy', 
    'replication_factor' : 1
    };

CREATE TABLE IF NOT EXISTS microservpoc.users (
    id UUID, 
    firstname text, 
    lastname text, 
    age int, 
    email text, 
    city text, 
    PRIMARY KEY (id)
    );
