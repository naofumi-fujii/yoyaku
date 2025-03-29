import type { Meta, StoryObj } from '@storybook/react';
import SplitLayout from '../../components/templates/SplitLayout';

const meta: Meta<typeof SplitLayout> = {
  component: SplitLayout,
  title: 'Templates/SplitLayout',
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
  decorators: [
    (Story) => (
      <div className="p-6">
        <Story />
      </div>
    ),
  ],
};

export default meta;
type Story = StoryObj<typeof SplitLayout>;

export const Default: Story = {
  args: {
    leftTitle: '左カラム',
    rightTitle: '右カラム',
    leftContent: (
      <div className="p-4 bg-blue-100 min-h-[300px] rounded-lg flex items-center justify-center">
        <p className="text-xl text-blue-700">左カラムのコンテンツ</p>
      </div>
    ),
    rightContent: (
      <div className="p-4 bg-green-100 min-h-[300px] rounded-lg flex items-center justify-center">
        <p className="text-xl text-green-700">右カラムのコンテンツ</p>
      </div>
    ),
  },
};

export const WithoutTitles: Story = {
  args: {
    leftContent: (
      <div className="p-4 bg-blue-100 min-h-[300px] rounded-lg flex items-center justify-center">
        <p className="text-xl text-blue-700">左カラムのコンテンツ</p>
      </div>
    ),
    rightContent: (
      <div className="p-4 bg-green-100 min-h-[300px] rounded-lg flex items-center justify-center">
        <p className="text-xl text-green-700">右カラムのコンテンツ</p>
      </div>
    ),
  },
};

export const CustomWidths: Story = {
  args: {
    leftTitle: '広い左カラム',
    rightTitle: '狭い右カラム',
    leftWidth: 'w-full lg:w-3/4',
    rightWidth: 'w-full lg:w-1/4',
    leftContent: (
      <div className="p-4 bg-blue-100 min-h-[300px] rounded-lg flex items-center justify-center">
        <p className="text-xl text-blue-700">左カラムのコンテンツ</p>
      </div>
    ),
    rightContent: (
      <div className="p-4 bg-green-100 min-h-[300px] rounded-lg flex items-center justify-center">
        <p className="text-xl text-green-700">右カラムのコンテンツ</p>
      </div>
    ),
  },
};