# NATS JetStream Pub-Sub Mechanism
This is a basic pub-sub project using NATS JetStream. There is a publisher who publishes a "Hello World" message to a subject called "jet" and a subscriber who receives the message by subscribing to "jet" as soon as it is published. The main difference between core NATS server and JetStream is that core NATS server does not persist the messages published by publisher systems, and NATS JetStream persists the messages into a persistent store and those messages can be replayed from the persistent store, with ACK, rate matching/limiting, etc.

## Docker Installation of NATS
[Installing via Docker](https://docs.nats.io/running-a-nats-service/introduction/installation#installing-via-docker)

## Steps to Run

### Step 1: Run NATS JetStream server
```sudo docker run -p 4222:4222 -ti nats:latest -js```

### Step 2: Run nats-box for nats cli in a different terminal
[nats box](https://github.com/nats-io/nats-box)  
```sudo docker run --rm -it synadia/nats-box:latest```

### Step 3: Create a stream
Enter nats stream, add <Stream name> (in this example we will name the stream "my_stream"), then enter "jet" as the subject name and hit Enter to use the defaults for all the other stream attributes:
```nats stream add my_stream --server [Your IP Address]```

### Step 4: Run publisher
```go run publisher/main.go```

### Step 5: Run consumer
```go run consumer/main.go```

## Resources
[NATS - Go Client](https://github.com/nats-io/nats.go)  
[JetStream Walkthrough](https://docs.nats.io/nats-concepts/jetstream/js_walkthrough)  
[Building Distributed Event Streaming Systems In Go With NATS JetStream](https://shijuvar.medium.com/building-distributed-event-streaming-systems-in-go-with-nats-jetstream-3938e6dc7a13)