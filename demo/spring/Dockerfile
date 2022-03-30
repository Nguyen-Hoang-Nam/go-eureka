FROM adoptopenjdk/openjdk11-openj9:jre-11.0.11_9_openj9-0.26.0-alpine
WORKDIR /
COPY wait-for wait-for
COPY target/spring-1.0.0.jar spring.jar
ENTRYPOINT ["./wait-for", "eureka:9000", "--", "java", "-jar", "spring.jar" ]
