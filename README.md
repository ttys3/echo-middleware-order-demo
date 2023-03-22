# echo middleware apply and execution order

## demo log

```
2023/03/22 14:55:17.796945 main.go:14: http server starting...
2023/03/22 14:55:17.796991 main.go:18: use Middleware1
2023/03/22 14:55:17.796995 main.go:20: use Middleware2
2023/03/22 14:55:17.796997 main.go:22: use Middleware3
⇨ http server started on [::]:1323
2023/03/22 14:55:18.587755 main.go:80: Middleware3 wrap
2023/03/22 14:55:18.587766 main.go:64: Middleware2 wrap
2023/03/22 14:55:18.587769 main.go:46: Middleware1 wrap
2023/03/22 14:55:18.587778 main.go:49: ::Middleware1 begin exec
2023/03/22 14:55:18.688078 main.go:67: ::Middleware2 begin exec
2023/03/22 14:55:18.888305 main.go:69: Middleware2: get header from middleware1: ::middleware1: 2023-03-22 22:55:18.687977316 +0800 HKT m=+0.891412976
2023/03/22 14:55:18.888319 main.go:82: ::Middleware3 begin exec
2023/03/22 14:55:19.339010 main.go:26: ::handler exec
2023/03/22 14:55:19.339054 main.go:86: Middleware3 took 450
2023/03/22 14:55:19.339059 main.go:71: Middleware2 took 650
2023/03/22 14:55:19.339063 main.go:54: Middleware1 took 751
```