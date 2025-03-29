import type { Meta, StoryObj } from '@storybook/react';
import ReservationItem from '../../components/molecules/ReservationItem';

const meta: Meta<typeof ReservationItem> = {
  component: ReservationItem,
  title: 'Molecules/ReservationItem',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    (Story) => (
      <div style={{ maxWidth: '400px', width: '100%' }}>
        <ul className="divide-y divide-gray-200 border border-gray-200 rounded-md">
          <Story />
        </ul>
      </div>
    ),
  ],
};

export default meta;
type Story = StoryObj<typeof ReservationItem>;

export const Default: Story = {
  args: {
    reservation: {
      id: '1',
      startTime: '2024-03-29T10:00:00Z',
      endTime: '2024-03-29T11:00:00Z',
      createdAt: '2024-03-28T08:00:00Z',
      updatedAt: '2024-03-28T08:00:00Z',
    },
    onDelete: (id) => console.log(`Delete reservation with id: ${id}`),
  },
};

export const Afternoon: Story = {
  args: {
    reservation: {
      id: '2',
      startTime: '2024-03-29T14:30:00Z',
      endTime: '2024-03-29T16:00:00Z',
      createdAt: '2024-03-28T09:00:00Z',
      updatedAt: '2024-03-28T09:00:00Z',
    },
    onDelete: (id) => console.log(`Delete reservation with id: ${id}`),
  },
};