# Wallet service

```console
//Проверить баланс
curl -v -H 'Content-Type: application/json' -X POST -d '{"id":1}' http://localhost:8080/checkBalance

//Зачислить/списать средства
curl -v -H 'Content-Type: application/json' -X POST -d '{"id":1, "delta":10}' http://localhost:8080/changeBalance

//Перевести средства
curl -v -H 'Content-Type: application/json' -X POST -d '{"IdSender":1, "IdReceiver":3, "delta":10}' http://localhost:8080/transfer
```