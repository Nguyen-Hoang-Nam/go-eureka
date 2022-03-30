const express = require("express");
const morgan = require("morgan");
const axios = require("axios");
const Eureka = require("eureka-js-client").Eureka;

const client = new Eureka({
    instance: {
        app: "express",
        hostName: "express",
        ipAddr: "0.0.0.0",
        vipAddress: "express",
        instanceId: "express:8003",
        port: {
            $: 8003,
            "@enabled": "true",
        },
        dataCenterInfo: {
            "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
            name: "MyOwn",
        },
    },
    eureka: {
        host: "localhost",
        port: 9000,
        servicePath: "/eureka/apps/",
    },
});

client.start();

const app = express();

app.use(morgan("dev"));
app.use(express.urlencoded({ extended: false }));
app.use(express.json());

app.post("/", (req, res) => {
    const { name } = req.body;

    res.json({ message: "Express hi, " + name });
});

app.get("/hello/:frameowork", (req, res) => {
    const framework = req.params.frameowork;
    const service = client.getInstancesByAppId(framework)[0];

    const hostname = service.hostName;
    const port = service.port["$"];

    const url = "http://" + hostname + ":" + port;
    axios
        .post(url, {
            name: "Express",
        })
        .then((response) => {
            res.send(response.data.message);
        })
        .catch((err) => {
            console.log(err);
        });
});

const server = app.listen(8003, () => {
    console.log("Node server is running..");
});
