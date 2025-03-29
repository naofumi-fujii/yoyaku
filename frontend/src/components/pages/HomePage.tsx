'use client';

import { useState, useEffect } from 'react';
import axios from 'axios';
import toast from 'react-hot-toast';
import { Reservation } from '@/types';
import Calendar from '../organisms/Calendar';
import ReservationModal from '../organisms/ReservationModal';
import ReservationList from '../organisms/ReservationList';
import MainLayout from '../templates/MainLayout';
import SplitLayout from '../templates/SplitLayout';
import Typography from '../atoms/Typography';
import Spinner from '../atoms/Spinner';

export default function HomePage() {
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
      <div className="flex min-h-screen flex-col items-center justify-center p-6">
        <div className="text-center">
          <Typography variant="h1" className="mb-8">予約システム</Typography>
          <div className="flex flex-col items-center space-y-4">
            <Spinner size="large" />
            <Typography variant="body" className="text-xl">
              {connectionError 
                ? 'データベースへの接続に失敗しました。再試行中...' 
                : 'データベースに接続中...'}
            </Typography>
          </div>
        </div>
      </div>
    );
  }

  return (
    <MainLayout title="予約システム">
      <SplitLayout
        leftTitle="カレンダー"
        rightTitle="予約一覧"
        leftContent={
          <Calendar 
            reservations={reservations}
            onSelect={handleSelect}
            onDelete={handleDeleteReservation}
          />
        }
        rightContent={
          <ReservationList 
            reservations={reservations} 
            isLoading={isLoading} 
            onDelete={handleDeleteReservation}
          />
        }
      />
      
      <ReservationModal 
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        selectedRange={selectedRange}
        onConfirm={handleConfirmReservation}
      />
    </MainLayout>
  );
}