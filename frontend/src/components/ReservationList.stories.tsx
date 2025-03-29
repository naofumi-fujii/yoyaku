import type { Meta, StoryObj } from '@storybook/react';
import ReservationList from './ReservationList';

const meta: Meta<typeof ReservationList> = {
  component: ReservationList,
  title: 'Components/ReservationList',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof ReservationList>;

const mockReservations = [
  {
    id: '1',
    startTime: '2024-12-01T10:00:00Z',
    endTime: '2024-12-01T11:00:00Z',
    createdAt: '2024-11-30T08:00:00Z',
    updatedAt: '2024-11-30T08:00:00Z',
  },
  {
    id: '2',
    startTime: '2024-12-02T14:00:00Z',
    endTime: '2024-12-02T15:00:00Z',
    createdAt: '2024-11-30T09:00:00Z',
    updatedAt: '2024-11-30T09:00:00Z',
  },
  {
    id: '3',
    startTime: '2024-12-03T16:00:00Z',
    endTime: '2024-12-03T17:00:00Z',
    createdAt: '2024-11-30T10:00:00Z',
    updatedAt: '2024-11-30T10:00:00Z',
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