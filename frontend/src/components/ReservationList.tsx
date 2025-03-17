'use client';

import { formatDate } from '@/lib/date';
import { Reservation } from '@/types';

interface ReservationListProps {
  reservations: Reservation[];
  isLoading: boolean;
  onDelete: (id: string) => void;
}

export default function ReservationList({ reservations, isLoading, onDelete }: ReservationListProps) {
  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-40">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-700"></div>
      </div>
    );
  }

  if (!reservations || reservations.length === 0) {
    return (
      <div className="bg-white p-4 rounded-lg shadow-md text-center">
        <p className="text-gray-500">予約はまだありません</p>
      </div>
    );
  }

  // Sort reservations by start time (most recent first)
  const sortedReservations = [...(reservations || [])].sort(
    (a, b) => new Date(b.startTime).getTime() - new Date(a.startTime).getTime()
  );

  return (
    <div className="bg-white rounded-lg shadow-md overflow-hidden">
      <ul className="divide-y divide-gray-200">
        {sortedReservations.map((reservation) => (
          <li key={reservation.id} className="p-4 hover:bg-gray-50">
            <div className="flex justify-between items-start">
              <div>
                <p className="font-medium">
                  {formatDate(new Date(reservation.startTime))}
                </p>
                <p className="text-sm text-gray-600">
                  {new Date(reservation.startTime).toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' })}
                  {' - '}
                  {new Date(reservation.endTime).toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' })}
                </p>
              </div>
              <button
                onClick={() => onDelete(reservation.id)}
                className="text-red-600 hover:text-red-800 text-sm font-medium"
              >
                キャンセル
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}
