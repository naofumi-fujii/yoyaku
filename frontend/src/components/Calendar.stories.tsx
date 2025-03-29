import type { Meta, StoryObj } from '@storybook/react';
import Calendar from './Calendar';
import { Reservation } from '@/types';

const meta: Meta<typeof Calendar> = {
  component: Calendar,
  title: 'Components/Calendar',
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof Calendar>;

const today = new Date();
const tomorrow = new Date(today);
tomorrow.setDate(today.getDate() + 1);

// Helper to format dates for the mock data
const formatISOTime = (date: Date, hour: number, minute: number = 0): string => {
  const d = new Date(date);
  d.setHours(hour, minute, 0, 0);
  return d.toISOString();
};

const mockReservations: Reservation[] = [
  {
    id: '1',
    startTime: formatISOTime(today, 10),
    endTime: formatISOTime(today, 11),
    createdAt: '2024-03-28T08:00:00Z',
    updatedAt: '2024-03-28T08:00:00Z',
  },
  {
    id: '2',
    startTime: formatISOTime(tomorrow, 14),
    endTime: formatISOTime(tomorrow, 15),
    createdAt: '2024-03-28T09:00:00Z',
    updatedAt: '2024-03-28T09:00:00Z',
  },
];

export const Default: Story = {
  args: {
    reservations: mockReservations,
    onSelect: ({ start, end }) => {
      console.log('Selected time slot:', { start, end });
    },
    onDelete: (id) => {
      console.log(`Delete reservation with id: ${id}`);
    },
  },
};

export const Empty: Story = {
  args: {
    reservations: [],
    onSelect: ({ start, end }) => {
      console.log('Selected time slot:', { start, end });
    },
    onDelete: (id) => {
      console.log(`Delete reservation with id: ${id}`);
    },
  },
};