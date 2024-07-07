#!/bin/bash

echo "-----------------------Project Setup-----------------------"
read -r -p "Enter MySQL username: " username
if [ -z "$username" ]
then
    echo "Username cannot be empty"
    exit 1
fi

read -r -s -p "Enter MySQL password: " password
echo
if [ -z "$password" ]
then
    echo "Password cannot be empty"
    exit 1
fi

read -r -s -p "Enter Secret Key for Hashing: " secret_key
echo
if [ -z "$secret_key" ]
then
    echo "Secret Key cannot be empty"
    exit 1
fi

echo "Setting up Database..."
mysql -u $username -p$password -e "DROP DATABASE IF EXISTS Library; CREATE DATABASE Library;"
echo "Database 'Library' created successfully"

echo "Setting up files..."
echo "Setting up .env..."
touch .env
echo "
DB_HOST='localhost'
DB_PORT='3306'
DB_USER='$username'
DB_PASSWORD='$password'
DB_NAME='Library'
SECRET_KEY='$secret_key'
" > .env
echo ".env setup complete"

echo "Setting up Makefile..."
mv Makefile.example Makefile
echo "
migrate_up:
		@read -p \"Enter version to migrate up: \" v; \
		migrate -path database/migrations -database \"mysql://$username:$password@(127.0.0.1:3306)/Library\" up \$\$v
migrate_down:
		@read -p \"Enter version to migrate down: \" v; \
		migrate -path database/migrations -database \"mysql://$username:$password@(127.0.0.1:3306)/Library\" down \$\$v
" >> Makefile
echo "Makefile setup complete"

go mod vendor
go mod tidy

echo "Files Setup Complete"
echo "-------------------Project Setup Complete-------------------"

