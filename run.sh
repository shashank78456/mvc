read -p "Do you want to host on Apache Server? (y/n) : " choice

if [[ "$choice" == "y" || "$choice" == "Y" ]]
then
    read -p "Enter your E-mail ID: " email
    if [ -z "$email" ]
    then
        echo "E-mail ID cannot be empty"
        exit -1
    elif [ ! echo "$email" | grep -q "@" ]
    then
        echo "Please Enter Valid E-mail ID"
        exit -1
    fi

    if [[ command -v apache2 >/dev/null 2>&1 ]]
    then
        echo "Apache already installed"
    else
        echo "-------------Apache Installation-------------"
        sudo apt install apache2 -y
        echo "------------Installation Complete------------"
    fi

    echo "--------------Configuring Apache--------------"
    sudo a2enmod proxy proxy_http
    sudo bash -c 'cat >> /etc/apache2/sites-available/mvc.sdslabs.local.conf <<EOL
    <VirtualHost *:80>
        ServerName mvc.sdslabs.local
        ServerAdmin $email
        ProxyPreserveHost On
        ProxyPass / http://127.0.0.1:8000/
        ProxyPassReverse / http://127.0.0.1:8000/
        TransferLog /var/log/apache2/mvc_access.log
        ErrorLog /var/log/apache2/mvc_error.log
    </VirtualHost>
    EOL'

    sudo a2ensite /etc/apache2/sites-available/mvc.sdslabs.local.conf
    echo "127.0.0.1 mvc.sdslabs.local" | sudo tee -a /etc/hosts > /dev/null
    sudo a2dissite /etc/apache2/sites-available/000-default.conf
    sudo apache2ctl configtest
    sudo systemctl restart apache2
    sudo systemctl status apache2
    echo "---------------Configured Apache--------------"

    echo "Running on Apache Server..."
    echo "Control + C to stop server"
    go run ./cmd/main.go

elif [[ "$choice" == "n" || "$choice" == "N" ]]
then
    echo "Running Without Apache Server..."
    echo "Control + C to stop server"
    go run ./cmd/main.go
else
    echo "Please Enter a Valid Choice"
    exit 1
fi
