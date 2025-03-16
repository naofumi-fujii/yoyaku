'use client';

import { useState, useEffect } from 'react';
import Calendar from '@/components/Calendar';
import ReservationModal from '@/components/ReservationModal';
import ReservationList from '@/components/ReservationList';
import axios from 'axios';
import toast from 'react-hot-toast';
import { Reservation } from '@/types';

export default function Home() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedRange, setSelectedRange] = useState<{ start: Date; end: Date } | null>(null);
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  const fetchReservations = async () => {
    try {
      setIsLoading(true);
      const response = await axios.get('/api/reservations');
      setReservations(response.data);
    } catch (error) {
      console.error('Failed to fetch reservations:', error);
      toast.error('予約の取得に失敗しました');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchReservations();
  }, []);

  const handleSelect = (info: { start: Date; end: Date }) => {
    setSelectedRange({ start: info.start, end: info.end });
    setIsModalOpen(true);
  };

  const handleConfirmReservation = async () => {
    if (!selectedRange) return;

    try {
      await axios.post('/api/reservations', {
        startTime: selectedRange.start.toISOString(),
        endTime: selectedRange.end.toISOString(),
      });
      toast.success('予約が完了しました');
      setIsModalOpen(false);
      fetchReservations();
    } catch (error) {
      console.error('Failed to create reservation:', error);
      toast.error('予約の作成に失敗しました');
    }
  };

  const handleDeleteReservation = async (id: string) => {
    try {
      await axios.delete(`/api/reservations/${id}`);
      toast.success('予約をキャンセルしました');
      fetchReservations();
    } catch (error) {
      console.error('Failed to delete reservation:', error);
      toast.error('予約のキャンセルに失敗しました');
    }
  };

  return (
    <main className="flex min-h-screen flex-col items-center p-6 md:p-12">
      <h1 className="text-3xl font-bold mb-8">予約システム</h1>
      
      <div className="w-full max-w-6xl grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2">
          <Calendar 
            reservations={reservations}
            onSelect={handleSelect}
          />
        </div>
        
        <div>
          <h2 className="text-xl font-semibold mb-4">予約一覧</h2>
          <ReservationList 
            reservations={reservations} 
            isLoading={isLoading} 
            onDelete={handleDeleteReservation}
          />
        </div>
      </div>
      
      <ReservationModal 
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        selectedRange={selectedRange}
        onConfirm={handleConfirmReservation}
      />
    </main>
  );
}
