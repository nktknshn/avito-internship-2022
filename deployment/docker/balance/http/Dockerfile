FROM golang:1.24


RUN go install github.com/cespare/reflex@latest
COPY reflex.conf /

COPY start.sh /

ENTRYPOINT ["reflex", "-d", "none", "-c", "/reflex.conf"]