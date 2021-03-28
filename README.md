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

#### Step1

編輯proto檔，syntax使用proto3(最多人使用)，message則是server或client拿到的資料格式，可以想像成有型態的json， service則是像一個interface，定義了這個服務應該有的溝通方式卻不實現

```proto
syntax = "proto3";

package calculator;
option go_package = "proto/calculator";

message CalculatorRequest {
  int64 a = 1;
  int64 b = 2;
}

message CalculatorResponse {
  int64 result = 1;
}

service CalculatorService {
  rpc Sum(CalculatorRequest) returns (CalculatorResponse) {};
}
```

#### Step2
編譯proto檔，假設我們現在要做的是編譯出一個以go語言開發的server端服務

```bash
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    {path/to/protofile}
```
如前所述，這會產生兩個檔案，以server端為例，我會有兩個動作需要做：

* 實做一個server type並組合(類似繼承)Unimplemented{Service}ServiceServer
* 實現該type有的function

以我的加法計算為例：

```go
type Server struct{
	calculatorPB.UnimplementedCalculatorServiceServer // import from file compiled from proto file
}

func (*Server) Sum(ctx context.Context, req *calculatorPB.CalculatorRequest) (*calculatorPB.CalculatorResponse, error) {
	fmt.Printf("Sum function is invoked with %v \n", req)

	a := req.GetA()
	b := req.GetB()

	res := &calculatorPB.CalculatorResponse{
		Result: a + b,
	}

	return res, nil

}
```

#### Step3

使用grpc server的sdk

```bash
$ go get
```

啟動server就完成嘍，可以使用bloomrpc

```go
func main() {
	fmt.Println("starting gRPC server")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	calculatorPB.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
```
