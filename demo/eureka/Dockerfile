FROM adoptopenjdk/openjdk11-openj9:jre-11.0.11_9_openj9-0.26.0-alpine
WORKDIR /
COPY target/eureka-1.0.0.jar eureka.jar
ENTRYPOINT [ "java", "-jar", "eureka.jar" ]
