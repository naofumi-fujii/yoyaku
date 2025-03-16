'use client';

import { formatDate } from '@/lib/date';

interface ReservationModalProps {
  isOpen: boolean;
  onClose: () => void;
  selectedRange: { start: Date; end: Date } | null;
  onConfirm: () => void;
}

export default function ReservationModal({ isOpen, onClose, selectedRange, onConfirm }: ReservationModalProps) {
  if (!isOpen || !selectedRange) return null;

  const startDate = formatDate(selectedRange.start);
  const startTime = selectedRange.start.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
  const endTime = selectedRange.end.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white p-6 rounded-lg shadow-xl max-w-md w-full">
        <h2 className="text-xl font-bold mb-4">予約確認</h2>
        
        <div className="mb-6">
          <p className="text-gray-700 mb-2">
            <span className="font-medium">日付:</span> {startDate}
          </p>
          <p className="text-gray-700">
            <span className="font-medium">時間:</span> {startTime} - {endTime}
          </p>
        </div>
        
        <div className="flex justify-end space-x-3">
          <button
            onClick={onClose}
            className="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100"
          >
            キャンセル
          </button>
          <button
            onClick={onConfirm}
            className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
          >
            予約する
          </button>
        </div>
      </div>
    </div>
  );
}
