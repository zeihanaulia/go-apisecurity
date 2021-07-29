# Rate Limiter on code

Rate Limiter adalah pencegahan dari penyerangan DDos. 
Mekanismenya adalah membatasi request yang masuk. 
Idealnya rate limiter dipasang pada load balancer atau reverse proxy.
Pada contoh kali ini kita akan belajar membatasi dilevel code.

Contoh Kasus:

Membatasi hanya 1 request perdetik.

Di golang, kita bisa menggunakan library [tollbooth](https://github.com/didip/tollbooth).

Ubah code 
```
	router.HandleFunc("/spaces", createSpace).Methods("POST")
```

Menjadi

```go
    router.Handle(
		"/spaces",
		tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), createSpace),
	).Methods("POST")
```

library [tollbooth](https://github.com/didip/tollbooth) juga memiliki beberapa konfigurasi.
semisal jika ingin membatasi per IP address atau sepesifik authentication.