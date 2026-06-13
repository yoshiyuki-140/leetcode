#!/usr/bin/bash

# 必要なコマンドのチェック
for cmd in curl jq; do
    if ! command -v $cmd &> /dev/null; then
        echo "Error: $cmd is not installed."
        exit 1
    fi
done

# LeetCodeの全問題のメタデータを一時的に取得して、番号からSlugを逆引きする関数
get_slug_by_number() {
    local target_num=$1
    
    # 毎回全件取得すると重いので、キャッシュ（/tmp/lc_problems.json）を利用
    local cache_file="/tmp/lc_problems.json"
    
    # キャッシュが古い（1日以上経過）か存在しない場合は取得
    if [ ! -f "$cache_file" ] || [ $(find "$cache_file" -mtime +1 2>/dev/null | wc -l) -gt 0 ]; then
        echo "🔄 LeetCodeの問題リストを同期中...（初回のみ少し時間がかかります）" >&2
        curl -s "https://leetcode.com/api/problems/all/" > "$cache_file"
    fi

    # 入力された番号に一致する title_slug を抽出
    # frontend_question_id は数値型か文字列型か怪しいので、両方に対応できるよう文字列として比較
    jq -r --arg num "$target_num" '.stat_status_pairs[] | select(.stat.frontend_question_id | tostring == $num) | .stat.question__title_slug' "$cache_file"
}

# メインルーープ
while true; do
    echo "=============================================="
    read -p "📝 作成したいLeetCodeの問題番号を入力してください (例: 1, 703) [qで終了]: " input_num
    
    # 終了判定
    if [[ "$input_num" == "q" || "$input_num" == "Q" ]]; then
        echo "バイバイ！👋"
        break
    fi

    # 入力が数値かチェック
    if [[ ! "$input_num" =~ ^[0-3]?[0-9]{1,4}$ ]]; then
        echo "❌ 有効な数値を入力してください。"
        continue
    fi

    echo "🔍 問題情報を検索中..."
    slug=$(get_slug_by_number "$input_num")

    if [ -z "$slug" ] || [ "$slug" = "null" ]; then
        echo "❌ 問題番号 $input_num に一致する問題が見つかりませんでした。"
        continue
    fi

    echo "🎯 対象のSlugを発見: $slug"

    # GraphQL APIで問題の詳細（タイトル、問題本文、Code Snippets）を取得
    json_payload=$(jq -n --arg slug "$slug" '{
        query: "query questionData($titleSlug: String!) { question(titleSlug: $titleSlug) { questionFrontendId title titleSlug content codeSnippets { lang langSlug code } } }",
        variables: { titleSlug: $slug }
    }')

    response=$(curl -s -X POST https://leetcode.com/graphql \
        -H "Content-Type: application/json" \
        -d "$json_payload")

    # データの抽出
    title_id=$(echo "$response" | jq -r '.data.question.questionFrontendId')
    title_slug=$(echo "$response" | jq -r '.data.question.titleSlug')
    content_html=$(echo "$response" | jq -r '.data.question.content')

    if [ "$content_html" = "null" ] || [ -z "$content_html" ]; then
        echo "❌ 問題文の取得に失敗しました（Premium限定問題の可能性があります）。"
        continue
    fi

    # ディレクトリ名の決定 (例: s703-kth-largest-element-in-a-stream)
    dir_name="s${title_id}-${title_slug}"

    if [ -d "$dir_name" ]; then
        echo "⚠️ すでにディレクトリ '$dir_name' が存在します。上書きを避けるためスキップします。"
        continue
    fi

    # ディレクトリ作成
    mkdir -p "$dir_name"

    # HTML形式の問題文を簡易的にMarkdown（プレーンテキスト）に変換
    # ※ 本格的な変換をしたい場合は pandoc などを使うと綺麗になりますが、簡易的にタグを消去・置換します
    problem_md=$(echo "$content_html" | \
        sed -e 's/<p>/\n/g' -e 's/<\/p>//g' \
        -e 's/<code>/`/g' -e 's/<\/code>/`/g' \
        -e 's/<strong[^>]*>/**/g' -e 's/<\/strong>/**/g' \
        -e 's/<pre[^>]*>/```\n/g' -e 's/<\/pre>/\n```/g' \
        -e 's/<ul>/\n/g' -e 's/<\/ul>//g' \
        -e 's/<li>/* /g' -e 's/<\/li>//g' \
        -e 's/&nbsp;/ /g' -e 's/&lt;/</g' -e 's/&gt;/>/g' -e 's/&amp;/\&/g' -e 's/&quot;/"/g')

    # README.md の書き出し
    cat << EOF > "${dir_name}/README.md"
# [${title_id}] ${title_slug}

URL: https://leetcode.com/problems/${title_slug}/

## 問題文
${problem_md}
EOF

    echo "✨ テンプレートを作成しました！"
    echo "   📁 ディレクトリ: ${dir_name}/"
    echo "   📄 問題文: ${dir_name}/README.md"
    echo ""
done