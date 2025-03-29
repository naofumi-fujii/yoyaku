'use client';

interface SpinnerProps {
  size?: 'small' | 'medium' | 'large';
  color?: string;
}

export default function Spinner({ size = 'medium', color = 'border-blue-700' }: SpinnerProps) {
  const sizeClasses = {
    small: 'h-4 w-4',
    medium: 'h-8 w-8',
    large: 'h-12 w-12',
  };

  return (
    <div className={`animate-spin rounded-full ${sizeClasses[size]} border-b-2 ${color}`} />
  );
}