### 課題
- 課題1: doSomethingを非同期化してみる
- 課題2: doSomethingが実行し終わるのを待ってみる（ヒント: sync.WaitGroup）
- 課題3: doSomethingの並列数を管理してみる（ヒント: channelのcapacity飽和でlockされる仕組みを使うといいかと）
- 課題4: workerスタイルにしてみる（channelにqueueを投げて非同期で待機しているworkerが捌く）
### 追加課題
- 課題5: doSomething を cancel できるようにして、一定時間で timeout させる（ヒント: context.Context）
- 課題6: doSomething が panic する場合を想定して、プログラムが停止しないようにエラーハンドリングする
- 課題7:プログラムが外部から停止(SIGTERMなど)された場合に適切に終了処理を行う
