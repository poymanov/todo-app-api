FROM golang:1.23.5

ARG GROUP=1000
ARG USER=1000

ENV TZ=Europe/Moscow

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN groupadd -g $GROUP user
RUN useradd -u $USER -ms /bin/bash -g user user

WORKDIR /app