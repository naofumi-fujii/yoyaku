import type { Meta, StoryObj } from '@storybook/react';
import ReservationList from '../../components/organisms/ReservationList';

const meta: Meta<typeof ReservationList> = {
  component: ReservationList,
  title: 'Organisms/ReservationList',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  decorators: [
    (Story) => (
      <div style={{ maxWidth: '400px', width: '100%' }}>
        <Story />
      </div>
    ),
  ],
};

export default meta;
type Story = StoryObj<typeof ReservationList>;

const mockReservations = [
  {
    id: '1',
    startTime: '2024-03-29T10:00:00Z',
    endTime: '2024-03-29T11:00:00Z',
    createdAt: '2024-03-28T08:00:00Z',
    updatedAt: '2024-03-28T08:00:00Z',
  },
  {
    id: '2',
    startTime: '2024-03-30T14:00:00Z',
    endTime: '2024-03-30T15:00:00Z',
    createdAt: '2024-03-28T09:00:00Z',
    updatedAt: '2024-03-28T09:00:00Z',
  },
  {
    id: '3',
    startTime: '2024-03-31T16:00:00Z',
    endTime: '2024-03-31T17:00:00Z',
    createdAt: '2024-03-28T10:00:00Z',
    updatedAt: '2024-03-28T10:00:00Z',
  },
];

export const Default: Story = {
  args: {
    reservations: mockReservations,
    isLoading: false,
    onDelete: (id) => {
      console.log(`Delete reservation with id: ${id}`);
    },
  },
};

export const Loading: Story = {
  args: {
    reservations: [],
    isLoading: true,
    onDelete: () => {},
  },
};

export const Empty: Story = {
  args: {
    reservations: [],
    isLoading: false,
    onDelete: () => {},
  },
};