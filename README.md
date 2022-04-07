# payment-app

## コマンド

```bash
# .protoからgo生成
protoc -Iproto --go_out=plugins=grpc:. proto/pay.proto 
```

## Pay.jpテスト用トークン生成

curlで行う
Headerに`X-Payjp-Direct-Token-Generate: true`を指定することでカード情報を直接指定して作成できます。
※ テストモードのみ有効

```
curl https://api.pay.jp/v1/tokens \
    -u sk_test_8f556f10bb4640d8144eb16b: \
    -H "X-Payjp-Direct-Token-Generate: true" \
    -d "card[number]=4242424242424242" \
    -d "card[cvc]=123" \
    -d "card[exp_month]=02" \
    -d "card[exp_year]=2024"

curl https://api.pay.jp/v1/tokens/tok_544ee93db794b683c5b7859c70e9 \
    -u sk_test_8f556f10bb4640d8144eb16b:
```

```
{
  "card": {
    "address_city": null,
    "address_line1": null,
    "address_line2": null,
    "address_state": null,
    "address_zip": null,
    "address_zip_check": "unchecked",
    "brand": "Visa",
    "country": null,
    "created": 1649325805,
    "customer": null,
    "cvc_check": "passed",
    "exp_month": 2,
    "exp_year": 2024,
    "fingerprint": "e1d8225886e3a7211127df751c86787f",
    "id": "car_ca5c002a5cd51c3e19ae008fa686",
    "last4": "4242",
    "livemode": false,
    "metadata": {},
    "name": null,
    "object": "card",
    "three_d_secure_status": null
  },
  "created": 1649325805,
  "id": "tok_544ee93db794b683c5b7859c70e9",
  "livemode": false,
  "object": "token",
  "used": false
}%            
```




grpc_cli call localhost:8000 PayManager.Charge 'id: 1, token: "tok_544ee93db794b683c5b7859c70e9": amount: 4000, name: "aiueo", description: "aiueo"'
grpc_cli call localhost:8000 PayManager.Charge 'id: 1, token: tok_544ee93db794b683c5b7859c70e9: amount: 4000, name: aiueo, description: aiueo'
grpc_cli call localhost:8000 PayManager.Charge 'id: 1,token: "tok_544ee93db794b683c5b7859c70e9", amount: 4000, name: "aiueo", description: "aiueo"'







