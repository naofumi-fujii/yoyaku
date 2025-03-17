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
  const [isConnected, setIsConnected] = useState(false);
  const [connectionError, setConnectionError] = useState(false);

  const fetchReservations = async () => {
    try {
      setIsLoading(true);
      const response = await axios.get('/api/reservations');
      setReservations(response.data);
      setIsConnected(true);
      setConnectionError(false);
    } catch (error) {
      console.error('Failed to fetch reservations:', error);
      if (!isConnected) {
        setConnectionError(true);
        // Retry connection after 2 seconds
        setTimeout(fetchReservations, 2000);
      } else {
        toast.error('予約の取得に失敗しました');
      }
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

  if (!isConnected) {
    return (
      <main className="flex min-h-screen flex-col items-center justify-center p-6 md:p-12">
        <div className="text-center">
          <h1 className="text-3xl font-bold mb-8">予約システム</h1>
          <div className="flex flex-col items-center space-y-4">
            <div className="animate-spin rounded-full h-16 w-16 border-b-4 border-blue-700"></div>
            <p className="text-xl">
              {connectionError 
                ? 'データベースへの接続に失敗しました。再試行中...' 
                : 'データベースに接続中...'}
            </p>
          </div>
        </div>
      </main>
    );
  }

  return (
    <main className="flex min-h-screen flex-col items-center p-6 md:p-12">
      <h1 className="text-3xl font-bold mb-8">予約システム</h1>
      
      <div className="w-full max-w-6xl grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2">
          <Calendar 
            reservations={reservations}
            onSelect={handleSelect}
            onDelete={handleDeleteReservation}
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
