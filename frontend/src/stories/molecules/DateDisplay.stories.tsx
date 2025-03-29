import type { Meta, StoryObj } from '@storybook/react';
import DateDisplay from '../../components/molecules/DateDisplay';

const meta: Meta<typeof DateDisplay> = {
  component: DateDisplay,
  title: 'Molecules/DateDisplay',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof DateDisplay>;

export const Default: Story = {
  args: {
    date: new Date('2024-03-29'),
  },
};

export const Heading: Story = {
  args: {
    date: new Date('2024-03-29'),
    variant: 'h3',
  },
};

export const WithCustomClass: Story = {
  args: {
    date: new Date('2024-03-29'),
    className: 'text-green-600 italic',
  },
};