'use client';

import React from 'react';

interface TypographyProps {
  variant?: 'h1' | 'h2' | 'h3' | 'h4' | 'body' | 'caption' | 'label';
  color?: string;
  className?: string;
  children: React.ReactNode;
}

export default function Typography({
  variant = 'body',
  color = 'text-gray-900',
  className = '',
  children,
}: TypographyProps) {
  const variantClasses = {
    h1: 'text-2xl font-bold md:text-3xl',
    h2: 'text-xl font-bold md:text-2xl',
    h3: 'text-lg font-bold md:text-xl',
    h4: 'text-md font-medium md:text-lg',
    body: 'text-base',
    caption: 'text-sm',
    label: 'text-sm font-medium',
  };

  const combinedClassName = `${variantClasses[variant]} ${color} ${className}`;

  switch (variant) {
    case 'h1':
      return <h1 className={combinedClassName}>{children}</h1>;
    case 'h2':
      return <h2 className={combinedClassName}>{children}</h2>;
    case 'h3':
      return <h3 className={combinedClassName}>{children}</h3>;
    case 'h4':
      return <h4 className={combinedClassName}>{children}</h4>;
    case 'caption':
      return <span className={combinedClassName}>{children}</span>;
    case 'label':
      return <label className={combinedClassName}>{children}</label>;
    default:
      return <p className={combinedClassName}>{children}</p>;
  }
}