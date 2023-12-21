FROM nginx:latest

COPY nginx/access.conf /etc/nginx/conf.d
COPY nginx/access_ssl.conf /etc/nginx/conf.d

COPY assets /usr/share/nginx/html/assets/
COPY *.html /usr/share/nginx/html/

RUN mkdir -p /etc/nginx/ssl 
 #&& curl https://ssl-config.mozilla.org/ffdhe2048.txt > /etc/nginx/ssl/dhparam.pem

RUN openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -subj \
    "/C=US/CN=localhost" \
    -keyout /tmp/access.baf.key -out /tmp/access.baf.crt && \
    ls -la /tmp && \
    cp /tmp/access.baf.key /etc/nginx/ssl/ && \
    cp /tmp/access.baf.crt /etc/nginx/ssl/ && \
    ls -la /etc/nginx/ssl