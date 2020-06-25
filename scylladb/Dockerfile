FROM scylladb/scylla

COPY wrapper.sh /wrapper.sh
RUN chmod +x /wrapper.sh

ENTRYPOINT ["bash", "/wrapper.sh"]