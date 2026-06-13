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

## 問題の追加方法

```bash
$ cd solutions/
$ make sol 
./create_solution_template.sh
==============================================
📝 作成したいLeetCodeの問題番号を入力してください (例: 1, 703) [qで終了]: 1234
🔍 問題情報を検索中...
🎯 対象のSlugを発見: replace-the-substring-for-balanced-string
✨ テンプレートを作成しました！
   📁 ディレクトリ: s1234-replace-the-substring-for-balanced-string/
   📄 問題文: s1234-replace-the-substring-for-balanced-string/README.md

==============================================
📝 作成したいLeetCodeの問題番号を入力してください (例: 1, 703) [qで終了]: q
バイバイ！👋
```

## 実行方法

各問題のディレクトリに移動して実行:

```bash
cd ./solutions/<problem-directory>
go run main.go
```
