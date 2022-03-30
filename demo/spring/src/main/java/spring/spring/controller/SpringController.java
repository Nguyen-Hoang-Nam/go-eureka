package spring.spring.controller;

import com.netflix.appinfo.InstanceInfo;
import com.netflix.discovery.EurekaClient;
import com.netflix.discovery.shared.Application;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Mono;
import spring.spring.request.SpringRequest;
import spring.spring.response.SpringResponse;

@Controller
@RestController
public class SpringController {

    @Autowired
    private EurekaClient eurekaClient;

    public String getBaseUrl(String framework) {
        Application application = eurekaClient.getApplication(framework);
        InstanceInfo instanceInfo = application.getInstances().get(0);

        String hostname = instanceInfo.getHostName();
        int port = instanceInfo.getPort();
        return "http://" + hostname + ":" + port;
    }

    @PostMapping
    public SpringResponse hi(@RequestBody SpringRequest framework) {
        String message = "Spring hi, " + framework.getName();

        return new SpringResponse(message);
    }

    @GetMapping(path = "/hello/{framework}")
    public Mono<SpringResponse> hello(@PathVariable String framework) {
        WebClient webClient = WebClient.create();

        SpringRequest springRequest = new SpringRequest("Spring");

        String serviceUrl = getBaseUrl(framework);

        return webClient
            .post()
            .uri(serviceUrl)
            .body(Mono.just(springRequest), SpringRequest.class)
            .retrieve()
            .bodyToMono(SpringResponse.class);
    }
}
