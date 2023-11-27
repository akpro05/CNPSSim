
# This docker file is for development/local environment
FROM debian:stable-slim

WORKDIR /

COPY CNPSSim /
RUN mkdir /conf
COPY conf /conf

RUN mkdir /views
COPY views /views


RUN mkdir /static
COPY static /static

RUN chmod +x /CNPSSim


ENTRYPOINT ["/CNPSSim"]
