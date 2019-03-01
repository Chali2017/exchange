package com.ten.client;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;


@Service
public class RMIService {
    /*@Autowired
    RestTemplate restTemplate;*/

    @Qualifier("restTemplate")
    @Autowired
    private RestTemplate rest;

    public String hiService(String name) {
        return rest.getForObject("http://DEMO/hi?name="+name,String.class);
    }

    public String hiGoService(String name) {
        return rest.getForObject("http://127.0.0.1:8088/version",String.class);
    }
}
