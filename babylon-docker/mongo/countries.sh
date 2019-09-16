#!/bin/bash
mongoimport --db babylon --collection countries --file /docker-entrypoint-initdb.d/countries.json