FROM alpine:3.10
ADD app app
CMD ./app -port=$PORT