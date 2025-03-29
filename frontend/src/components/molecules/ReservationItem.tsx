'use client';

import { Reservation } from '@/types';
import DateDisplay from './DateDisplay';
import TimeDisplay from './TimeDisplay';
import Button from '../atoms/Button';

interface ReservationItemProps {
  reservation: Reservation;
  onDelete: (id: string) => void;
}

export default function ReservationItem({ reservation, onDelete }: ReservationItemProps) {
  const startDate = new Date(reservation.startTime);
  const endDate = new Date(reservation.endTime);

  return (
    <li className="p-4 hover:bg-gray-50">
      <div className="flex justify-between items-start">
        <div>
          <DateDisplay date={startDate} variant="body" className="font-medium" />
          <TimeDisplay 
            start={startDate} 
            end={endDate} 
            className="text-sm text-gray-600" 
          />
        </div>
        <Button 
          variant="danger" 
          size="small"
          onClick={() => onDelete(reservation.id)}
        >
          キャンセル
        </Button>
      </div>
    </li>
  );
}