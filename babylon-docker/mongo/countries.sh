#!/bin/bash
mongoimport --db babylon --collection countries --file /docker-entrypoint-initdb.d/countries.json
mongoimport --db babylon --collection languages --file /docker-entrypoint-initdb.d/languages-locale.json
mongoimport --db babylon --collection wage --file /docker-entrypoint-initdb.d/countries_minimum_wage.json
