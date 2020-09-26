FROM php:7.2.0-fpm-alpine

RUN apk add --update-cache libzip-dev && \
    docker-php-ext-install zip && \
    rm -rf /var/cache/apk/*

ADD www.conf /usr/local/etc/php-fpm.d/www.conf