FROM golang:1.21.5-bullseye

WORKDIR /app
COPY . .
COPY start-web.sh /start-web.sh
RUN chmod +x /start-web.sh


##
#RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
##ARG cert_location=/usr/local/share/ca-certificates
#### Get certificate from "github.com"
##RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
### Get certificate from "proxy.golang.org"
##RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
### Get certificate from "golang.org"
##RUN openssl s_client -showcerts -connect golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/golang.crt
## Temporarily disable Git SSL verification (not recommended for production)
#RUN git config --global http.sslVerify false
#
#ENV GOPROXY=direct
#ENV GIT_SSL_NO_VERIFY=true
#ENV GONOSUMDB=*
#ENV GOPRIVATE=*
#RUN go mod download
##
#
#RUN go build -o main main.go

COPY main /app/main

CMD ["./main"]