# HOWTO: Setup Lets Encrypt for MikaPost Golang Web-App on DigitalOcean using CentOS 7 OS
## Description
This article assumes you've completed the first two articles with setting up MikaPost. These instructions were modified from [DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-centos-7).

## Instruction
The following instructions are used to manually setup `letsencrypt` and automatically integrate with `nginx`.

1. Open up the ``nginx`` configuration file.

  ```
  $ sudo vi /etc/nginx/nginx.conf
  ```


2. And replace the following:

  ```
  server_name SERVER_DOMAIN_NAME_OR_IP;
  ```


3. With the following:

  ```
  server_name mikapost.com www.mikapost.com;
  ```


4. Confirm the ``nginx`` app works and restart it.

  ```
  $ sudo nginx -t;
  $ sudo systemctl reload nginx;
  ```


5. Install our **Lets Encrypt** client.

  ```
  $ sudo yum install -y certbot-nginx
  ```


6. Generate our certificate.

  ```
  $ sudo certbot --nginx -d mikapost.com -d www.mikapost.com
  ```


7. Follow the instructions and choose the most appropriate options.


8. **(Optional)** Please make a copy of the ``/etc/letsencrypt`` file.


9. Restart ``nginx``.

  ```
  $ sudo systemctl restart nginx
  ```


10. Upgrade the security by following the instructions - https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-centos-7


11. Restart the server.

    ```
    $ sudo systemctl restart nginx
    ```


12. Would you like to know more?


### HOW DO WE AUTO RENEW?
https://certbot.eff.org/lets-encrypt/centosrhel7-nginx.html

sudo crontab -e

# Add this to the crontab and save it:
0 0,12 * * * python -c 'import random; import time; time.sleep(random.random() * 3600)' && /usr/bin/certbot renew && systemctl restart nginx


### Nginx + SSL
If your SSL is not being populated at your address then follow these.
