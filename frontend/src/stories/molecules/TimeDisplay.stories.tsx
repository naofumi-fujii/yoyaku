import type { Meta, StoryObj } from '@storybook/react';
import TimeDisplay from '../../components/molecules/TimeDisplay';

const meta: Meta<typeof TimeDisplay> = {
  component: TimeDisplay,
  title: 'Molecules/TimeDisplay',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof TimeDisplay>;

export const Default: Story = {
  args: {
    start: new Date('2024-03-29T10:00:00'),
    end: new Date('2024-03-29T11:00:00'),
  },
};

export const Afternoon: Story = {
  args: {
    start: new Date('2024-03-29T14:30:00'),
    end: new Date('2024-03-29T15:45:00'),
  },
};

export const WithCustomClass: Story = {
  args: {
    start: new Date('2024-03-29T10:00:00'),
    end: new Date('2024-03-29T11:00:00'),
    className: 'text-blue-600 font-bold',
  },
};