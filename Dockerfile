FROM golang:1.21.1-alpine as build
RUN mkdir -p /go/src/github.com/johnny-yeng/drone-sonar-plugin
WORKDIR /go/src/github.com/johnny-yeng/drone-sonar-plugin 
COPY *.go ./
COPY vendor ./vendor/
RUN go mod init github.com/johnny-yeng/drone-sonarqube
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonar

FROM openjdk:17-jdk-slim

ARG SONAR_VERSION=5.0.1.3006-linux
ARG SONAR_SCANNER_CLI=sonar-scanner-cli-${SONAR_VERSION}
ARG SONAR_SCANNER=sonar-scanner-${SONAR_VERSION}

RUN apt-get update \
    && apt-get install -y nodejs curl zip\
    && apt-get clean

COPY --from=build /go/src/github.com/johnny-yeng/drone-sonar-plugin/drone-sonar /bin/
WORKDIR /bin

RUN curl https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/${SONAR_SCANNER_CLI}.zip -so /bin/${SONAR_SCANNER_CLI}.zip
RUN unzip ${SONAR_SCANNER_CLI}.zip \
    && rm ${SONAR_SCANNER_CLI}.zip 

ENV PATH $PATH:/bin/${SONAR_SCANNER}/bin

RUN sed -i "s|#sonar.host.url|sonar.host.url|g" /bin/${SONAR_SCANNER}/conf/sonar-scanner.properties

ENTRYPOINT /bin/drone-sonar
