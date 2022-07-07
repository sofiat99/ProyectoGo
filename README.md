# Arquitectura Propuesta para Stratton Oakmont

The repository contains the following files

- propuesta_arquitectura.pdf

  Proposed architecture for Stratton Oakmont, It describes why this proposal is chosen and a diagram of the recommended architecture to implement.


To run the project perform the following steps, create a test database in MySQL

    CREATE DATABASE test;
    
Then in the internal/db/migrations folder you will find the 1_initial_schema.up.sql file which you will need to run to create the test table. Once the table is created, clone the project on your computer with the following command

    git clone https://github.com/sofiat99/ProyectoGo.git

locate ourselves in the project

    cd .\ProyectoGo\

Finally run the docker with the following command

    docker compose up
    
