clean0:
	rm -rf cockroach-data certs private

clean1:
	rm -rf cockroach-data certs1 private

certs: clean0 clean1 certs0 certs1
	echo

certs0: #clean0
	mkdir -p certs private && \
	cockroach cert create-ca --certs-dir=certs --ca-key=private/ca.key && \
	cockroach cert create-node --overwrite --certs-dir=certs --ca-key=private/ca.key localhost mainframe 127.0.0.1 192.168.2.130 && \
	cockroach cert create-client root --certs-dir=certs --ca-key=private/ca.key && \
	cockroach cert create-client madmin --certs-dir=certs --ca-key=private/ca.key

certs1: #clean1
	cp -r certs certs1
	cockroach cert create-node --overwrite --certs-dir=certs1 --ca-key=private/ca.key localhost magnetosphere 127.0.0.1 192.168.2.118 magnetosphere.net

	#cockroach start-single-node --certs-dir=certs
	#--listen-addr=192.168.2.130 --advertise-addr=192.168.2.130
start0:
	cockroach start \
	--certs-dir=certs \
	--advertise-addr=192.168.2.130 \
	--join=192.168.2.118 \
	--cache=.25 \
	--max-sql-memory=.25 \
#	--background

start1:
	cockroach start \
	--certs-dir=certs \
	--advertise-addr=192.168.2.118 \
	--join=192.168.2.130 \
	--cache=.25 \
	--max-sql-memory=.25 \

init:
	cockroach init --certs-dir=certs --host=192.168.2.118,192.168.2.130

single-node:
	cockroach start-single-node --certs-dir=certs

#insecure: clean
#	cockroach start-single-node --insecure

db-secure:
	cockroach sql --certs-dir=certs -e  'CREATE USER IF NOT EXISTS madmin WITH PASSWORD "g00dyear"; CREATE DATABASE IF NOT EXISTS product; GRANT ALL ON DATABASE product TO madmin;'

#db-insecure:
#	cockroach sql --insecure -e 'CREATE USER IF NOT EXISTS madmin WITH PASSWORD "g00dyear"; CREATE DATABASE IF NOT EXISTS product; GRANT ALL ON DATABASE product TO madmin;'
