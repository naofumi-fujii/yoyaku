'use client';

import { useEffect, useRef, useState } from 'react';
import FullCalendar from '@fullcalendar/react';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import jaLocale from '@fullcalendar/core/locales/ja';
import { Reservation } from '@/types';
import { formatDate } from '@/lib/date';
import toast from 'react-hot-toast';

interface CalendarProps {
  reservations: Reservation[];
  onSelect: (info: { start: Date; end: Date }) => void;
  onDelete?: (id: string) => void;
}

export default function Calendar({ reservations, onSelect, onDelete }: CalendarProps) {
  const calendarRef = useRef<FullCalendar | null>(null);
  const [showDeleteModal, setShowDeleteModal] = useState(false);
  const [selectedEvent, setSelectedEvent] = useState<{id: string, start: Date, end: Date} | null>(null);

  useEffect(() => {
    // Force refresh when reservations change
    const calendarApi = calendarRef.current?.getApi();
    if (calendarApi) {
      calendarApi.refetchEvents();
    }
  }, [reservations]);

  const events = reservations.map((reservation) => ({
    id: reservation.id,
    title: '予約済み',
    start: new Date(reservation.startTime),
    end: new Date(reservation.endTime),
    backgroundColor: '#3788d8',
    borderColor: '#3788d8',
    extendedProps: { isReservation: true },
  }));

  const handleEventClick = (info: any) => {
    // Only handle reservation events
    if (info.event.extendedProps.isReservation) {
      const eventId = info.event.id;
      const start = info.event.start;
      const end = info.event.end;
      
      setSelectedEvent({ id: eventId, start, end });
      setShowDeleteModal(true);
    }
  };

  const handleDeleteReservation = () => {
    if (selectedEvent && onDelete) {
      onDelete(selectedEvent.id);
      setShowDeleteModal(false);
      toast.success('予約をキャンセルしました');
    }
  };

  return (
    <>
      <div className="bg-white p-4 rounded-lg shadow-md">
        <FullCalendar
          ref={calendarRef}
          plugins={[dayGridPlugin, timeGridPlugin, interactionPlugin]}
          initialView="timeGridWeek"
          headerToolbar={{
            left: 'prev,next today',
            center: 'title',
            right: 'dayGridMonth,timeGridWeek,timeGridDay',
          }}
          locales={[jaLocale]}
          locale="ja"
          selectable={true}
          selectMirror={true}
          dayMaxEvents={true}
          events={events}
          select={onSelect}
          eventClick={handleEventClick}
          height="auto"
          allDaySlot={false}
          slotMinTime="09:00:00"
          slotMaxTime="20:00:00"
        />
      </div>

      {/* Delete Confirmation Modal */}
      {showDeleteModal && selectedEvent && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
          <div className="bg-white p-6 rounded-lg shadow-xl max-w-md w-full">
            <h2 className="text-xl font-bold mb-4">予約キャンセル</h2>
            
            <div className="mb-6">
              <p className="text-gray-700 mb-2">
                <span className="font-medium">日付:</span> {formatDate(selectedEvent.start)}
              </p>
              <p className="text-gray-700">
                <span className="font-medium">時間:</span> {selectedEvent.start.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' })} - {selectedEvent.end.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' })}
              </p>
              <p className="mt-3 text-red-600">この予約をキャンセルしますか？</p>
            </div>
            
            <div className="flex justify-end space-x-3">
              <button
                onClick={() => setShowDeleteModal(false)}
                className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100"
              >
                戻る
              </button>
              <button
                onClick={handleDeleteReservation}
                className="px-4 py-2 bg-red-600 text-white rounded-md hover:bg-red-700"
              >
                キャンセルする
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  );
}
