FROM golang:1.23.5

ARG GROUP=1000
ARG USER=1000

ENV TZ=Europe/Moscow

WORKDIR /app

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN groupadd -g $GROUP user
RUN useradd -u $USER -ms /bin/bash -g user user

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

USER user

CMD ["air"]