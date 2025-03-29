# Storybook 導入手順

このプロジェクトには Storybook が導入されています。以下の手順で実行できます。

## Storybook の実行方法

```bash
npm run storybook
```

これにより、デフォルトでは http://localhost:6006 で Storybook が起動します。

## Storybook のビルド方法

```bash
npm run build-storybook
```

これにより、静的な Storybook サイトが生成されます。

## コンポーネントの Story ファイル

以下のコンポーネントに Story ファイルが追加されています：

- Calendar.tsx → Calendar.stories.tsx
- ReservationList.tsx → ReservationList.stories.tsx
- ReservationModal.tsx → ReservationModal.stories.tsx

各 Story ファイルには複数のバリエーション（ストーリー）があり、コンポーネントの異なる状態を確認できます。

## Story ファイルの追加方法

新しいコンポーネントを作成した場合は、以下の命名規則に従って Story ファイルを作成してください：

```tsx
import type { Meta, StoryObj } from '@storybook/react';
import YourComponent from './YourComponent';

const meta: Meta<typeof YourComponent> = {
  component: YourComponent,
  title: 'Components/YourComponent',
  parameters: {
    layout: 'centered', // or 'fullscreen' for full-width components
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof YourComponent>;

export const Default: Story = {
  args: {
    // コンポーネントのpropsをここに記述
  },
};

// 他のバリエーションを追加
export const AnotherVariant: Story = {
  args: {
    // 異なるpropsをここに記述
  },
};
```

## 参考リンク

- [Storybook for Next.js](https://storybook.js.org/docs/get-started/frameworks/nextjs)
- [Writing Stories](https://storybook.js.org/docs/writing-stories)
- [Component Story Format (CSF)](https://storybook.js.org/docs/api/csf)
