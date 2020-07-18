FROM php:7.3

RUN pecl install apcu-5.1.18 && docker-php-ext-enable apcu

RUN mkdir /app
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
ADD ./app/package.json ./app/yarn.lock /app/
COPY ./app/ /app/

RUN composer install
ENTRYPOINT ["php", "-S", "0.0.0.0:8081", "index.php"]

EXPOSE 8081
