# d2l (discord2line)
Discord のメッセージを LINE に転送する小さなボットです。

## Requirements
- Go 1.24+

## 事前準備

### Discord 側
- **Bot の作成**: Discord Developer Portal でアプリ/ボットを作成し、Bot Token を取得します。
- **BOT_ID の確認**: 反応させたいボットの **ユーザーID**（Discord 上で bot を右クリック → ID をコピー）を取得します。
- **Gateway Intents**:
  - **Message Content Intent** を有効化してください（メッセージ本文を読むため）。
  - サーバーに招待する際は `bot` スコープで招待します。

### LINE 側
- Messaging API のチャネル（公式アカウント）を作成し、以下を取得します。
  - **Channel Secret**（`LINE_SECRET`）
  - **Channel Access Token (long-lived)**（`LINE_TOKEN`）
- 本実装は **Broadcast** で送信します（公式アカウントの友だち全員に送信されます）。運用前に送信範囲/権限/料金体系を必ず確認してください。

## 設定（.env）
本プロジェクトは起動時に **カレントディレクトリの `.env`** を読み込みます（`go run .` をリポジトリ直下で実行してください）。

`.env` の例:
```dotenv
DISCORD_TOKEN=xxxxxxxx
BOT_ID=123456789012345678
LINE_SECRET=xxxxxxxx
LINE_TOKEN=xxxxxxxx
```

- **DISCORD_TOKEN**: Discord Bot Token
- **BOT_ID**: ボットのユーザーID（メンション判定に利用）
- **LINE_SECRET**: LINE Channel Secret
- **LINE_TOKEN**: LINE Channel Access Token (long-lived)

## 使い方
### ローカル実行
```bash
go test ./...
go run .
```

### 動作仕様
- Discord で **ボットへのメンションが1つだけ**含まれるメッセージに反応します（メンションが0個/2個以上なら無視）。
- メッセージ本文から `<...>` 形式の部分（メンション等）を除去して LINE に送信します。
- LINE 側は `BroadcastMessage` を使って送信します。

## トラブルシュート
- **本文が転送されない / 空になる**: Discord 側の **Message Content Intent** が無効な可能性があります。
- **LINE に届かない**: `LINE_SECRET`/`LINE_TOKEN` が正しいか、Broadcast の送信条件（権限/送信範囲）を満たしているか確認してください。
