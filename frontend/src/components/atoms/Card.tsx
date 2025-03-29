'use client';

import React from 'react';

interface CardProps {
  children: React.ReactNode;
  className?: string;
  hasShadow?: boolean;
  padding?: 'none' | 'small' | 'medium' | 'large';
}

export default function Card({
  children,
  className = '',
  hasShadow = true,
  padding = 'medium',
}: CardProps) {
  const paddingClasses = {
    none: 'p-0',
    small: 'p-2',
    medium: 'p-4',
    large: 'p-6',
  };

  const shadowClass = hasShadow ? 'shadow-md' : '';
  
  return (
    <div className={`bg-white rounded-lg ${shadowClass} ${paddingClasses[padding]} ${className}`}>
      {children}
    </div>
  );
}