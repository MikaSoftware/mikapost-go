# HOWTO: Setup ComicsCantian on DigitalOcean using CentOS 7 OS

## Instructions
### Setup Web-App Environment Variables
Developers Note: These steps are necessary if you want to do development on the remote server. If you do not want to do this then you can skip this section. If you would like to continue, please begin by running the following commands as the ``techops`` user account.

1. While being logged in as ``techops`` run the following:

    ```
    $ sudo vi /etc/profile.d/mikapost_env.sh
    ```


2. Copy and paste the following. **Please change the variables to meet your own.**

    ```
    #!/bin/bash
    export MIKAPOST_GORM_CONFIG="postgres://lucha:YOUR_PASSWORD@localhost/mikapost_db?sslmode=disable"
    export MIKAPOST_GORM_CONFIG="host=localhost port=5432 user=lucha dbname=mikapost_db password=YOUR_PASSWORD sslmode=disable"
    export MIKAPOST_SECRET="YOUR_SECRET_RANDOM_STRING"
    export MIKAPOST_ADDRESS="127.0.0.1:8080"  # Do not change!!!
    export MIKAPOST_UNIT_TEST_GORM_CONFIG="host=localhost port=5432 user=lucha dbname=mikapost_test_db password=YOUR_PASSWORD sslmode=disable"
    ```


3. Afterwords run the following.

    ```
    $ sudo chmod 0755 /etc/profile.d/mikapost_env.sh
    ```


4. On your next logon you'll have those environment variables loaded in. Feel free to log off and log back on. Afterwords to confirm the variables are loaded, run the following:

    ```
    $ sudo printenv

    OR

    $ sudo printenv MIKAPOST_GORM_CONFIG
    ```


### Setup Web-App Database

1. While being logged in as ``techops`` run the following:

    ```
    $ sudo -i -u postgres
    $ psql
    ```


2. Then run the following to support our web-applications database.

    ```sql
    drop database mikapost_db;
    create database mikapost_db;
    \c mikapost_db;
    CREATE USER lucha WITH PASSWORD 'YOUR_PASSWORD';
    GRANT ALL PRIVILEGES ON DATABASE mikapost_db to lucha;
    ALTER USER lucha CREATEDB;
    ALTER ROLE lucha SUPERUSER;
    CREATE EXTENSION postgis;
    ```


3. And re-run to support our database used for unit testing.

    ```sql
    drop database mikapost_test_db;
    create database mikapost_test_db;
    \c mikapost_test_db;
    CREATE USER lucha WITH PASSWORD 'YOUR_PASSWORD';
    GRANT ALL PRIVILEGES ON DATABASE mikapost_test_db to lucha;
    ALTER USER lucha CREATEDB;
    ALTER ROLE lucha SUPERUSER;
    CREATE EXTENSION postgis;
    ```


### Setup Web-App from GitHub
Please run the following commands as the ``lucha`` user account.

1. Get the project.

    ```
    $ go get github.com/luchacomics/mikapost-go
    ```


2. Install the dependencies.

    ```
    $ cd /opt/lucha/go/src/github.com/luchacomics/mikapost-go/;
    $ ./requirements.sh;
    ```


3. Build our project.

   ```
   $ go install github.com/luchacomics/mikapost-go
   ```


4. Enable permission and security while you are a ``techops`` user.

    ```
    $ sudo setcap 'cap_net_bind_service=+ep' /opt/lucha/go/bin/mikapost-go
    $ sudo setsebool -P httpd_can_network_connect 1
    $ sudo semanage permissive -a httpd_t
    $ sudo chcon -Rt httpd_sys_content_t /opt/django/workery-django/workery/static
    ```


### Integrate Nginx with Golang
Please run the following commands as the ``techops`` user account.

1. Load up ``Nginx``.

   ```
   $ sudo vi /etc/nginx/nginx.conf
   ```


2. Replace with the following code.

    ```
    server {
        listen       80;
        server_name  SERVER_DOMAIN_NAME_OR_IP;

        charset utf-8;

        location / {
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host            $http_host;

            proxy_pass http://127.0.0.1:8080;
        }
    }
    ```


3. Restart ``Nginx``.

    ```
    sudo systemctl restart nginx
    ```


4. Run your go app manually.

    ```
    # go run github.com/luchacomics/mikapost-go
    ```


5. Now in your browser go to ``http://SERVER_DOMAIN_NAME_OR_IP`` and you should see the app!


6. Special thanks to:

  * https://beego.me/docs/deploy/nginx.md


### Integrate Systemd with Golang

This section explains how to integrate our project with ``systemd`` so our operating system will handle stopping, restarting or starting.

1. **(OPTIONAL)** If you cannot access the server, please stop and review the steps above. If everything is working proceed forward.


2. While you are logged in as a ``techops`` user, please write the following into the console.

    ```
    $ sudo vi /etc/systemd/system/mikapost-go.service
    ```


3. Copy and paste the following. **Please change the variables to meet your own.** Please do not change the IP``127.0.0.1:8080``. Note: ``Nginx`` will be communicating with the ``golang`` app through this IP.

    ```
    Description=ComicsCantina Backend Webservice
    Wants=network.target
    After=network.target

    [Service]
    Environment=MIKAPOST_GORM_CONFIG=postgres://lucha:YOUR_PASSWORD@localhost/mikapost_db?sslmode=disable
    Environment=MIKAPOST_SECRET=YOUR_SECRET_RANDOM_STRING
    Environment=MIKAPOST_ADDRESS=127.0.0.1:8080
    Environment=MIKAPOST_GORM_CONFIG=postgres://lucha:YOUR_PASSWORD@localhost/mikapost_test_db?sslmode=disable
    Type=simple
    DynamicUser=yes
    WorkingDirectory=/opt/lucha/go/bin
    ExecStartPre=$MIKAPOST_GORM_CONFIG
    ExecStartPre=$MIKAPOST_SECRET
    ExecStartPre=$MIKAPOST_UNIT_TEST_GORM_CONFIG
    ExecStart=/opt/lucha/go/bin/mikapost-go
    Restart=always
    RestartSec=3
    SyslogIdentifier=mikapost_go

    [Install]
    WantedBy=multi-user.target
    ```


4. Grant access.

   ```
   $ sudo chmod 755 /etc/systemd/system/mikapost-go.service
   ```


5. (Optional) If you've updated the above, you will need to run the following before proceeding.

    ```
    $ sudo systemctl daemon-reload
    ```


6. We can now start the Gunicorn service we created and enable it so that it starts at boot:

    ```
    $ sudo systemctl start mikapost-go
    $ sudo systemctl enable mikapost-go
    ```


7. Confirm our service is running.

    ```
    $ systemctl status mikapost-go.service
    $ journalctl -f -u mikapost-go.service
    ```


8. And verify the URL works in the browser.

    ```text
    http://SERVER_DOMAIN_NAME_OR_IP/en/
    ```
