'use client';

import React from 'react';
import Typography from '../atoms/Typography';

interface SplitLayoutProps {
  leftContent: React.ReactNode;
  rightContent: React.ReactNode;
  leftTitle?: string;
  rightTitle?: string;
  leftWidth?: string;
  rightWidth?: string;
}

export default function SplitLayout({
  leftContent,
  rightContent,
  leftTitle,
  rightTitle,
  leftWidth = 'w-full lg:w-2/3',
  rightWidth = 'w-full lg:w-1/3',
}: SplitLayoutProps) {
  return (
    <div className="flex flex-col lg:flex-row gap-6">
      {/* 左カラム */}
      <div className={`${leftWidth}`}>
        {leftTitle && (
          <Typography variant="h3" className="mb-4">
            {leftTitle}
          </Typography>
        )}
        {leftContent}
      </div>
      
      {/* 右カラム */}
      <div className={`${rightWidth}`}>
        {rightTitle && (
          <Typography variant="h3" className="mb-4">
            {rightTitle}
          </Typography>
        )}
        {rightContent}
      </div>
    </div>
  );
}