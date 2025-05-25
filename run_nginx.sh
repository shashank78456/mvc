#!/bin/bash

read -p "Do you want to host on nginx? (y/n) : " choice

if [[ "$choice" == "y" || "$choice" == "Y" ]]
then
    if command -v brew >/dev/null 2>&1
    then
        echo "Homebrew already installed"

        if brew list nginx >/dev/null 2>&1
        then
            echo "nginx already installed"
        else
            echo "nginx not installed. Installing nginx..."
            echo "-------------nginx Installation-------------"
            brew install nginx
            echo "-----------Installation Complete------------"
        fi

        echo "--------------Configuring nginx--------------"
        NGINX_PATH="/opt/homebrew/etc/nginx/"
        if [ ! -d "$NGINX_PATH/servers" ]
        then
            mkdir -p "$NGINX_PATH/servers"
        fi
        PROJECT_PATH=$(cd "$(dirname "$0")" && pwd)

        sudo bash -c 'cat > /opt/homebrew/etc/nginx/servers/mvc.conf <<EOL
server {
    listen 80;
    server_name mvc.sdslabs.local;

    root "'"$PROJECT_PATH"'";

    location / {
        proxy_pass http://127.0.0.1:3000;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
        
    location /static/ {
        alias "'"$PROJECT_PATH/static/"'";
    }

    # Deny access to hidden files
    location ~ /\. {
        deny all;
    }
}
EOL'
        grep -q "mvc.sdslabs.local" /etc/hosts || echo "127.0.0.1 mvc.sdslabs.local" | sudo tee -a /etc/hosts > /dev/null
        echo "---------------Configured nginx--------------"

        echo "Running on nginx Server"
        sudo nginx
        echo "Visit http://mvc.sdslabs.local in your browser"
        echo "Use 'sudo nginx -s stop' to stop nginx"
        echo "Control + C to stop server"
        go run ./cmd/main.go

    else
        echo "Homebrew not installed. Please install homebrew first."
    fi

elif [[ "$choice" == "n" || "$choice" == "N" ]]
then
    echo "Running Without nginx..."
    echo "Control + C to stop server"
    go run ./cmd/main.go

else
    echo "Please Enter a Valid Choice"
    exit 1
fi