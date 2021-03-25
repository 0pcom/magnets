#clean:
#	rm -rf cockroach-data certs private

#secure: clean
#	mkdir -p certs private && \
	cockroach cert create-ca --certs-dir=certs --ca-key=private/ca.key && \
	#cockroach cert create-node mainframe --certs-dir=certs --ca-key=private/ca.key && \
	cockroach cert create-node localhost --certs-dir=certs --ca-key=private/ca.key && \
	cockroach cert create-client root --certs-dir=certs --ca-key=private/ca.key && \
	cockroach cert create-client madmin --certs-dir=certs --ca-key=private/ca.key && \
	cockroach start-single-node --certs-dir=certs

rerun:
	cockroach start-single-node --certs-dir=certs


#insecure: clean
#	cockroach start-single-node --insecure

db-secure:
	cockroach sql --certs-dir=certs -e  'CREATE USER IF NOT EXISTS madmin WITH PASSWORD "g00dyear"; CREATE DATABASE IF NOT EXISTS product; GRANT ALL ON DATABASE product TO madmin;'

#db-insecure:
#	cockroach sql --insecure -e 'CREATE USER IF NOT EXISTS madmin WITH PASSWORD "g00dyear"; CREATE DATABASE IF NOT EXISTS product; GRANT ALL ON DATABASE product TO madmin;'
