FROM scratch

ENV PORT 32000
EXPOSE $PORT

COPY hello-eks /
