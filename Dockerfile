FROM scratch
COPY ./deploygif-be /deploygif-be
ENTRYPOINT ["/deploygif-be"]
EXPOSE 8000
