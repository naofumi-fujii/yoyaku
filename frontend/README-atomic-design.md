# Atomic Design パターン導入

このプロジェクトでは、Brad Frostが提唱したAtomic Designパターンを採用してコンポーネント構造を整理しています。

## 構造

コンポーネントは以下の5つのレベルに分類されています：

### 1. Atoms（原子）
最小単位の基本的なUIコンポーネント。これらは単独で存在し、他のコンポーネントに依存しません。

- `Button.tsx`: 様々なスタイルのボタン
- `Card.tsx`: コンテンツを包むカードコンポーネント
- `Spinner.tsx`: ローディングスピナー
- `Typography.tsx`: テキスト表示用コンポーネント

### 2. Molecules（分子）
複数のAtomsを組み合わせた、やや複雑なUIコンポーネント。

- `TimeDisplay.tsx`: 開始時間と終了時間を表示
- `DateDisplay.tsx`: 日付を表示
- `ReservationItem.tsx`: 予約アイテム1件分の表示
- `ConfirmationButtons.tsx`: 確認/キャンセルボタンのセット

### 3. Organisms（有機体）
複数のMoleculesやAtomsを組み合わせた、より複雑な機能を持つコンポーネント。

- `Calendar.tsx`: カレンダーコンポーネント
- `ReservationList.tsx`: 予約リスト
- `ReservationModal.tsx`: 予約確認モーダル

### 4. Templates（テンプレート）
ページのレイアウト構造を定義するコンポーネント。

- `MainLayout.tsx`: ヘッダー、フッターを含むメインレイアウト
- `SplitLayout.tsx`: 2カラムレイアウト

### 5. Pages（ページ）
実際のページを表すコンポーネント。

- `HomePage.tsx`: メインページコンポーネント

## Storybook統合

各コンポーネントレベルのStoryファイルは`src/stories/`ディレクトリに対応するフォルダ構造で格納されています：

```
src/
  ├── components/
  │   ├── atoms/
  │   ├── molecules/
  │   ├── organisms/
  │   ├── templates/
  │   └── pages/
  └── stories/
      ├── atoms/
      ├── molecules/
      ├── organisms/
      ├── templates/
      └── pages/
```

## 新しいコンポーネントを追加する場合

1. コンポーネントの役割と責任範囲を考慮し、適切なレベルを決定します
2. 対応するディレクトリに新しいコンポーネントを作成します
3. 同様のパターンに従い、必要なpropsとインターフェースを定義します
4. 対応するStoryファイルを作成します

## 利点

- **再利用性**: 小さなコンポーネントの再利用性が高まります
- **一貫性**: デザインの一貫性が保たれます
- **保守性**: コードの保守性が向上します
- **効率性**: 開発の効率が向上します
- **ドキュメント化**: Storybookとの統合によりコンポーネントカタログができます

## 参考リンク

- [Atomic Design by Brad Frost](https://atomicdesign.bradfrost.com/)
- [Storybook Documentation](https://storybook.js.org/docs/react/get-started/introduction)