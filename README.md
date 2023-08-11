# Rate Listener

Service for receiving updated rates for three instruments: ETH-BTC, BTC-USD, BTC-EUR

### To launch the service execute following commands:

To run docker-compose with MySQL
```
task start-db
```

If you don't have `migrate`
```
task install-migrate
```

Then run necessary migrations
```
task run-migrate
```

After that you can run the service with
```
go run cmd\app\main.go
```