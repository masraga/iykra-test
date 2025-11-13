# IYKRA TEST

Dokumentasi instalasi aplikasi pegawai

### Instalasi

1. Pastikan [docker](https://www.docker.com/) sudah terinstall diperangkat kamu
2. Pastikan service docker dan kubernetes sudah aktif
3. Buka project dan arahkan terminal/cmd ke root directory project
4. Build dockerfile

```sh
docker build -t iykra-golang:latest .
```

5. Deploy aplikasi ke kubernetes

```sh
kubectl apply -f deployment.yaml
```

6. Ekspos aplikasi golang

```sh
kubectl port-forward svc/golang-app-service 3000:3000
```

7. Ekspos adminer database administrator

```sh
kubectl port-forward svc/adminer-service 8080:8080
```

8. Buka adminer database administrator di browser

```sh
http://localhost:8080
```

9. Login ke sistem adminer
   system=PostgreSQL
   user=root
   pass=root
   service=db
   database=iykra-test
10. Import _Database.sql_ yang ada didalam directory project ke adminer database administrator

### Endpoint

Tambah karyawan

```sh
curl --location 'localhost:3000/employees' \
--header 'Content-Type: application/json' \
--data '{
    "name": "RAGA MULIA KUSUMA",
    "position": "Backend Engineer",
    "salary": 9000000
}'
```

Tampilkan karyawan

```sh
curl --location 'localhost:3000/employees'
```

Tampilkan karyawan berdasarkan ID

```sh
curl --location 'localhost:3000/employees/1'
```

Edit karyawan

```sh
curl --location --request PUT 'localhost:3000/employees/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "RAGA MULIA KUSUMA UPDATE",
    "position": "Backend Engineer",
    "salary": 9000000
}'
```

hapus karyawan

```sh
curl --location --request DELETE 'localhost:3000/employees/1' \
--data ''
```
