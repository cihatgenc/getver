FROM alpine

ENV TMPDIR /tmp

RUN mkdir /opt && \
	apk add --update ca-certificates && \
	update-ca-certificates

COPY getver /opt

EXPOSE 9080

CMD ["/opt/getver"]