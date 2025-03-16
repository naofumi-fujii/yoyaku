'use client';

import { useEffect, useRef } from 'react';
import FullCalendar from '@fullcalendar/react';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import jaLocale from '@fullcalendar/core/locales/ja';
import { Reservation } from '@/types';

interface CalendarProps {
  reservations: Reservation[];
  onSelect: (info: { start: Date; end: Date }) => void;
}

export default function Calendar({ reservations, onSelect }: CalendarProps) {
  const calendarRef = useRef<FullCalendar | null>(null);

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
  }));

  return (
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
        height="auto"
        allDaySlot={false}
        slotMinTime="09:00:00"
        slotMaxTime="20:00:00"
      />
    </div>
  );
}
