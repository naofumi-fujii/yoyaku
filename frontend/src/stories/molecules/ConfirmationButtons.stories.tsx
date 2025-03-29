import type { Meta, StoryObj } from '@storybook/react';
import ConfirmationButtons from '../../components/molecules/ConfirmationButtons';

const meta: Meta<typeof ConfirmationButtons> = {
  component: ConfirmationButtons,
  title: 'Molecules/ConfirmationButtons',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof ConfirmationButtons>;

export const Default: Story = {
  args: {
    onCancel: () => console.log('Cancel clicked'),
    onConfirm: () => console.log('Confirm clicked'),
    cancelText: 'キャンセル',
    confirmText: '確認',
    confirmVariant: 'primary',
  },
};

export const CustomText: Story = {
  args: {
    onCancel: () => console.log('Cancel clicked'),
    onConfirm: () => console.log('Confirm clicked'),
    cancelText: '戻る',
    confirmText: '送信する',
    confirmVariant: 'primary',
  },
};

export const DangerAction: Story = {
  args: {
    onCancel: () => console.log('Cancel clicked'),
    onConfirm: () => console.log('Confirm clicked'),
    cancelText: '戻る',
    confirmText: '削除する',
    confirmVariant: 'danger',
  },
};