# LeetCode Solutions

このリポジトリには、LeetCode問題の解答が含まれています。言語に特定の指定はありませんが、主にGo言語を使用して解答しています。

## 構造

```bash
.
├── README.md # this file
├── cz.json # git-cz-rsの設定ファイル
├── go.mod
└── solutions
    ├── Makefile
    ├── create_solution_template.sh # 問題番号からディレクトリを作ってくれるbash script
    ├── s1-two-sum
    ...
    └── s9-palindrome-number
```


## 実行方法

各問題のディレクトリに移動して実行:

```bash
cd ./solutions/<problem-directory>
go run main.go
```
