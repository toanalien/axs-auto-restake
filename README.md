axs-auto-restake
====

Mình đang stake AXS trên https://stake.axieinfinity.com/ nhưng vì lười nên mình quyết định thử viết script để auto restake.

Sau khi code thử và chạy thì mình biết được https://api.roninchain.com/rpc chỉ là RPC readonly của Roninchain.

Vì bản tính tò mò và ham học hỏi nên mình quyết định download source code của [Ronin Wallet](https://chrome.google.com/webstore/detail/ronin-wallet/fnjhmkhhmkbjkkabndcnnogagogbneec) trên Chrome Store bằng [extension này](https://chrome.google.com/webstore/detail/chrome-extension-source-v/jifpbeccnghkjeaalbbjmodiffmgedin) nhưng source code đã được _[obfuscated](https://en.wikipedia.org/wiki/Obfuscation_(software))_ nên mình lại thử dùng [Burp Suite](https://portswigger.net/burp) để track xem wallet gọi đi đâu.

Cuối cùng mình đã tìm được RPC chính hiệu là https://proxy.roninchain.com/free-gas-rpc

Nhờ sự giúp sức tận tình của **[lugondev](https://github.com/lugondev)** đã dựng testnet trên bsc để test code.

Vì còn một số chỗ nữa thì mới chạy ngon được nên các việc cần làm tiếp theo ở bên dưới.

## Update 6 Nov 2021
Vì cần phải tìm thời gian thích hợp để chạy restake nên mình cần biết cách tính time remain trên Web UI. Phần này mình xem trong code js thì ở đây trùng với request này

![code js](https://i.imgur.com/40Wu2JH.png)

![request](https://i.imgur.com/94bTK9H.png)

### Cách lấy ABI từ manifest

Mình search MethodID là `restakeRewards` ra đoạn code này, mình đoán là ABI

![](https://i.imgur.com/lKJ6woO.png)

Tiến hành copy nguyên array của `t.default` trong function `28711` và `36572` để ghép thành 1 array hoàn chỉnh.

Dùng VSCode để mông má lại:
- replace `!1` = `false`
- replace `!0` = `true`
![](https://i.imgur.com/1tQE0fZ.png)
- vì json lấy về phần `key` không nằm trong double quote nên chưa đúng định dạng cần replace lại như sau:
`\s.+?([a-zA-z].*):` thành `"$1":`

![before](https://i.imgur.com/6LyfZZO.png)

![after](https://i.imgur.com/y4bScD5.png)

![](https://i.imgur.com/mNg38hv.png)

Sau khi có file abi mình sẽ tiến hành convert sang go module bằng abigen ([cách cài đặt ở đây](https://goethereumbook.org/smart-contract-compile/))

```bash
abigen --abi=erc20.abi --pkg=token --out=erc20.go
```

Cuối cùng import file module `token` vào code và gọi lên contract.

# TODO
- [x] Check thời gian *thích hợp* để restake?
- [x] Cào abi từ manifest
- [x] Tự động restake
- [ ] Check restake đã success chưa?
- [ ] Scripts / docs deploy to GCP Cloudfunctions
- [ ] Thêm option dùng private key thay vì mnemonic

# RUN

```bash
# cp .env.example .env
# thêm dòng
# MNEMONIC=là 12 chữ cái private lúc tạo ví Ronin
go run main.go
```