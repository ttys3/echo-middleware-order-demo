# echo middleware apply and execution order

## demo log

```
2023/07/26 13:32:06.648755 main.go:87: Middleware3 wrap
2023/07/26 13:32:06.648766 main.go:71: Middleware2 wrap
2023/07/26 13:32:06.648770 main.go:51: Middleware1 wrap
2023/07/26 13:32:06.648773 main.go:54: ::Middleware1 begin exec
2023/07/26 13:32:06.749109 main.go:74: ::Middleware2 begin exec
2023/07/26 13:32:06.949474 main.go:76: Middleware2: get header from middleware1: ::middleware1: 2023-07-26 21:32:06.749024698 +0800 HKT m=+9.272094115
2023/07/26 13:32:06.949495 main.go:89: ::Middleware3 begin exec
2023/07/26 13:32:07.350006 main.go:117: RouterSpecificMiddleware2 wrap
2023/07/26 13:32:07.350023 main.go:102: RouterSpecificMiddleware1 wrap
2023/07/26 13:32:07.350027 main.go:104: ::RouterSpecificMiddleware1 begin exec
2023/07/26 13:32:07.650383 main.go:119: ::RouterSpecificMiddleware2 begin exec
2023/07/26 13:32:08.400908 main.go:31: /foo ::handler exec
2023/07/26 13:32:08.400941 main.go:123: RouterSpecificMiddleware2 took 750
2023/07/26 13:32:08.400946 main.go:108: RouterSpecificMiddleware1 took 1050
2023/07/26 13:32:08.400949 main.go:93: Middleware3 took 1451
2023/07/26 13:32:08.400951 main.go:78: Middleware2 took 1651
2023/07/26 13:32:08.400954 main.go:61: Middleware1 took 1752
```