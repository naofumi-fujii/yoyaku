import type { Meta, StoryObj } from '@storybook/react';
import Card from '../../components/atoms/Card';

const meta: Meta<typeof Card> = {
  component: Card,
  title: 'Atoms/Card',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Card>;

export const Default: Story = {
  args: {
    children: <div className="p-4">Card Content</div>,
  },
};

export const NoShadow: Story = {
  args: {
    hasShadow: false,
    children: <div className="p-4">Card without shadow</div>,
  },
};

export const NoPadding: Story = {
  args: {
    padding: 'none',
    children: <div className="p-4">Card with no padding (content has its own padding)</div>,
  },
};

export const SmallPadding: Story = {
  args: {
    padding: 'small',
    children: <div>Card with small padding</div>,
  },
};

export const MediumPadding: Story = {
  args: {
    padding: 'medium',
    children: <div>Card with medium padding</div>,
  },
};

export const LargePadding: Story = {
  args: {
    padding: 'large',
    children: <div>Card with large padding</div>,
  },
};