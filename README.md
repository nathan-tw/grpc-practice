# grpc-practice

3/20[BESG](https://github.com/BESG-TW)的gRPC分享，紀錄學習gRPC遇到的問題及使用方式。

## What is gRPC

gRPC是一個基於http/2的遠端呼叫系統，使用protocol buffer作為界面描述語言，常用於服務之間的資料交換，
client和server都會由同一個.proto文件編譯出相對應語言的檔案檔案內會包含request和response的方法，client和server只要import後就可以調用了！整體流程如下圖

<img src=images/grpc.png width=800px/>

至於如何將proto檔編譯成對應的程式語言呢？我們會用到[protoc](http://google.github.io/proto-lens/installing-protoc.html)這個cli，並下載產生相對應語言的binary plugins，
編譯後產生兩個檔案，一個負責grpc的服務，一個則是protocol buffer的應用，如下圖

<img src=images/protocol_buffer.png width=400px/>

## Installation

我是按照[官網](https://grpc.io/)的步驟安裝：

#### 安裝 [protoc cli](http://google.github.io/proto-lens/installing-protoc.html)
因為macOS很簡單，`brew install`就結束了，所以我以linux為例：

```bash
PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

接著下載相對應語言的plugin binary，記得將其放在$PATH中，下載方式可參考[官網支援的語言](以go語言為例)，以go語言為例：

```bash
$ export GO111MODULE=on  # Enable module mode
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

安裝步驟完成～

## Usage

(todo)

