'use client';

import Card from '../atoms/Card';
import Typography from '../atoms/Typography';
import DateDisplay from '../molecules/DateDisplay';
import TimeDisplay from '../molecules/TimeDisplay';
import ConfirmationButtons from '../molecules/ConfirmationButtons';

interface ReservationModalProps {
  isOpen: boolean;
  onClose: () => void;
  selectedRange: { start: Date; end: Date } | null;
  onConfirm: () => void;
}

export default function ReservationModal({ isOpen, onClose, selectedRange, onConfirm }: ReservationModalProps) {
  if (!isOpen || !selectedRange) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
      <Card padding="large" className="max-w-md w-full">
        <Typography variant="h2" className="mb-4">予約確認</Typography>
        
        <div className="mb-6">
          <div className="mb-2">
            <Typography variant="label" className="mr-2">日付:</Typography>
            <DateDisplay date={selectedRange.start} />
          </div>
          <div>
            <Typography variant="label" className="mr-2">時間:</Typography>
            <TimeDisplay start={selectedRange.start} end={selectedRange.end} />
          </div>
        </div>
        
        <ConfirmationButtons
          onCancel={onClose}
          onConfirm={onConfirm}
          cancelText="戻る"
          confirmText="予約する"
        />
      </Card>
    </div>
  );
}