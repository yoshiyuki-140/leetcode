この問題は株取引の問題です。
pricesは単位時間ごとの株価を表していて、未来の株価がわかる前提の下利益が最大となる取引を1度だけ行う。
その最大の利益を返却する関数を実装する。

つまり、下の画像でいうと緑のところで買って、赤のところで売ればよい。

![](https://storage.googleapis.com/zenn-user-upload/7ecfffb840bd-20250903.png)

まず、変数lとrを各々0と1で初期化する。-> `l,r := 0,1`

![](https://storage.googleapis.com/zenn-user-upload/ea86974a0092-20250903.png)

`prices[r]`のほうが`prices[l]`よりも小さいことが図よりわかる。
この場合は、rの値をlに代入し、rは一つ右隣に移動する。
すなわち、

`l = r`
`r++`

![](https://storage.googleapis.com/zenn-user-upload/71328bc4aa0b-20250903.png)

