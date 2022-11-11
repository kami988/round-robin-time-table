FROM golang:1.19-buster

LABEL Name="maker-go"
LABEL Version="1.0.0"

ENV GITHUB_USER_NAME=GITHUB-user-name
ENV GITHUB_TOKEN=GITHUB-token
ENV GITHUB_EMAIL=GITHUB-email

WORKDIR /
COPY ./init.sh init.sh

RUN mkdir -p /repo
VOLUME /repo
WORKDIR /repo

CMD ["bash", "/init.sh"]
