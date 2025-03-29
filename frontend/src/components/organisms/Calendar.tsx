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
import Card from '../atoms/Card';
import Typography from '../atoms/Typography';
import ConfirmationButtons from '../molecules/ConfirmationButtons';

interface CalendarProps {
  reservations: Reservation[];
  onSelect: ({ start, end }: { start: Date; end: Date }) => void;
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

  const events = reservations?.map((reservation) => ({
    id: reservation.id,
    title: '予約済み',
    start: new Date(reservation.startTime),
    end: new Date(reservation.endTime),
    backgroundColor: '#3788d8',
    borderColor: '#3788d8',
    extendedProps: { isReservation: true },
  })) || [];

  const handleEventClick = (eventInfo: any) => {
    // Only handle reservation events
    if (eventInfo.event.extendedProps.isReservation) {
      const eventId = eventInfo.event.id;
      const start = eventInfo.event.start;
      const end = eventInfo.event.end;
      
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
      <Card>
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
      </Card>

      {/* Delete Confirmation Modal */}
      {showDeleteModal && selectedEvent && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
          <Card padding="large" className="max-w-md w-full">
            <Typography variant="h2" className="mb-4">予約キャンセル</Typography>
            
            <div className="mb-6">
              <div className="mb-2">
                <Typography variant="label" className="mr-2">日付:</Typography>
                <Typography variant="body">{formatDate(selectedEvent.start)}</Typography>
              </div>
              <div className="mb-3">
                <Typography variant="label" className="mr-2">時間:</Typography>
                <Typography variant="body">
                  {selectedEvent.start.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' })} - {selectedEvent.end.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' })}
                </Typography>
              </div>
              <Typography variant="body" className="text-red-600">この予約をキャンセルしますか？</Typography>
            </div>
            
            <ConfirmationButtons
              onCancel={() => setShowDeleteModal(false)}
              onConfirm={handleDeleteReservation}
              cancelText="戻る"
              confirmText="キャンセルする"
              confirmVariant="danger"
            />
          </Card>
        </div>
      )}
    </>
  );
}