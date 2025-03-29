import type { Meta, StoryObj } from '@storybook/react';
import MainLayout from '../../components/templates/MainLayout';

const meta: Meta<typeof MainLayout> = {
  component: MainLayout,
  title: 'Templates/MainLayout',
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof MainLayout>;

export const Default: Story = {
  args: {
    title: '予約システム',
    children: (
      <div className="p-4 bg-gray-100 min-h-[300px] flex items-center justify-center">
        <p className="text-xl text-gray-700">ページコンテンツがここに表示されます</p>
      </div>
    ),
  },
};

export const WithCustomTitle: Story = {
  args: {
    title: 'カスタムタイトル',
    children: (
      <div className="p-4 bg-gray-100 min-h-[300px] flex items-center justify-center">
        <p className="text-xl text-gray-700">ページコンテンツがここに表示されます</p>
      </div>
    ),
  },
};