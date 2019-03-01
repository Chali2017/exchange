package com.ten.client;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestController {
    @Autowired RMIService hiService;

    @RequestMapping(value = "/hi")
    public String hi(@RequestParam String name){
        return hiService.hiService(name);
    }

    @RequestMapping(value = "/hiGo")
    public String hiGo(@RequestParam String name){
        return hiService.hiGoService(name);
    }

    @RequestMapping(value="/callbygo")
    public String callByGo(){
        return "hi go,I am java!";
    }
}
