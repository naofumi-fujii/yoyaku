'use client';

import { formatDate } from '@/lib/date';
import Typography from '../atoms/Typography';

interface DateDisplayProps {
  date: Date;
  className?: string;
  variant?: 'h1' | 'h2' | 'h3' | 'h4' | 'body' | 'caption' | 'label';
}

export default function DateDisplay({ 
  date, 
  className = '',
  variant = 'body' 
}: DateDisplayProps) {
  return (
    <Typography variant={variant} className={className}>
      {formatDate(date)}
    </Typography>
  );
}