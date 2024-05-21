ZOOKEEPRER_VERSION ?= 3.7.0

run:
ifeq ("$(wildcard apache-zookeeper-$(ZOOKEEPRER_VERSION)-bin.tar.gz)", "")
	wget "https://archive.apache.org/dist/zookeeper/zookeeper-$(ZOOKEEPRER_VERSION)/apache-zookeeper-$(ZOOKEEPRER_VERSION)-bin.tar.gz"
endif
	tar -xvf apache-zookeeper-$(ZOOKEEPRER_VERSION)-bin.tar.gz
	mv apache-zookeeper-$(ZOOKEEPRER_VERSION)-bin zookeeper
	cp zookeeper/conf/zoo_sample.cfg zookeeper/conf/zoo.cfg
	./zookeeper/bin/zkServer.sh start

stop:
	./zookeeper/bin/zkServer.sh stop
