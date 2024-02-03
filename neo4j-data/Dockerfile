# This is a Dockerfile for the neo4j-data image
FROM neo4j:5.16.0

# load the neo4j.conf file into the container
COPY ./neo4j.conf /var/lib/neo4j/conf/neo4j.conf

RUN chmod 770 -R /data
RUN chown neo4j:neo4j -R /data
