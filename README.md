# Arkham Server

Powered by custom [Core](https://github.com/Lazy-Parser/Collector.git) library.

---

### TODO
- [ ] Mexc exchange already added. Create process router / endpoints to manage processes
- [ ] Create NATS + some Publisher, ( [Exchange1, exchange2, ...] -> Publisher -> NATS -> [Consumer1, consumer2, ...])   

---

### Docs

**Swagger:**

```bash
swag init
```

You can open Swagger Documentation by running server:

```go
go run .
```

And open the url in your browser:

```bash
http://localhost:8080/swagger/index.html
```

---

### Login
1. Enter login on the site
2. The Bot will send you one-time password
3. Enter this password

---

### Future plans:

- [ ]  Metrics
- [ ]  Websockets
- [ ]  Integrate exchanges from [Core](https://github.com/Lazy-Parser/Collector.git) (50%)
- [ ]  And a lot more