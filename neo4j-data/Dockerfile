# This is a Dockerfile for the neo4j-data image
FROM neo4j:5.16.0

# load the neo4j.conf file into the container
COPY ./neo4j.conf /var/lib/neo4j/conf/neo4j.conf

# create the data and logs directories
# and set the permissions
RUN chmod 770 -R /data

# remove in production -- 'exec' instead
RUN chmod 770 -R /logs 
