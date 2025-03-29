import type { Meta, StoryObj } from '@storybook/react';
import Spinner from '../../components/atoms/Spinner';

const meta: Meta<typeof Spinner> = {
  component: Spinner,
  title: 'Atoms/Spinner',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Spinner>;

export const Small: Story = {
  args: {
    size: 'small',
  },
};

export const Medium: Story = {
  args: {
    size: 'medium',
  },
};

export const Large: Story = {
  args: {
    size: 'large',
  },
};

export const CustomColor: Story = {
  args: {
    size: 'medium',
    color: 'border-red-500',
  },
};