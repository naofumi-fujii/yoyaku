import type { Meta, StoryObj } from '@storybook/react';
import ReservationModal from './ReservationModal';

const meta: Meta<typeof ReservationModal> = {
  component: ReservationModal,
  title: 'Components/ReservationModal',
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof ReservationModal>;

// Set fixed dates for the story
const startDate = new Date('2024-12-01T10:00:00Z');
const endDate = new Date('2024-12-01T11:00:00Z');

export const Default: Story = {
  args: {
    isOpen: true,
    selectedRange: {
      start: startDate,
      end: endDate,
    },
    onClose: () => console.log('Modal closed'),
    onConfirm: () => console.log('Reservation confirmed'),
  },
};

export const Hidden: Story = {
  args: {
    isOpen: false,
    selectedRange: {
      start: startDate,
      end: endDate,
    },
    onClose: () => console.log('Modal closed'),
    onConfirm: () => console.log('Reservation confirmed'),
  },
};