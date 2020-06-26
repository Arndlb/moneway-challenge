FROM scylladb/scylla

COPY wrapper.sh /wrapper.sh
RUN chmod +x /wrapper.sh

EXPOSE 9042

ENTRYPOINT ["/wrapper.sh"]