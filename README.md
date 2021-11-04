axs-auto-restake
====

Mình đang stake AXS trên https://stake.axieinfinity.com/ nhưng vì lười nên mình quyết định thử viết script để auto restake.

Sau khi code thử và chạy thì mình biết được https://api.roninchain.com/rpc chỉ là RPC readonly của Roninchain.

Vì bản tính tò mò và ham học hỏi nên mình quyết định download source code của [Ronin Wallet](https://chrome.google.com/webstore/detail/ronin-wallet/fnjhmkhhmkbjkkabndcnnogagogbneec) trên Chrome Store bằng [extension này](https://chrome.google.com/webstore/detail/chrome-extension-source-v/jifpbeccnghkjeaalbbjmodiffmgedin) nhưng source code đã được _[obfuscated](https://en.wikipedia.org/wiki/Obfuscation_(software))_ nên mình lại thử dùng [Burp Suite](https://portswigger.net/burp) để track xem wallet gọi đi đâu.

Cuối cùng mình đã tìm được RPC chính hiệu là https://proxy.roninchain.com/free-gas-rpc

Nhờ sự giúp sức tận tình của **[lugondev](https://github.com/lugondev)** đã dựng testnet trên bsc để test code.

Vì còn một số chỗ nữa thì mới chạy ngon được nên các việc cần làm tiếp theo ở bên dưới.

# TODO
- [ ] Check thời gian *thích hợp* để restake?
- [ ] Check restake đã success chưa?
- [ ] Scripts / docs deploy to GCP Cloudfunctions

# RUN

```bash
# cp .env.example .env
# thêm dòng
# MNEMONIC=là 12 chữ cái private lúc tạo ví Ronin
go run main.go
```