https://activemq.apache.org/kahadb

ajustar volume para activemq, unmask 000

This sample project delves into the implementation of persistence mechanisms along ActiveMQ.

## Overview
We deal with various booking types — flights, accommodations, and vehicle rentals. Our goal is to
post some messages and see how they are stored in a persistent tool until they are consumed. We will
explore two ways to persist messages.

The Go project will simply post two messages, one for plane booking and other for hotel booking. We
will not implement any consumer to leave the message stored in the persistence store. It is important
to notice that we need to tag the messages with the header 'persistent:true' when sending them to the
broker.

## KahaDB
The first and simplest is the default KahaDB, a file-based datastore optimized for efficiency. In
production scenarios, this would be most likely the approach used. It follows a straightforward
journal model in an append-only approach.

KahaDB files are binaries that cannot be humanly interpreted and need outside tools to troubleshoot
and understand.

This is the default approach, so a standard docker image starts the container with KahaDB capabilities.
To make things more clear, though, we will explicitly mount the activemq.xml config file to the container.
It can be found [here](../../../resources/channels/persistence/persistent-activemq-kaha-db.xml).

It still was not possible to mount a volume to make the storage testable. Needs extra studying.

## JDBC | MySQL
Another common approach is to use a JDBC driver (PostgreSQL, MySQL, etc.) that will register messages,
ACKs and locks in a traditional database that can be easily accessed. We will use MySQL in our example.

To make this the chosen persistence approach, we need to configure our activemq.xml config file to
contain the following lines:
```
    <bean id="mysql-ds" class="org.apache.commons.dbcp2.BasicDataSource" destroy-method="close">
        <property name="driverClassName" value="com.mysql.cj.jdbc.Driver"/>
        <property name="url" value="jdbc:mysql://mysql:3306/activemq_db?useSSL=false&amp;allowPublicKeyRetrieval=true"/>
        <property name="username" value="activemq"/>
        <property name="password" value="activemq"/>
    </bean>

    <broker>
        <persistenceAdapter>
            <jdbcPersistenceAdapter dataSource="#mysql-ds"/>
        </persistenceAdapter>
    </broker>
```
The full file can be found [here](../../../resources/channels/persistence/persistent-activemq-mysql.xml).
We also need to pass the MySQL .jar driver to inside the container. This can be done through a customized
Dockerfile or, as in our case, manually placing it there after downloading.

We can easily access the database through these credentials. It will also be necessary to start a mysql
container with the correct information, as shown in our docker-compose.yml file.

## How to Run
The Makefile in the root can spin the docker-compose with the ActiveMQ broker and the Go project.
Simply run `make channels/persistence` and the messages will be stored in the corresponding persistence
stores. No output will be displayed by the command.
