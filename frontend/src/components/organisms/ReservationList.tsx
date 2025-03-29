'use client';

import { Reservation } from '@/types';
import Card from '../atoms/Card';
import Spinner from '../atoms/Spinner';
import Typography from '../atoms/Typography';
import ReservationItem from '../molecules/ReservationItem';

interface ReservationListProps {
  reservations: Reservation[];
  isLoading: boolean;
  onDelete: (id: string) => void;
}

export default function ReservationList({ reservations, isLoading, onDelete }: ReservationListProps) {
  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-40">
        <Spinner size="medium" />
      </div>
    );
  }

  if (!reservations || reservations.length === 0) {
    return (
      <Card>
        <Typography variant="body" className="text-center text-gray-500">
          予約はまだありません
        </Typography>
      </Card>
    );
  }

  // Sort reservations by start time (most recent first)
  const sortedReservations = [...(reservations || [])].sort(
    (a, b) => new Date(b.startTime).getTime() - new Date(a.startTime).getTime()
  );

  return (
    <Card className="overflow-hidden p-0">
      <ul className="divide-y divide-gray-200">
        {sortedReservations.map((reservation) => (
          <ReservationItem 
            key={reservation.id}
            reservation={reservation}
            onDelete={onDelete}
          />
        ))}
      </ul>
    </Card>
  );
}