-- CREATE USER admin WITH PASSWORD 'admin';
-- CREATE DATABASE python_service;
-- CREATE DATABASE go_service;
-- GRANT ALL PRIVILEGES ON DATABASE python_service TO admin;
-- GRANT ALL PRIVILEGES ON DATABASE go_service TO admin;

#!/bin/bash
psql -U postgres <<-EOSQL
    CREATE USER admin WITH PASSWORD 'admin';
    CREATE DATABASE python_service;
    CREATE DATABASE go_service;
    GRANT ALL PRIVILEGES ON DATABASE python_service TO admin;
    GRANT ALL PRIVILEGES ON DATABASE go_service TO admin;
EOSQL