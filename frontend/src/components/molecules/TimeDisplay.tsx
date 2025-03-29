'use client';

import Typography from '../atoms/Typography';

interface TimeDisplayProps {
  start: Date;
  end: Date;
  className?: string;
}

export default function TimeDisplay({ start, end, className = '' }: TimeDisplayProps) {
  const formatTime = (date: Date) => {
    return date.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
  };

  return (
    <Typography variant="body" className={className}>
      {formatTime(start)} - {formatTime(end)}
    </Typography>
  );
}