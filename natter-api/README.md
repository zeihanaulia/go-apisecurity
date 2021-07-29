# Vulnerable API

Kode yang kita buat sebelumnya masuk kedalam [top 10 owasp](https://owasp.org/www-project-api-security/).
Artinya ada sercurity vulnerability yang serius.

Injection Attact bisa terjadi pada setiap API yang menerima inputan dari user dan mengesekusi query database.
## Reproduce

Hit api dengan data

```
curl --location --request POST 'localhost:4567/spaces' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test",
    "owner": "'\''); DELETE FROM spaces;  ## "
}'
```

Data diatas akan berdampak mengerikan. Data dalam table space akan terhapus semua.

### Penjelasan

Dengan data seperti itu dan input tidak divalidasi maka akan membentuk query seperti ini

```sql
INSERT INTO spaces(name, owner) VALUES('test',''); DELETE FROM spaces;  ## ')
```

Artinya, kita seharusnya hanya menjalankan query insert. Tapi karena tidak divalidasi inputannya.
Penyerang bisa menyisipi statement lain yang dapat merusak data kita.

## Cara mencegah

Cara mencegahnya ada 2, yaitu:

1. Melalui Code
2. Melalui SQL Permission

### Melalui Code

Ini adalah cara yang **DiRecomendasikan**, memperbaiki code dengan memvalidasi inputan.

Untuk di golang hanya perlu mengubah cara execute querynya.

Yang sebelumnya seperti ini

```go
db.Exec("INSERT INTO spaces(name, owner) VALUES('" + space.Name + "','" + space.Owner + "')")
```

menjadi

```go
db.Exec("INSERT INTO spaces(name, owner) VALUES(?,?)", space.Name, space.Name)
```

Dengan menggunakan argument, inputan yang akan diexecute akan dibersihkan dulu.
Karakter karakter yang merusak query akan diubah formatnya, sehinga menjadi string biasa.

Selain itu, agar lebih akurat sebaiknya kita memvalidasi setiap inputan yang kita terima.
Misalnya seperti memeriksa jumlah karakter dan karakter yang boleh dimasukan kedalam database. 

## Note

untuk me-reproduce issue koneksi ke db harus disetmultistatement=true

```bash
root:root@tcp(127.0.0.1:3306)/natter?multiStatements=true
```

