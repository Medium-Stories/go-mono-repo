**Run all services**
```shell
docker-compose up
```

**User service**
* Runs on port 8001 (configurable)
* Main service for user accounts management
* Will publish an event that corresponds to an account action (created, deleted...)

**Listener service**
* Demo listener for user account events
* Listen for events and do some additional actions

**Product service**
* Runs on port 8002 (configurable)
* Main service to get products to be displayed
* Communicates to Discounts service through grpc to apply discounts on given products

**Discount service**
* Runs on port 8003 (configurable)
* Used to apply discounts on given products

**Gateway**
* Runs on port 8000 (configurable)
* Exposes http endpoints for services: User, Product